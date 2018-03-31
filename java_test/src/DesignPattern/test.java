package DesignPattern;


import DesignPattern.Observer.CurrentConditionsDisplay;
import DesignPattern.Observer.WeatherData;

/**
 * Created by Administrator on 2018/3/23.
 */
public class test {
    public static void main(String[] args) {
        WeatherData weatherData = new WeatherData();

        CurrentConditionsDisplay currentConditionsDisplay = new CurrentConditionsDisplay(weatherData);

        weatherData.setMeasurements(80, 65, 30.4f);
        weatherData.setMeasurements(70, 65, 30.4f);
        weatherData.setMeasurements(60, 65, 30.4f);
    }
}
