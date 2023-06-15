package mediator

import "fmt"

/**
Components may go in and out of a system at any time
	- Chat room participants
 	- Players in an MMORPG
It makes no sens for them to have direct references to one another
	- Those references may go dead
Solution : have then all refer to some central component that facilitates communication

Mediator
A component that facilitates communication between other components without them
necessarily being aware of each other or having direct access to each other.
*/

type Person struct {
	Name    string
	Room    *ChatRoom
	chatLog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) receive(sender, message string) {
	s := fmt.Sprintf("%s: '%s'", sender, message)
	fmt.Printf("[%s's chat session] %s\n", p.Name, s)
	p.chatLog = append(p.chatLog, s)
}
func (p *Person) say(message string) {
	p.Room.Broadcast(p.Name, message)
}
func (p *Person) privateMessage(who, message string) {
	p.Room.Message(p.Name, who, message)
}

type ChatRoom struct {
	people []*Person
}

func (c *ChatRoom) Broadcast(source, message string) {
	for _, p := range c.people {
		if p.Name != source {
			p.receive(source, message)
		}
	}
}

func (c *ChatRoom) Message(src, dst, msg string) {
	for _, p := range c.people {
		if p.Name == dst {
			p.receive(src, "(private) "+msg)
		}
	}
}

func (c *ChatRoom) Join(p *Person) {
	joinMsg := p.Name + " joins the chat"
	c.Broadcast("Room", joinMsg)
	p.Room = c
	c.people = append(c.people, p)
}

func Start() {
	room := ChatRoom{}
	john := NewPerson("John")
	park := NewPerson("Park")

	room.Join(john)
	room.Join(park)

	john.say("hi room")
	park.say("oh, hey john")

	jane := NewPerson("Jane")
	room.Join(jane)
	jane.say("hi everyone!")

	park.privateMessage("Jane", "glad you could join us!")
}

/**
Create the mediator and have each object in system point to it
Mediator engages in bidirectional communication with its connected components
Mediator has functions the components can call
Components have functions the mediator can call
Evnet processing (e.g. Rx) libraries make communication easier to implement
*/
