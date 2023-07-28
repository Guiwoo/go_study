package facade

import "testing"

func TestFacade(t *testing.T) {
	homeTheater := NewHomeTheaterFacade(
		NewAmplifier(),
		NewTuner(),
		NewStreamingPlayer(),
		NewProjector(),
		NewTheaterLights(),
		NewScreen(),
		NewPopcornPopper(),
	)

	homeTheater.watchMovie("Inception")
}
