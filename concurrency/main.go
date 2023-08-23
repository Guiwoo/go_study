package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"time"
)

type list struct {
	signal chan interface{}
	name   int
}
type Handler struct {
	list map[int]list
	sync sync.Mutex
}

func handlerStream(done <-chan interface{}) <-chan interface{} {
	stream := make(chan interface{})
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		defer func() {
			log.Println("handler stream closed")
			close(stream)
			ticker.Stop()
		}()
		// do something int stream handler
		for {
			select {
			case _, ok := <-done:
				if !ok {
					return
				}
				log.Println("got done signal close")
				return
			case <-ticker.C:
				time.Sleep(1 * time.Second)
				stream <- "something on your mind"
			}
		}
	}()
	return stream
}

func (h *Handler) Handle(a, b int) {
	log.Printf("got %d and %d", a, b)

	if b == 0 {
		log.Println("got 0")
		if handler, ok := h.list[a]; ok {
			h.sync.Lock()
			close(handler.signal)
			delete(h.list, a)
			h.sync.Unlock()
		}
	} else if b == -1 {
		for _, v := range h.list {
			fmt.Printf("go routine runngin :%d\n", v.name)
		}
	} else {
		//생성하는 로직
		if _, ok := h.list[a]; ok {
			return
		} else {
			log.Println("create go routine")
			h.list[a] = list{make(chan interface{}), a}
		}
		go func() {
			defer log.Println("go routine done")
			for _ = range handlerStream(h.list[a].signal) {
			}
		}()
	}
}

func NewHandler() *Handler {
	return &Handler{
		list: make(map[int]list),
		sync: sync.Mutex{},
	}
}

func main() {

	go func() {
		http.ListenAndServe("localhost:4000", nil)
	}()

	handler := NewHandler()
	reader := bufio.NewReader(os.Stdin)
	for {
		var a, b int
		set := make(chan interface{})
		go func() {
			defer close(set)
			fmt.Fscanln(reader, &a, &b)
			set <- "done"
		}()
		<-set
		log.Println("got input")

		handler.Handle(a, b)

		log.Println("cycle done")
	}

	log.Println("go routine done")
}
