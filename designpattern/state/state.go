package state

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Consider an ordinary telephone
What you do with it depends on the state of the phone/line
	- If it's ringing or you want to make a call, you can pick it up
	- Phone must be off the hook to talk/make a call
	- If you try calling someone, and it's busy,you put the handset down

Changes in state can be explicit or in response to event(Observer pattern)

A pattern in which the object's behavior is determined by its state.
An object transactions from one state to another ( something needs to trigger a transition).

A formalized construct which manages state and transitions is called a state machine.
*/

type Switch struct {
	State State
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

type State interface {
	On(s *Switch)
	Off(s *Switch)
}

type BaseState struct {
}

func (b *BaseState) On(s *Switch) {
	fmt.Println("Light is already on")
}

func (b *BaseState) Off(s *Switch) {
	fmt.Println("Light is already off")
}

var _ State = (*BaseState)(nil)

type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{BaseState{}}
}

func (o *OnState) Off(s *Switch) {
	fmt.Println("Turning off the light")
	s.State = NewOffState()
}

type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{BaseState{}}
}

func (o *OffState) On(s *Switch) {
	fmt.Println("Turning on light")
	s.State = NewOnState()
}

func NewSwitch() *Switch {
	return &Switch{NewOnState()}
}

func Ex01() {
	sw := NewSwitch()

	sw.On()
	sw.Off()
	sw.Off()
}

type State2 int

const (
	OffHook State2 = iota
	Connecting
	Connected
	OnHold
	OnHook
)

func (s State2) String() string {
	switch s {
	case OffHook:
		return "Off Hook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connencted"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	}
	return "Unknown"
}

type Trigger int

const (
	CallDialed Trigger = iota
	HungUp
	CallConnected
	PlacedOnHold
	TakenOffHold
	LeftMessage
)

func (t Trigger) String() string {
	switch t {
	case CallDialed:
		return "CallDialed"
	case HungUp:
		return "HungUp"
	case CallConnected:
		return "CallConnected"
	case PlacedOnHold:
		return "PlacedOnHold"
	case TakenOffHold:
		return "TakenOffHold"
	case LeftMessage:
		return "LeftMessage"
	}
	return "Unknown"
}

type TriggerResult struct {
	Trigger Trigger
	State   State2
}

var rules = map[State2][]TriggerResult{
	OffHook: {
		{CallDialed, Connecting},
	},
	Connecting: {
		{HungUp, OffHook},
		{CallConnected, Connected},
	},
	Connected: {
		{LeftMessage, OffHook},
		{HungUp, OffHook},
		{PlacedOnHold, OnHold},
	},
	OnHold: {
		{TakenOffHold, Connected},
		{HungUp, OffHook},
	},
}

func Ex02() {
	state, exitState := OffHook, OnHook

	for ok := true; ok; ok = state != exitState {
		fmt.Println("The phone is currently", state)
		fmt.Println("Select a trigger:")

		for i := 0; i < len(rules[state]); i++ {
			tr := rules[state][i]
			fmt.Println(strconv.Itoa(i), ".", tr.Trigger)
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		i, _ := strconv.Atoi(string(input))

		tr := rules[state][i]
		state = tr.State
	}

	fmt.Println("We are done using the phone")
}

type State3 int

const (
	Locked State3 = iota
	Failed
	Unlocked
)

func Start() {
	code := "1234"
	state := Locked
	entry := strings.Builder{}

	for {
		switch state {
		case Locked:
			r, _, _ := bufio.NewReader(os.Stdin).ReadRune()
			entry.WriteRune(r)

			if entry.String() == code {
				state = Unlocked
				break
			}

			if strings.Index(code, entry.String()) != 0 {
				state = Failed
			}
		case Failed:
			fmt.Println("FAILED")
			entry.Reset()
			state = Locked
		case Unlocked:
			fmt.Println("UNLOCKED")
			return
		}
	}
}

/**
Given sufficient complexity, it pays to formally define possible states and events/triggers
Can define
	- State entry/exit behaviors
	- Action when a particular event causes a transition
	- Guard conditions enabling/disabling a transition
	- Default action when no transitions are found for an event
*/
