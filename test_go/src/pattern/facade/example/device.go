package facade

import "fmt"

//调谐器
type Tuner struct {
	description string
	amplifier   *Amplifier
	frequency   float64 //频率
}

func NewTuner(description string, amplifier *Amplifier) *Tuner {
	this := new(Tuner)
	this.description = description
	this.amplifier = amplifier
	return this
}

func (this *Tuner)On() {
	fmt.Println(this.description + " on")
}

func (this *Tuner)Off() {
	fmt.Println(this.description + " off")
}

func (this *Tuner)SetFrequency(frequency float64) {
	fmt.Println(this.description + " setting frequency to " + fmt.Sprintf("%f",frequency))
	this.frequency = frequency
}

func (this *Tuner)SetAm() {
	fmt.Println(this.description + " setting AM mode")
}

func (this *Tuner)SetFm() {
	fmt.Println(this.description + " setting FM mode")
}

func (this *Tuner)String() string {
	return this.description
}

//DVD播放器
type DvdPlayer struct {
	description  string
	currentTrack int
	amplifier    *Amplifier
	movie        string
}

func NewDvdPlayer(description string, amplifier *Amplifier) *DvdPlayer {
	this := new(DvdPlayer)
	this.description = description
	this.amplifier = amplifier
	return this
}

func (this *DvdPlayer)On() {
	fmt.Println(this.description + " on")
}

func (this *DvdPlayer)Off() {
	fmt.Println(this.description + " off")
}

func (this *DvdPlayer)Eject() {
	this.movie = ""
	fmt.Println(this.description + " eject")
}

func (this *DvdPlayer)Play(movie string) {
	this.movie = movie
	this.currentTrack = 0
	fmt.Println(this.description + " playing \"" + movie + "\"")
}

func (this *DvdPlayer)Play2(track int) {
	if this.movie == "" {
		fmt.Println(this.description + " can't play track " + string(track) + " no dvd inserted")
	} else {
		this.currentTrack = track
		fmt.Println(this.description + " playing track " + string(this.currentTrack) + " of \"" + this.movie + "\"")
	}
}

func (this *DvdPlayer)Stop() {
	this.currentTrack = 0
	fmt.Println(this.description + " stopped \"" + this.movie + "\"")
}

func (this *DvdPlayer)Pause() {
	fmt.Println(this.description + " paused \"" + this.movie + "\"")
}

func (this *DvdPlayer)SetTwoChannelAudio() {
	fmt.Println(this.description + " set two channel audio")
}

func (this *DvdPlayer)SetSurroundAudio() {
	fmt.Println(this.description + " set surround audio")
}

func (this *DvdPlayer)String() string {
	return this.description
}

//CD机
type CdPlayer struct {
	description  string
	currentTrack int
	amplifier    *Amplifier
	title        string
}

func NewCdPlayer(description string, amplifier *Amplifier) *CdPlayer {
	this := new(CdPlayer)
	this.description = description
	this.amplifier = amplifier
	return this
}

func (this *CdPlayer)On() {
	fmt.Println(this.description + " on")
}

func (this *CdPlayer)Off() {
	fmt.Println(this.description + " off")
}

func (this *CdPlayer)Eject() {
	this.title = ""
	fmt.Println(this.description + " eject")
}

func (this *CdPlayer)Play(title string) {
	this.title = title
	this.currentTrack = 0
	fmt.Println(this.description + " playing \"" + title + "\"")
}

func (this *CdPlayer)Play2(track int) {
	if this.title == "" {
		fmt.Println(this.description + " can't play track " + string(this.currentTrack) + " no cd inserted")
	} else {
		this.currentTrack = track
		fmt.Println(this.description + " playing track " + string(this.currentTrack))
	}
}

func (this *CdPlayer)Stop() {
	this.currentTrack = 0
	fmt.Println(this.description + " stopped ")
}

func (this *CdPlayer)Pause() {
	fmt.Println(this.description + " paused \"" + this.title + "\"")
}

func (this *CdPlayer)String() string {
	return this.description
}

//功放
type Amplifier struct {
	description string
	tuner       *Tuner     //调谐器
	dvd         *DvdPlayer //DVD
	cd          *CdPlayer  //CD
}

func NewAmplifier(description string) *Amplifier {
	this := new(Amplifier)
	this.description = description
	return this
}

func (this *Amplifier)On() {
	fmt.Println(this.description + " on")
}

func (this *Amplifier)Off() {
	fmt.Println(this.description + " off")
}

func (this *Amplifier)SetStereoSound() {
	fmt.Println(this.description + " stereo mode on")
}

func (this *Amplifier)SetSurroundSound() {
	fmt.Println(this.description + " surround sound on (5 speakers, 1 subwoofer)")
}

func (this *Amplifier)SetVolume(level int) {
	fmt.Println(this.description + " setting volume to " + string(level))
}

func (this *Amplifier)SetTuner(tuner *Tuner) {
	fmt.Println(this.description + " setting tuner to " + tuner.String())
	this.tuner = tuner
}

func (this *Amplifier)SetDvd(dvd *DvdPlayer) {
	fmt.Println(this.description + " setting DVD player to " + dvd.String())
	this.dvd = dvd
}

func (this *Amplifier)SetCd(cd *CdPlayer) {
	fmt.Println(this.description + " setting CD player to " + cd.String())
	this.cd = cd
}

func (this *Amplifier)String() string {
	return this.description
}

//爆米花机
type PopcornPopper  struct {
	description string
}

func NewPopcornPopper(description string) *PopcornPopper {
	this := new(PopcornPopper)
	this.description = description
	return this
}

func (this *PopcornPopper)On() {
	fmt.Println(this.description + " on")
}

func (this *PopcornPopper)Off() {
	fmt.Println(this.description + " off")
}

func (this *PopcornPopper)Pop() {
	fmt.Println(this.description + " popping popcorn!")
}

func (this *PopcornPopper)String() string {
	return this.description
}

//投影仪
type Projector struct {
	description string
	dvd         *DvdPlayer
}

func NewProjector(description string, dvd *DvdPlayer) *Projector {
	this := new(Projector)
	this.description = description
	this.dvd = dvd
	return this
}

func (this *Projector)On() {
	fmt.Println(this.description + " on")
}

func (this *Projector)Off() {
	fmt.Println(this.description + " off")
}

func (this *Projector)WideScreenMode() {
	fmt.Println(this.description + " in widescreen mode (16x9 aspect ratio)")
}

func (this *Projector)TvMode() {
	fmt.Println(this.description + " in tv mode (4x3 aspect ratio)")
}

func (this *Projector)String() string {
	return this.description
}

//电影院灯
type TheaterLights struct {
	description string
}

func NewTheaterLights(description string) *TheaterLights {
	this := new(TheaterLights)
	this.description = description
	return this
}

func (this *TheaterLights)On() {
	fmt.Println(this.description + " on")
}

func (this *TheaterLights)Off() {
	fmt.Println(this.description + " off")
}

func (this *TheaterLights)Dim(level int) {
	fmt.Println(this.description + " dimming to " + string(level) + "\n")
}

func (this *TheaterLights)String() string {
	return this.description
}

//荧幕
type Screen struct {
	description string
}

func NewScreen(description string) *Screen {
	this := new(Screen)
	this.description = description
	return this
}

func (this *Screen)Up() {
	fmt.Println(this.description + " going up")
}

func (this *Screen)Down() {
	fmt.Println(this.description + " going down")
}

func (this *Screen)String() string {
	return this.description
}