package main

import (
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	for {
		go func() {
			arr := make([]int, 0)
			for cap(arr) < 1_000_000 {
				arr = append(arr, rand.Intn(100))
				time.Sleep(500 * time.Millisecond)
			}
		}()
		time.Sleep(time.Second * 500)
	}
}
