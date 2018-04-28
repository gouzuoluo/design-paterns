package facade

import "fmt"

type HomeTheaterFacade struct {
	amp       *Amplifier
	tuner     *Tuner
	dvd       *DvdPlayer
	cd        *CdPlayer
	projector *Projector
	lights    *TheaterLights
	screen    *Screen
	popper    *PopcornPopper
}

func NewHomeTheaterFacade(amp *Amplifier,
tuner *Tuner,
dvd *DvdPlayer,
cd *CdPlayer,
projector *Projector,
screen *Screen,
lights *TheaterLights,
popper *PopcornPopper) *HomeTheaterFacade {
	this := new(HomeTheaterFacade)
	this.amp = amp
	this.tuner = tuner
	this.dvd = dvd
	this.cd = cd
	this.projector = projector
	this.screen = screen
	this.lights = lights
	this.popper = popper
	return this
}

func (this *HomeTheaterFacade)WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	this.popper.On()
	this.popper.Pop()
	this.lights.Dim(10)
	this.screen.Down()
	this.projector.On()
	this.projector.WideScreenMode()
	this.amp.On()
	this.amp.SetDvd(this.dvd)
	this.amp.SetSurroundSound()
	this.amp.SetVolume(5)
	this.dvd.On()
	this.dvd.Play(movie)
}

func (this *HomeTheaterFacade)EndMovie() {
	fmt.Println("Shutting movie theater down...")
	this.popper.Off()
	this.lights.On()
	this.screen.Up()
	this.projector.Off()
	this.amp.Off()
	this.dvd.Stop()
	this.dvd.Eject()
	this.dvd.Off()
}

func (this *HomeTheaterFacade)ListenToCd(cdTitle string) {
	fmt.Println("Get ready for an audiopile experence...")
	this.lights.On()
	this.amp.On()
	this.amp.SetVolume(5)
	this.amp.SetCd(this.cd)
	this.amp.SetStereoSound()
	this.cd.On()
	this.cd.Play(cdTitle)
}

func (this *HomeTheaterFacade)EndCd() {
	fmt.Println("Shutting down CD...")
	this.amp.Off()
	this.amp.SetCd(this.cd)
	this.cd.Eject()
	this.cd.Off()
}

func (this *HomeTheaterFacade)ListenToRadio(frequency float64) {
	fmt.Println("Tuning in the airwaves...")
	this.tuner.On()
	this.tuner.SetFrequency(frequency)
	this.amp.On()
	this.amp.SetVolume(5)
	this.amp.SetTuner(this.tuner)
}

func (this *HomeTheaterFacade)EndRadio() {
	fmt.Println("Shutting down the tuner...")
	this.tuner.Off()
	this.amp.Off()
}