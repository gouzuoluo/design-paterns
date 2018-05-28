package eg1

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	weatherData := NewWeatherData()

	NewStatisticsDisplay(weatherData)
	NewForecastDisplay(weatherData)
	NewHeatIndexDisplay(weatherData)
	currentDisplay := NewCurrentConditionsDisplay()
	currentDisplay.AddAttention(weatherData) //添加关注

	weatherData.SetMeasurements(80, 65, 30.4)

	fmt.Println("-------------------------")

	currentDisplay.CancelAttention(weatherData) //取消关注
	weatherData.SetMeasurements(70, 65, 30.4)
}
