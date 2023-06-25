package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test01(t *testing.T) {

	locale := func(done <-chan interface{}) (string, error) {
		select {
		case <-done:
			return "", fmt.Errorf("canceled")
		case <-time.After(5 * time.Second):
		}
		return "EN/US", nil
	}

	genGreeting := func(done <-chan interface{}) (string, error) {
		switch locale, err := locale(done); {
		case err != nil:
			return "", err
		case locale == "EN/US":
			return "hello", nil
		}
		return "", fmt.Errorf("unsupported locale")
	}

	printGreeting := func(done <-chan interface{}) error {
		greeting, err := genGreeting(done)
		if err != nil {
			return err
		}
		fmt.Printf("%s Wrold!\n", greeting)
		return nil
	}
	genFarewell := func(done <-chan interface{}) (string, error) {
		switch locale, err := locale(done); {
		case err != nil:
			return "", err
		case locale == "EN/US":
			return "bye bye", nil
		}
		return "", fmt.Errorf("upsupported locale")
	}

	printFareWell := func(done <-chan interface{}) error {
		farewell, err := genFarewell(done)
		if err != nil {
			return err
		}
		fmt.Printf("%s world!\n", farewell)
		return nil
	}

	var wg sync.WaitGroup

	done := make(chan interface{})
	defer close(done)

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := printGreeting(done); err != nil {
			fmt.Printf("error is : %s", err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := printFareWell(done); err != nil {
			fmt.Printf("error is : %s", err)
		}
	}()

	wg.Wait()
}

func Test02(t *testing.T) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	locale := func(ctx context.Context) (string, error) {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(5 * time.Second):
		}
		return "EN/US", nil
	}

	genGreeting := func(ctx context.Context) (string, error) {
		ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
		defer cancel()
		switch loc, err := locale(ctx); {
		case err != nil:
			return "", err
		case loc == "EN/US":
			return "hello", nil
		}
		return "", fmt.Errorf("unsupported")
	}

	printGreeting := func(ctx context.Context) error {
		greeting, err := genGreeting(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("%s world\n", greeting)
		return nil
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("Can not printing greeting : %s", err)
			cancel()
		}
	}()

	go func() {
		defer wg.Done()
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("Can not printing farewell : %s", err)
			cancel()
		}
	}()

	wg.Wait()
}

func Test03(t *testing.T) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(2)

	locale := func(ctx context.Context) (string, error) {
		if deadline, ok := ctx.Deadline(); ok {
			if deadline.Sub(time.Now().Add(1*time.Minute)) <= 0 {
				return "", fmt.Errorf("unsupported locale")
			}
		}
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(1 * time.Minute):
			return "EN/US", nil
		}
	}

	genGreeting := func(ctx context.Context) (string, error) {
		ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
		defer cancel()
		defer wg.Done()
		switch loc, err := locale(ctx); {
		case err != nil:
			return "", err
		case loc == "EN/US":
			return "hello", nil
		}
		return "", fmt.Errorf("unsupported")
	}

	printGreeting := func(ctx context.Context) error {
		greeting, err := genGreeting(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("%s world!\n", greeting)
		return nil
	}
	printFarewell := func(ctx context.Context) error {
		defer wg.Done()
		farewell, err := genGreeting(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("%s world \n", farewell)
		return nil
	}

	go func() {
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("Erorr ouccr on print Greeting : %s", err)
			cancel()
		}
	}()

	go func() {
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("Error occur on print Farewell : %s", err)
		}
	}()

	wg.Wait()
	fmt.Println("All go routines done")
}

func Test04(t *testing.T) {
	var id, token string

	HandleResponse := func(ctx context.Context) {
		id = ctx.Value("userId").(string)
		token = ctx.Value("token").(string)
		fmt.Printf("handling response for %v %v", ctx.Value("userId"), ctx.Value("token"))
	}

	processRequest := func(id, token string) {
		ctx := context.WithValue(context.Background(), "userId", id)
		ctx = context.WithValue(ctx, "token", token)
		HandleResponse(ctx)
	}

	processRequest("guiwoo", "abc123")

	if id != "guiwoo" || token != "abc123" {
		t.Errorf("does not store the values")
	}
}

func Test_06(t *testing.T) {
	type foo int
	type bar int

	m := make(map[any]string)

	m[foo(1)] = "This is Foo"
	m[bar(1)] = "This is Bar"

	fmt.Printf("%+v", m)
}

func Test_07(t *testing.T) {
	type ctxKey int
	const (
		ctxUserId ctxKey = iota
		ctxBank
	)
	userId := func(ctx context.Context) string {
		return ctx.Value(ctxUserId).(string)
	}
	bank := func(ctx context.Context) string {
		return ctx.Value(ctxBank).(string)
	}

	HandleResponse := func(ctx context.Context) {
		fmt.Printf("Handling response is id : %+v,%+v", userId(ctx), bank(ctx))
	}
	processRequest := func(id, bank string) {
		ctx := context.WithValue(context.Background(), ctxUserId, id)
		ctx = context.WithValue(ctx, ctxBank, bank)
		HandleResponse(ctx)
	}
	processRequest("guiwoo", "hyundai")
}
