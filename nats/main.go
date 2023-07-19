package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"sync"
	"time"
)

func synchronous(nc *nats.Conn) {
	sub, err := nc.SubscribeSync("updates")

	msg, err := sub.NextMsg(10 * time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Reply: ", string(msg.Data))
}

func asynchronous(nc *nats.Conn) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	if _, err := nc.Subscribe("updates", func(msg *nats.Msg) {
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	synchronous(nc)
	guiwoo.SomethingOnYourMind()
}
