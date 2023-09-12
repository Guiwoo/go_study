package main

import (
	"github.com/nats-io/nats.go"
)

type NatsClient struct {
	*nats.EncodedConn
	conn *nats.EncodedConn
}

func (n *NatsClient) subscribe(topic string, handler nats.Handler) {
	n.conn.Subscribe(topic, handler)
}

func NewNatsClient() *NatsClient {
	consumer := &NatsClient{}

	c, err := nats.Connect("localhost:4222")
	if err != nil {
		panic(err)
	}

	consumer.conn, _ = nats.NewEncodedConn(c, nats.JSON_ENCODER)

	return consumer
}
