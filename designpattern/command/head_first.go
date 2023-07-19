package command

import "fmt"

/**
요청하는 쪽과 그 작업을 처리하는 쪽을 분리할수 있음
리모컨에서 작업을 요청 하면 업체에서 제공한 클래스가 그 작업을 처리한다고 보면 이 패턴을 적용할 수 있지 않을까요?

커맨드 객체를 추가 => 측정 작업요청을 캡슐화해 줌 , 버트마다 커맨드 객체를 저장해, 사용자가 버튼을 눌렀을때 커맨드 객체로 작업을 처리,

음식의 주문과정
1. 고객이 종업원에게 주문을 합니다. => createOrder()
2. 종업원은 문을 받아서 카운터에 전달하고 이야기해준다. takeOrder() orderUp()
3. 주방장이 주문대로 음식을 조리한다. makeBurger(), makeShake()

주문서는 주문 내용을 캡슐화 한다.
- 주문서는 주문 내용을 요구하는 객체 식사 준비에 필요한 행동을 캡슐화 한 메소드 가 있다.
- 주문서를 받고 그저 orderUp() 을 호출 해서 식사준비를 해줌
- takeOrder() 여러고객의 주문서를 매개변수로 전달,
- 주방장은 식사를 준비하는ㄴ 데 필요한 정보를 가지고 있다.

어떤것을 요구 하는 객체와 , 그것을 받아들이고 처리하는 개체의 분리
리모컨 API 에 대입해 본다면 ? 리모컨 버튼이 눌렸을 때 호출되는 코드와 실제로 일을 처리하는 코드를 분리해야 합니다.
리모컨 슬롯에 객체마을 식당의 문서 같은 객체가 들어있다면 어떨까요 ? 버튼이 눌렸을때 orderUP() 같은 메소드가 호출되면서,

*/

type CommandHead interface {
	execute()
	undo()
}

type Light struct{}

func (l *Light) on() {
	fmt.Println("전구가 켜집니다.")
}
func (l *Light) off() {
	fmt.Println("전구가 꺼집니다.")
}

type GarageDoor struct{}

func (g *GarageDoor) up() {
	fmt.Println("차고 문 올라가요")
}
func (g *GarageDoor) down() {
	fmt.Println("차고 문 내려가요")
}
func (g *GarageDoor) stop() {
	fmt.Println("차고 문 멈춥니다.")
}
func (g *GarageDoor) lightOn() {
	fmt.Println("차고 에 불이 켜집니다.")
}
func (g *GarageDoor) lightOff() {
	fmt.Println("차고 에 불이 꺼집니다.")
}

type GarageOnCommand struct {
	g *GarageDoor
}

func (g *GarageOnCommand) execute() {
	g.g.lightOn()
	g.g.up()
	g.g.stop()
}
func (g *GarageOnCommand) undo() {
	g.g.lightOff()
	g.g.down()
}
func (g *GarageOnCommand) String() string {
	return "차고 ON 명령"
}

type GarageDoorOffCommand struct {
	g *GarageDoor
}

func (g *GarageDoorOffCommand) execute() {
	g.g.lightOff()
	g.g.down()
}
func (g *GarageDoorOffCommand) undo() {
	g.g.lightOn()
	g.g.up()
	g.g.stop()
}

type LightOnCommand struct {
	light Light
}

func (l *LightOnCommand) execute() {
	l.light.on()
}
func (l *LightOnCommand) undo() {
	l.light.off()
}

type SimpleController struct {
	c CommandHead
}

func (s *SimpleController) SetCommand(c CommandHead) {
	s.c = c
}

func (s *SimpleController) ButtonWasPressed() {
	s.c.execute()
}

/**
커맨드 패턴을 사용하면 요청 내역을 객체로 캡슐화해서 객체를 서로 다른 요청내역에 따라 매개변수화할 수 있습니다.
이러면 요청을 큐에 저장하거나 로그로 기록하건ㅏ 작업ㅊ 쉬초 기능을 사용할수 있습니다.

특정 행동과 리비서를 한객체에 넣고,execute 라는 메소드를 호출하면 해당 요청이 처리된다.

*/

type NoCommand struct{}

func (n *NoCommand) execute() {
	fmt.Println("There's no command")
}
func (n *NoCommand) undo() {
	fmt.Println("There's no command")
}

type RemoteController struct {
	OnCommands  []CommandHead
	OffCommands []CommandHead
	undoCommand CommandHead
}

