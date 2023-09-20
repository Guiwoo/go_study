package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"testing"
	"time"
)

func TestError01(t *testing.T) {
	webCall := func() error {
		n := time.Now().Unix()
		if n%2 == 0 {
			return nil
		}
		return errors.New("지금은 홀수 입니다.")
	}

	if err := webCall(); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Life is Good")
}

func TestError02(t *testing.T) {
	var (
		ErrBadRequest   = errors.New("Bad Request")
		ErrBadPageMoved = errors.New("Page Moved")
		webCall         = func() error {
			n := time.Now().Unix()
			if n%2 == 0 {
				return ErrBadRequest
			}
			return ErrBadPageMoved
		}
	)
	if err := webCall(); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad Request")
			return
		case ErrBadPageMoved:
			fmt.Println("The Page Moved")
			return
		}
	}
	fmt.Println("Life is good")
}

func TestError03(t *testing.T) {
	type user struct {
		Name string
	}

	var u user
	err := json.Unmarshal([]byte(`{"name":"bill"}`), u)
	switch e := err.(type) {
	case *json.InvalidUnmarshalError:
		fmt.Printf("Invalid input type %v\n", e.Type)
	default:
		fmt.Println(err)
	}

	fmt.Println("Name", u.Name)
}

type client struct {
	name   string
	reader *bufio.Reader
}

func (c *client) TypeAsContext() {
	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			switch e := err.(type) {
			case *net.OpError:
				if !e.Temporary() {
					log.Println("Temporary: Client leaving chat")
					return
				}

			case *net.AddrError:
				if !e.Temporary() {
					log.Println("Temporary: Client leaving chat")
					return
				}

			case *net.DNSConfigError:
				if !e.Temporary() {
					log.Println("Temporary: Client leaving chat")
					return
				}

			default:
				if err == io.EOF {
					log.Println("EOF: Client leaving chat")
					return
				}

				log.Println("read-routine", err)
			}
		}
		fmt.Println(line)
	}
}

type CustomOneError struct {
	message string
}

func (c *CustomOneError) Error() string {
	return "one error" + c.message
}
func (c *CustomOneError) IsNotInteger() bool {
	return true
}

type CustomTwoError struct {
	message string
}

func (c *CustomTwoError) IsNotInteger() bool {
	return true
}

func (c *CustomTwoError) Error() string {
	return "two error" + c.message
}

type ErrNotInteger interface {
	IsNotInteger() bool
}

func TestError04_a(t *testing.T) {
	errChecker := func() error {
		n := time.Now().Unix()
		if n%2 == 0 {
			return &CustomOneError{"wrong"}
		} else {
			return &CustomTwoError{"something"}
		}
	}
	err := errChecker()
	switch e := err.(type) {
	case *CustomOneError:
		if e.IsNotInteger() {
			fmt.Println("Not Integer", e.message)
		}
	case *CustomTwoError:
		if e.IsNotInteger() {
			fmt.Println("Not Integer", e.message)
		}
	}
}

func TestError04(t *testing.T) {
	errChecker := func() error {
		n := time.Now().Unix()
		if n%2 == 0 {
			return &CustomOneError{"wrong"}
		} else {
			return &CustomTwoError{"something"}
		}
	}
	err := errChecker()
	switch e := err.(type) {
	case ErrNotInteger:
		fmt.Println("Not Integer Error ", e.IsNotInteger())
	default:
		fmt.Println("default ", err)
	}
}

func TestError05(t *testing.T) {

	thirdCall := func(i int) error {
		return &CustomOneError{"hoit"}
	}
	secondCall := func(i int) error {
		if err := thirdCall(i); err != nil {
			return fmt.Errorf("third call error %w", err)
		}
		return nil
	}
	firstCall := func(i int) error {
		if err := secondCall(i); err != nil {
			return fmt.Errorf("second call error %w", err)
		}
		return nil
	}
	for i := 0; i < 1; i++ {
		if err := firstCall(i); err != nil {
			var v *CustomOneError
			switch {
			case errors.As(err, &v):
				fmt.Println("custom One error found")
			}

			fmt.Println("Stack Trace")
			fmt.Printf("%+v\n", err)
			fmt.Printf("Stack Done")
		}
	}
}
