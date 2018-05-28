package eg1

//主题类，管理着某些数据
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

/*===================================================================================================================*/

//天气数据类
type WeatherData struct {
	observers                       []Observer //精髓之处
	temperature, humidify, pressure float32
}

func NewWeatherData() *WeatherData {
	this := new(WeatherData)
	this.observers = make([]Observer, 0, 10)
	return this
}

func (this *WeatherData) RegisterObserver(observer Observer) {
	this.observers = append(this.observers, observer)
}

func (this *WeatherData) RemoveObserver(observer Observer) {
	for i, o := range this.observers {
		if o == observer {
			this.observers = append(this.observers[:i], this.observers[i+1:]...)
			break
		}
	}
}

func (this *WeatherData) NotifyObservers() {
	for _, o := range this.observers {
		o.Update(this.temperature, this.humidify, this.pressure)
	}
}

func (this *WeatherData) MeasurementChanged() {
	this.NotifyObservers()
}

func (this *WeatherData) SetMeasurements(temperature, humidify, pressure float32) {
	this.temperature = temperature
	this.humidify = humidify
	this.pressure = pressure

	this.MeasurementChanged()
}

func (this *WeatherData) GetTemperature() float32 {
	return this.temperature
}

func (this *WeatherData) GetHumidity() float32 {
	return this.humidify
}

func (this *WeatherData) GetPressure() float32 {
	return this.pressure
}
