package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"testing"
	"time"
)

type userPattern struct {
	name string
}

type userKey int

func TestPatternWithValue(t *testing.T) {
	const uk userKey = 0

	u := userPattern{
		name: "Guiwoo",
	}

	ctx := context.WithValue(context.Background(), uk, &u)

	if user, ok := ctx.Value(uk).(*userPattern); ok {
		fmt.Println("User : ", user)
	}

	if _, ok := ctx.Value(0).(*userPattern); !ok {
		fmt.Println("User not found")
	}
}

func TestPatternWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Moving on")
	case <-ctx.Done():
		fmt.Println("work completed")
	}
}

type dataPattern struct {
	UserID string
}

func TestPatternWithDeadLine(t *testing.T) {
	deadline := time.Now().Add(150 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	ch := make(chan dataPattern, 1)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- dataPattern{"123"}
	}()

	select {
	case d := <-ch:
		fmt.Println("work completed ", d)
	case <-ctx.Done():
		fmt.Println("work cancelled by deadline")
	}
}

func TestPatternWithTimeout(t *testing.T) {
	duration := 150 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan dataPattern, 1)
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch <- dataPattern{"123"}
	}()

	select {
	case d := <-ch:
		fmt.Println("work completed ", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}

func TestPatternRequestAndResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "https://www.ardanlabs.com/blog/post/index.xml", nil)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithTimeout(req.Context(), 50*time.Millisecond)
	defer cancel()

	tr := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := http.Client{
		Transport: &tr,
	}

	ch := make(chan error, 1)
	go func() {
		log.Println("Starting Request")
		resp, err := client.Do(req)
		if err != nil {
			ch <- err
			return
		}
		defer resp.Body.Close()

		io.Copy(os.Stdout, resp.Body)
		ch <- nil
	}()

	select {
	case <-ctx.Done():
		log.Println("timeout, cancel work...")
		cancel()
		log.Println(<-ch)
	case err := <-ch:
		if err != nil {
			log.Println(err)
		}
	}
}
