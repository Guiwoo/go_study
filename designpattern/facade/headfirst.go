package facade

import "fmt"

/**
다른 이유로 어댑터를 변경하는 방법
인터페이스를 단순하게 변가ㅕㅇ하기 위해 사용
"퍼사드 패턴"
*/

type Screen struct{}

func (s *Screen) up() {
	fmt.Println("[Screen] up")
}
func (s *Screen) down() {
	fmt.Println("[Screen] down")
}
func NewScreen() *Screen {
	return &Screen{}
}

type PopcornPopper struct{}

func (p *PopcornPopper) on() {
	fmt.Println("[PopcornPopper] on")
}
func (p *PopcornPopper) off() {
	fmt.Println("[PopcornPopper] off")
}
func (p *PopcornPopper) pop() {
	fmt.Println("[PopcornPopper] pop")
}
func NewPopcornPopper() *PopcornPopper {
	return &PopcornPopper{}
}

type TheaterLights struct{}

func (t *TheaterLights) on() {
	fmt.Println("[TheaterLights] on")
}
func (t *TheaterLights) off() {
	fmt.Println("[TheaterLights] off")
}
func (t *TheaterLights) dim() {
	fmt.Println("[TheaterLights] dim")
}
func NewTheaterLights() *TheaterLights {
	return &TheaterLights{}
}

type StreamingPlayer struct{}

func (s *StreamingPlayer) on() {
	fmt.Println("[StreamingPlayer] on")
}
func (s *StreamingPlayer) off() {
	fmt.Printf("[StreamingPlayer] off\n")
}
func (s *StreamingPlayer) pause() {
	fmt.Printf("[StreamingPlayer] pause\n")
}
func (s *StreamingPlayer) play(target string) {
	fmt.Printf("[StreamingPlayer %s] play\n", target)
}
func (s *StreamingPlayer) setSurroundAudio() {
	fmt.Printf("[StreamingPlayer] setSurroundAudio\n")
}
func (s *StreamingPlayer) setTwoChannelAudio() {
	fmt.Printf("[StreamingPlayer] setTwoChannelAudio\n")
}
func (s *StreamingPlayer) stop() {
	fmt.Printf("[StreamingPlayer] stop\n")
}
func (s *StreamingPlayer) String() string {
	return "StreamingPlayer"
}
func NewStreamingPlayer() *StreamingPlayer {
	return &StreamingPlayer{}
}

type Projector struct {
	player StreamingPlayer
}

func (p *Projector) on() {
	fmt.Printf("[%s Projector] on\n", p.player)
}
func (p *Projector) off() {
	fmt.Printf("[%s Projector] off\n", p.player)
}
func (p *Projector) tvMode() {
	fmt.Printf("[%s Projector] tvMode\n", p.player)
}
func (p *Projector) wideScreenMode() {
	fmt.Printf("[%s Projector] wideScreenMode\n", p.player)
}
func NewProjector() *Projector {
	return &Projector{}
}

type Tuner struct{}

func (t *Tuner) on() {
	fmt.Println("[Tuner] on")
}
func (t *Tuner) off() {
	fmt.Println("[Tuner] off")
}
func (t *Tuner) setAm() {
	fmt.Println("[Tuner] setAm")
}
func (t *Tuner) setFm() {
	fmt.Println("[Tuner] setFm")
}
func (t *Tuner) setFrequency() {
	fmt.Println("[Tuner] setFrequency")
}
func (t *Tuner) String() string {
	return "Tuner"
}
func NewTuner() *Tuner {
	return &Tuner{}
}

type Amplifier struct {
	tuner  *Tuner
	player *StreamingPlayer
}

func (a *Amplifier) on() {
	fmt.Printf("[%s][%s][Amplifier] on\n", a.tuner, a.player)
}
func (a *Amplifier) off() {
	fmt.Printf("[%s][%s][Amplifier] off\n", a.tuner, a.player)
}
func (a *Amplifier) setStreamingPlayer() {
	fmt.Printf("[%s][%s][Amplifier] setStreamingPlayer\n", a.tuner, a.player)
}
func (a *Amplifier) setSurroundSound() {
	fmt.Printf("[%s][%s][Amplifier] setSurroundSound\n", a.tuner, a.player)
}
func (a *Amplifier) setVolume() {
	fmt.Printf("[%s][%s][Amplifier] setVolume\n", a.tuner, a.player)
}
func (a *Amplifier) String() string {
	return fmt.Sprintf("[%s][%s][Amplifier]", a.tuner, a.player)
}
func NewAmplifier() *Amplifier {
	return &Amplifier{}
}

type HomeTheaterFacade struct {
	amp           *Amplifier
	tuner         *Tuner
	player        *StreamingPlayer
	projector     *Projector
	lights        *TheaterLights
	screen        *Screen
	popcornPopper *PopcornPopper
}

func (h *HomeTheaterFacade) watchMovie(movie string) {
	fmt.Printf("Get ready to watch a movie...\n")
	h.popcornPopper.on()
	h.popcornPopper.pop()
	h.lights.dim()
	h.screen.down()
	h.projector.on()
	h.projector.wideScreenMode()
	h.amp.on()
	h.amp.setStreamingPlayer()
	h.amp.setSurroundSound()
	h.amp.setVolume()
	h.player.on()
	h.player.play(movie)
}

func NewHomeTheaterFacade(
	amp *Amplifier,
	tuner *Tuner,
	player *StreamingPlayer,
	projector *Projector,
	lights *TheaterLights,
	screen *Screen,
	popcornPopper *PopcornPopper,
) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		amp:           amp,
		tuner:         tuner,
		player:        player,
		projector:     projector,
		lights:        lights,
		screen:        screen,
		popcornPopper: popcornPopper,
	}
}
