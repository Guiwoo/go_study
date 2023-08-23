package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var (
	global_one = time.NewTicker(1 * time.Second)
	global_two = time.NewTicker(2 * time.Second)
	chan_map   = make(map[int]chan interface{})
	sc         = &sync.Mutex{}
)

func streamGen(done <-chan interface{}) <-chan interface{} {
	stream := make(chan interface{})
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case <-global_one.C:
				stream <- "ticker sent data"
			}
		}
	}()
	return stream
}

func main() {

	go func() {
		http.ListenAndServe("localhost:4000", nil)
	}()

	go func() {
		defer fmt.Println("함수종료 왜 ?")
		for {
			target := int(time.Now().Unix() % 2)
			select {
			case <-global_two.C:
				if target%2 == 0 {
					if _, ok := chan_map[target]; !ok {
						// 채널 생성해서 넣어주고
						chan_map[target] = make(chan interface{})
					}
					for data := range streamGen(chan_map[target]) {
						log.Println(data)
					}
				} else {
					sc.Lock()
					close(chan_map[target])
					delete(chan_map, target)
					sc.Unlock()
				}
			}
		}
	}()

	fmt.Println("finish as normal exit")
}
