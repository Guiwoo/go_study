package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	basicSendRecv := func() {
		ch := make(chan interface{})
		go func() {
			ch <- "hello"
		}()
		fmt.Println(<-ch)
	}
	signalClose := func() {
		ch := make(chan interface{})
		go func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("signal event")
			close(ch)
		}()
		<-ch
		fmt.Println("event received")
	}

	fmt.Printf("\n => Basics of a send and receive\n")
	basicSendRecv()
	fmt.Printf("\n=> Close a channel to signal an event\n")
	signalClose()
}

func TestDoubleChannel(t *testing.T) {
	signalAck := func() {
		ch := make(chan string)
		go func() {
			fmt.Println(<-ch)
			ch <- "ok done"
		}()
		ch <- "do this"
		fmt.Println(<-ch)
	}

	selectRecv := func() {
		ch := make(chan string)
		defer close(ch)
		go func() {
			fmt.Println("run")
			defer func() {
				fmt.Println("go routine return")
			}()
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "work"
		}()
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-time.After(100 * time.Millisecond):
			fmt.Println("time out")
		}
	}

	selectSend := func() {
		ch := make(chan string)
		defer close(ch)

		go func() {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "work"
		}()
		select {
		case ch <- "work":
			fmt.Println("send work")
		case <-time.After(100 * time.Millisecond):
			fmt.Println("time out")
		}
	}
	selectDrop := func() {
		ch := make(chan int, 5)
		defer close(ch)

		go func() {
			defer func() {
				fmt.Println("return go routine")
			}()
			for v := range ch {
				fmt.Println("recv", v)
			}
		}()

		for i := 0; i < 20; i++ {
			select {
			case ch <- i:
				fmt.Println("send work", i)
			default:
				fmt.Println("drop", i)
			}
		}
	}

	fmt.Printf("\n=> Double signal\n")
	signalAck()

	fmt.Printf("\n=> Select and receive\n")
	selectRecv()

	fmt.Printf("\n=> Select and send\n")
	selectSend()

	fmt.Printf("\n=> Select and drop\n")
	selectDrop()
}

func TestTennis(t *testing.T) {
	player := func(name string, court chan int) {
		for {
			ball, ok := <-court
			if !ok {
				fmt.Printf("Player %s Won\n", name)
				return
			}
			n := rand.Intn(100)
			if n%13 == 0 {
				fmt.Printf("Player %s Missed\n", name)
				close(court)
				return
			}
			fmt.Printf("Player %s Hit %d \n", name, ball)
			ball++
			court <- ball
		}
	}

	court := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		player("Hoanh", court)
	}()

	go func() {
		defer wg.Done()
		player("Andrew", court)
	}()

	court <- 1

	wg.Wait()
}

func TestRace(t *testing.T) {
	var wg sync.WaitGroup
	Runner := func(track chan int) {}
	Runner = func(track chan int) {
		const maxEchange = 4
		var exchange int

		baton := <-track
		fmt.Printf("Runner %d Running With Baton\n", baton)

		if baton < maxEchange {
			exchange = baton + 1
			fmt.Printf("Runner %d To The Line \n", exchange)
			go Runner(track)
		}
		time.Sleep(100 * time.Millisecond)
		if baton == maxEchange {
			fmt.Printf("Runner %d Finished, Race Over \n", baton)
			wg.Done()
			return
		}
		fmt.Printf("Runner %d Exchange WIth Runner %d\n", baton, exchange)
		track <- exchange
	}

	track := make(chan int)
	wg.Add(1)

	go Runner(track)
	track <- 1

	wg.Wait()
}

type result struct {
	id  int
	op  string
	err error
}

func TestBufferedChan(t *testing.T) {
	insertUser := func(id int) result {
		r := result{
			id: id,
			op: fmt.Sprintf("insert Users value %d", id),
		}
		if rand.Intn(10) == 0 {
			r.err = fmt.Errorf("unfortunely insert fail %d into User table", id)
		}
		return r
	}

	insertTrans := func(id int) result {
		r := result{
			id: id,
			op: fmt.Sprintf("insert trans value %d", id),
		}
		if rand.Intn(10) == 0 {
			r.err = fmt.Errorf("unfortunely insert fail %d into User table", id)
		}
		return r
	}

	const (
		rout   = 10
		insert = rout * 2
	)

	ch := make(chan result, insert)
	waitInsert := insert

	for i := 0; i < rout; i++ {
		go func(id int) {
			ch <- insertUser(id)
			ch <- insertTrans(id)
		}(i)
	}
	for waitInsert > 0 {
		r := <-ch
		log.Printf("N : %d Id : %d OP: %s Err : %v\n", waitInsert, r.id, r.op, r.err)
		waitInsert--
	}
	log.Printf("Insert Complete\n")
}

func TestShutDown(t *testing.T) {
	const timeoutSec = 3 * time.Second

	var (
		sigChan  = make(chan os.Signal, 1)
		timeout  = time.After(timeoutSec)
		complete = make(chan error)
		shutdown = make(chan interface{})
	)

	checkShutdown := func() bool {
		select {
		case <-shutdown:
			log.Println("checkShutdown - Shutdown Early")
			return true
		default:
			return false
		}
	}

	doWork := func() error {
		log.Println("Processor - Task 1")
		time.Sleep(2 * time.Second)

		if checkShutdown() {
			return errors.New("Early Shutdown")
		}
		log.Println("Processor - Task 2")
		time.Sleep(1 * time.Second)

		if checkShutdown() {
			return errors.New("Early Shutdown")
		}

		log.Println("Processor - Task 3")
		time.Sleep(1 * time.Second)
		return nil
	}

	processor := func(comp chan<- error) {
		log.Println("Processor - Starting")
		var err error
		defer func() {
			if r := recover(); r != nil {
				log.Println("Processor - Panic", r)
			}
			comp <- err
		}()

		err = doWork()
		log.Println("Processor - Completed")
	}

	log.Println("Starting Process")
	signal.Notify(sigChan, os.Interrupt)

	log.Println("Launching Processors")
	go processor(complete)

CounterLoop:
	for {
		select {
		case <-sigChan:
			log.Println("os interrupt")
			close(shutdown)
			sigChan = nil
		case <-timeout:
			log.Println("tiem out - killing program")
			os.Exit(1)
		case err := <-complete:
			log.Printf("task complete : error [%s]", err)
			break CounterLoop
		}
	}

	log.Println("Process end")
}
