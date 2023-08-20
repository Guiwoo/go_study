package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type address struct {
	streetAddress, city, country string
}

type person struct {
	name    string
	address *address
	friends []string
}

func deepCopyingProblem() {
	john := &person{
		name:    "John",
		address: &address{"123 london rd", "london", "uik"},
	}

	//jane := john
	//jane.name = "Jane"
	//jane.address.streetAddress = "321 Baker street"

	jane := john
	jane.address = &address{
		john.address.streetAddress,
		john.address.city,
		john.address.country,
	}

	jane.address.streetAddress = "321 Baker street"

	// deep copying when you copying object , pointer ,slice copy the original object
	// it doesn't affect for the copy the object it might occur the problem to pointer copies
}

// DeepCopy function make a receiver function
func (a *address) DeepCopy() *address {
	return &address{}
}

// slice ,string, map still needs copy triple and double check the type not idle way
// one possible of approach to handle deep copy

func (p *person) DeepCopy() *person {
	q := p
	q.address = p.address.DeepCopy()
	copy(q.friends, p.friends)
	return q
}

func (p *person) DeepCopyUtil() *person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := &person{}
	_ = d.Decode(result)

	return result
}

/**
Prototype factory
*/

type employee struct {
	name   string
	office address
}

func (p *employee) DeepCopy() *employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := &employee{}
	_ = d.Decode(result)

	return result
}

var mainOffice = employee{"", address{"0", "123 East Dr", "London"}}
var auxOffice = employee{"", address{"0", "66 West Dr", "London"}}

func newEmployee(proto *employee, name, city string) *employee {
	result := proto.DeepCopy()
	result.name = name
	result.office.city = city
	return result
}

func newMainOfficeEmployee(name, city string) *employee {
	return newEmployee(&mainOffice, name, city)
}

func newAuxOfficeEmployee(name, city string) *employee {
	return newEmployee(&auxOffice, name, city)
}

type Keyboard struct {
	Layout string
	Switch []string
	KeyCap string
}

func (k *Keyboard) Clone() Cloneable {
	cloneSlice := make([]string, len(k.Switch))
	copy(cloneSlice, k.Switch)
	return &Keyboard{
		Layout: k.Layout,
		Switch: cloneSlice,
		KeyCap: k.KeyCap,
	}

}

type Cloneable interface {
	Clone() Cloneable
}

func deepCopyProblem(k *Keyboard) *Keyboard {

	var leopold *Keyboard

	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)
	_ = encoder.Encode(k)

	decoder := gob.NewDecoder(&buf)
	_ = decoder.Decode(&leopold)

	return leopold
}

func Run() {
	sixty := &Keyboard{"60%", []string{"Cherry MX Blue", "Cherry MX Brown", "Cherry MX Red"}, "DSA"}

	keyboardClone := sixty.Clone()

	fmt.Println(keyboardClone)
}
