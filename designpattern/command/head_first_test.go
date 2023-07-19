package command

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	remote := &SimpleController{&LightOnCommand{Light{}}}

	remote.ButtonWasPressed()

	garage := &GarageOnCommand{&GarageDoor{}}
	remote.SetCommand(garage)

	remote.ButtonWasPressed()
}

func Test02(t *testing.T) {
	rmt := NewRemoteController()

	light := &Light{}
	garage := &GarageDoor{}
	stereo := &Stereo{}

	rmt.SetCommand(0, &LightOnCommand{*light}, &LightOffCommand{*light})
	rmt.SetCommand(1, &GarageOnCommand{garage}, &GarageDoorOffCommand{garage})
	rmt.SetCommand(2, NewStereoOnWithCDCommand("guiwoo"), &StereoOffWithCDCommand{*stereo})

	rmt.OnButtonWasPushed(1)
	rmt.OffButtonWasPushed(1)
}

func Test03(t *testing.T) {
	rmt := NewRemoteController()

	light := &Light{}
	garage := &GarageDoor{}
	stereo := &Stereo{}

	rmt.SetCommand(0, &LightOnCommand{*light}, &LightOffCommand{*light})
	rmt.SetCommand(1, &GarageOnCommand{garage}, &GarageDoorOffCommand{garage})
	rmt.SetCommand(2, NewStereoOnWithCDCommand("guiwoo"), &StereoOffWithCDCommand{*stereo})

	rmt.OnButtonWasPushed(1)

	rmt.undoButtonWasPushed()

	fmt.Println(rmt.undoCommand)
}

func Test04(t *testing.T) {
	rmt := NewRemoteController()

	fan := NewCeilingFan("Living Room")

	rmt.SetCommand(0, &CeilingFanMediumCommand{fan: fan}, &CeilingFanOffCommand{fan: fan})
	rmt.SetCommand(1, &CeilingFanHighCommand{fan: fan}, &CeilingFanOffCommand{fan: fan})
}

func Test05_macro(t *testing.T) {
	light := &Light{}
	grage := &GarageDoor{}
	stereo := &Stereo{}

	//on command 작성
	lightOn := &LightOnCommand{*light}
	grageOn := &GarageOnCommand{grage}
	stereoOn := NewStereoOnWithCDCommand("guiwoo")
	onCmd := []CommandHead{lightOn, grageOn, stereoOn}

	//off command 작성
	lightOff := &LightOffCommand{*light}
	grageOff := &GarageDoorOffCommand{grage}
	stereoOff := &StereoOffWithCDCommand{*stereo}
	offCmd := []CommandHead{lightOff, grageOff, stereoOff}

	//slot 에 넣어주기
	rmt := NewRemoteController()
	rmt.SetCommand(0, &MacroCommand{onCmd}, &MacroCommand{offCmd})

	rmt.OnButtonWasPushed(0)
}
