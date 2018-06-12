package eg1

import "testing"

func TestAll(t *testing.T) {
	var amp *Amplifier = NewAmplifier("Top-O-Line Amplifier")
	var tuner *Tuner = NewTuner("Top-O-Line AM/FM Tuner", amp)
	var dvd *DvdPlayer = NewDvdPlayer("Top-O-Line DVD Player", amp)
	var cd *CdPlayer = NewCdPlayer("Top-O-Line CD Player", amp)
	var projector *Projector = NewProjector("Top-O-Line Projector", dvd)
	var lights *TheaterLights = NewTheaterLights("Theater Ceiling Lights")
	var screen *Screen = NewScreen("Theater Screen")
	var popper *PopcornPopper = NewPopcornPopper("Popcorn Popper")

	//外观模式使用
	var homeTheater *HomeTheaterFacade = NewHomeTheaterFacade(amp, tuner, dvd, cd, projector, screen, lights, popper)
	homeTheater.WatchMovie("Raiders of the Lost Ark")
	homeTheater.EndMovie()
}