func (r *RemoteController) SetCommand(slot int, onCmd, offCmd CommandHead) {
	r.OnCommands[slot] = onCmd
	r.OffCommands[slot] = offCmd
}
func (r *RemoteController) OnButtonWasPushed(slot int) {
	r.OnCommands[slot].execute()
	r.undoCommand = r.OnCommands[slot]
}
func (r *RemoteController) OffButtonWasPushed(slot int) {
	r.OffCommands[slot].execute()
	r.undoCommand = r.OffCommands[slot]
}
func (r *RemoteController) undoButtonWasPushed() {
	r.undoCommand.undo()
}

func NewRemoteController() *RemoteController {
	size := 7
	onCmd := make([]CommandHead, size)
	offCmd := make([]CommandHead, size)
	for i := range onCmd {
		onCmd[i] = &NoCommand{}
		offCmd[i] = &NoCommand{}
	}
	return &RemoteController{onCmd, offCmd, &NoCommand{}}
}

type LightOffCommand struct {
	light Light
}

func (l *LightOffCommand) execute() {
	l.light.off()
}
func (l *LightOffCommand) undo() {
	l.light.on()
}

type Stereo struct {
	cd, dvd string
	volume  int
}

func (s *Stereo) On() {
	fmt.Println("오디오 를 켭니다.")
}
func (s *Stereo) Off() {
	fmt.Println("오디오 를 끕니다.")
}
func (s *Stereo) SetCd(cd string) {
	s.cd = cd
	fmt.Printf("cd 를 녛습니다 : %s\n", s.cd)
}
func (s *Stereo) SetDvd(dvd string) {
	s.dvd = dvd
	fmt.Printf("dvd 를 넣습니다 : %s\n", s.dvd)
}
func (s *Stereo) SetVolume(v int) {
	s.volume = v
	fmt.Printf("볼륨을 설정합니다 : %d\n", s.volume)
}

type StereoOnWithCDCommand struct {
	stereo Stereo
}

func (s *StereoOnWithCDCommand) execute() {
	s.stereo.On()
	s.stereo.SetCd(s.stereo.cd)
	s.stereo.SetVolume(11)
}
func (s *StereoOnWithCDCommand) undo() {
	s.stereo.Off()
}
func NewStereoOnWithCDCommand(cd string) *StereoOnWithCDCommand {
	return &StereoOnWithCDCommand{
		Stereo{
			cd: cd,
		},
	}
}

type StereoOffWithCDCommand struct {
	stereo Stereo
}

func (s *StereoOffWithCDCommand) execute() {
	s.stereo.Off()
}
func (s *StereoOffWithCDCommand) undo() {
	s.stereo.On()
	s.stereo.SetCd(s.stereo.cd)
	s.stereo.SetVolume(11)
}

const (
	Off = iota
	Low
	Medium
	High
)

type CeilingFan struct {
	Speed    int
	Location string
}

func (c *CeilingFan) high() {
	c.Speed = High
}
func (c *CeilingFan) medium() {
	c.Speed = Medium
}
func (c *CeilingFan) low() {
	c.Speed = Low
}
func (c *CeilingFan) off() {
	c.Speed = Off
}
func (c *CeilingFan) getSpeed() int {
	return c.Speed
}

func NewCeilingFan(location string) *CeilingFan {
	return &CeilingFan{Off, location}
}

type CeilingFanHighCommand struct {
	fan       *CeilingFan
	prevSpeed int
}

func (c *CeilingFanHighCommand) execute() {
	c.prevSpeed = c.fan.getSpeed()
	c.fan.high()
}
func (c *CeilingFanHighCommand) undo() {
	switch c.prevSpeed {
	case High:
		c.fan.high()
	case Medium:
		c.fan.medium()
	case Low:
		c.fan.low()
	default:
		c.fan.off()
	}
}

type CeilingFanMediumCommand struct {
	fan       *CeilingFan
	prevSpeed int
}

func (c *CeilingFanMediumCommand) execute() {
	c.prevSpeed = c.fan.getSpeed()
	c.fan.medium()
}
func (c *CeilingFanMediumCommand) undo() {
	switch c.prevSpeed {
	case High:
		c.fan.high()
	case Medium:
		c.fan.medium()
	case Low:
		c.fan.low()
	default:
		c.fan.off()
	}
}

type CeilingFanOffCommand struct {
	fan       *CeilingFan
	prevSpeed int
}

func (c *CeilingFanOffCommand) execute() {
	c.prevSpeed = c.fan.getSpeed()
	c.fan.off()
}
func (c *CeilingFanOffCommand) undo() {
	switch c.prevSpeed {
	case High:
		c.fan.high()
	case Medium:
		c.fan.medium()
	case Low:
		c.fan.low()
	default:
		c.fan.off()
	}
}

type MacroCommand struct {
	cmd []CommandHead
}

func (m *MacroCommand) execute() {
	for _, v := range m.cmd {
		v.execute()
	}
}
func (m *MacroCommand) undo() {}
