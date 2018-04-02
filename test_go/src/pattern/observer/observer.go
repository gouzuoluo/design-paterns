package observer

import (
	"fmt"
)

type Observer interface {
	Update(temp, humidify, pressure float32)
}

type DisplayElement interface {
	Display()
}

/*----------------------------------------------------------------------------------------------------------------*/

//布告板1
type CurrentConditionsDisplay struct {
	temperature float32
	humidity    float32
	//可关注多个不同的主题
	attentions []Subject //这里针对接口的编程方式更好，在运行时可以更换关注的主题
}

func NewCurrentConditionsDisplay() *CurrentConditionsDisplay {
	this := new(CurrentConditionsDisplay)
	this.attentions = make([]Subject, 0, 8)
	return this
}

//添加关注
func (this *CurrentConditionsDisplay) AddAttention(s Subject) bool {
	for _, v := range this.attentions {
		if v == s {
			return false
		}
	}
	this.attentions = append(this.attentions, s)
	s.RegisterObserver(this)
	return true
}

//取消关注
func (this *CurrentConditionsDisplay) CancelAttention(s Subject) bool {
	for i, v := range this.attentions {
		if v == s {
			this.attentions = append(this.attentions[:i], this.attentions[:i+1]...)
			s.RemoveObserver(this)
			return true
		}
	}
	return false
}

func (this *CurrentConditionsDisplay) Update(temp, humidify, pressure float32) {
	this.temperature = temp
	this.humidity = humidify

	this.Display()
}

func (this *CurrentConditionsDisplay) Display() {
	fmt.Println(fmt.Sprintf("Current conditions: %f%s%f%s",
		this.temperature, "F degrees and ", this.humidity, "% humidity"))
}

//布告板2
type StatisticsDisplay struct {
	maxTemp, minTemp, tempSum float32
	numReadings               int
	weatherData               *WeatherData //这里针对实现的编程不是很好
}

func NewStatisticsDisplay(weatherData *WeatherData) *StatisticsDisplay {
	this := new(StatisticsDisplay)
	this.weatherData = weatherData
	weatherData.RegisterObserver(this)
	return this
}

func (this *StatisticsDisplay) Update(temp, humidify, pressure float32) {
	this.tempSum += temp
	this.numReadings++

	if temp > this.maxTemp {
		this.maxTemp = temp
	}

	if temp < this.minTemp {
		this.minTemp = temp
	}

	this.Display()
}

func (this *StatisticsDisplay) Display() {
	fmt.Println(fmt.Sprintf("Avg/Max/Min temperature = %f%s%f%s%f",
		this.tempSum/float32(this.numReadings), "/", this.maxTemp, "/", this.minTemp))
}

//布告板3
type ForecastDisplay struct {
	currentPressure, lastPressure float32
	weatherData                   *WeatherData //这里针对实现的编程不是很好
}

func NewForecastDisplay(weatherData *WeatherData) *ForecastDisplay {
	this := new(ForecastDisplay)
	this.weatherData = weatherData
	this.currentPressure = 29.92
	weatherData.RegisterObserver(this)
	return this
}

func (this *ForecastDisplay) Update(temp, humidify, pressure float32) {
	this.lastPressure = this.currentPressure
	this.currentPressure = pressure

	this.Display()
}

func (this *ForecastDisplay) Display() {
	fmt.Sprintln("Forecast: ")
	if this.currentPressure > this.lastPressure {
		fmt.Println("Improving weather on the way!")
	} else if this.currentPressure == this.lastPressure {
		fmt.Println("More of the same")
	} else if this.currentPressure < this.lastPressure {
		fmt.Println("Watch out for cooler, rainy weather")
	}
}

//布告板4
type HeatIndexDisplay struct {
	heatIndex   float32
	weatherData *WeatherData //这里针对实现的编程不是很好
}

func NewHeatIndexDisplay(weatherData *WeatherData) *HeatIndexDisplay {
	this := new(HeatIndexDisplay)
	this.weatherData = weatherData
	weatherData.RegisterObserver(this)
	return this
}

func (this *HeatIndexDisplay) Update(temp, humidify, pressure float32) {
	this.heatIndex = this.computeHeatIndex(temp, humidify)
	this.Display()
}

func (this *HeatIndexDisplay) computeHeatIndex(t, rh float32) float32 {
	index := (16.923 + (0.185212 * t) + (5.37941 * rh) - (0.100254 * t * rh) +
		(0.00941695 * (t * t)) + (0.00728898 * (rh * rh)) +
		(0.000345372 * (t * t * rh)) - (0.000814971 * (t * rh * rh)) +
		(0.0000102102 * (t * t * rh * rh)) - (0.000038646 * (t * t * t)) + (0.0000291583 *
		(rh * rh * rh)) + (0.00000142721 * (t * t * t * rh)) +
		(0.000000197483 * (t * rh * rh * rh)) - (0.0000000218429 * (t * t * t * rh * rh)) +
		0.000000000843296*(t*t*rh*rh*rh)) -
		(0.0000000000481975 * (t * t * t * rh * rh * rh))
	return index
}

func (this *HeatIndexDisplay) Display() {
	fmt.Println("Heat index is ", this.heatIndex)
}
