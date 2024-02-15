package main

import (
	"fmt"
	"math/rand"
	"testing"
)

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Locker
	Mover
}

type bike struct{}

func (*bike) Move() {
	fmt.Println("bike is moving")
}
func (*bike) Lock() {
	fmt.Println("Locking the bike")
}
func (*bike) Unlock() {
	fmt.Println("Unlocking the bike")
}

func Test_Bike(t *testing.T) {
	var (
		ml MoveLocker
		m  Mover
	)
	ml = &bike{}
	m = ml

	b, ok := m.(*bike)
	fmt.Println("Does m has value of bike ? ", ok)

	ml = b
}

type car struct{}

func (car) String() string {
	return "Vroom!"
}

type cloud struct{}

func (cloud) String() string {
	return "Big data!"
}

func TestRunTimeAssertion(t *testing.T) {
	mvs := []fmt.Stringer{car{}, cloud{}}
	for i := 0; i < 10; i++ {
		n := rand.Intn(2)

		if v, ok := mvs[n].(cloud); ok {
			fmt.Println("Got lucky :", v)
			continue
		}
		fmt.Println("got unlucky")
	}
}

/*
*
Server
*/
type PubSub struct {
	host string
}

func New(host string) *PubSub {
	return &PubSub{
		host,
	}
}

func (ps *PubSub) Publish(key string, v interface{}) error {
	fmt.Println("Actual PubSub: Publish")
	return nil
}
func (ps *PubSub) Subscribe(key string) error {
	fmt.Println("Actual PubSub: Subscribe")
	return nil
}

/**
Client
*/

type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

type mock struct{}

func (m *mock) Publish(key string, v interface{}) error {
	fmt.Println("Mock PubSub: Publish")
	return nil
}
func (m *mock) Subscribe(key string) error {
	fmt.Println("Mock Subscribe: subscribe")
	return nil
}

var _ publisher = (*mock)(nil)

func Test_Mocking(t *testing.T) {
	pubs := []publisher{
		&mock{},
		New("localhost"),
	}

	for _, p := range pubs {
		p.Publish("key", "value")
		p.Subscribe("key")
	}
}
