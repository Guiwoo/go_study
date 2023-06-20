package observer

import (
	"container/list"
	"fmt"
	"time"
)

/**
We need to be informed when certain things happen
	- Objects field changes
	- Object does something
	- Some external event occurs

We want to listen to events and notified when they occur
Two participants : observable and observer

An observer is an object that wishes to be informed about events happening in the system.
Then entity generating the events is and observable.
*/

// Observable, Observer

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) UnSubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type Person struct {
	Observable
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		Name:       name,
	}
}

func (p *Person) CatchCold() {
	p.Fire(p.Name)
}

type DoctorService struct{}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s", data.(string))
}

type PropertyChange struct {
	Name  string
	Value interface{}
}

type Person2 struct {
	Observable
	age int
}

func NewPerson2(age int) *Person2 {
	return &Person2{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

func (p *Person2) Age() int {
	return p.age
}

func (p *Person2) SetAge(age int) {
	if age == p.age {
		return
	}

	prev := p.CanVote()
	p.age = age

	p.Fire(PropertyChange{"Age", p.age})

	if prev != p.CanVote() {
		p.Fire(PropertyChange{"CanVote", p.CanVote()})
	}
}

func (p *Person2) CanVote() bool {
	return p.age >= 18
}

type TrafficManagement struct {
	o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 18 {
			fmt.Println("You can drive now!")
			t.o.UnSubscribe(t)
		}
	}
}

type ElectoralRoll struct {
}

func (e *ElectoralRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can vote!")
		}
	}
}

func Start() {
	p := NewPerson2(0)
	er := &ElectoralRoll{}
	p.Subscribe(er)

	for i := 10; i < 20; i++ {
		fmt.Println("Setting Age to ", i)
		p.SetAge(i)
		time.Sleep(1 * time.Second)
	}
}

func ex02() {
	p := NewPerson2(15)
	t := &TrafficManagement{p.Observable}
	p.Subscribe(t)

	for i := 16; i <= 20; i++ {
		fmt.Println("Setting the age to ", i)
		p.SetAge(i)
	}
}
func ex01() {
	p := NewPerson("Guiwoo")

	ds := &DoctorService{}

	p.Subscribe(ds)

	p.CatchCold()
}

/**
- Observer is an intrusive approach
- Must provide a way of clients to subscribe
- Event data sent from observable to all observers
- Data represented as interface{}
*/
