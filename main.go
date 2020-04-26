package main

import (
	"fmt"
	"forecast/mail"
	"forecast/openweather"
)

func main() {
	openWeatherRequests := openweather.NewRequests()

	currentWeather, err := openWeatherRequests.GetCurrentWeather()
	if err != nil {
		fmt.Println(err)
		return
	}

	formatter := openweather.NewFormat()

	err = mail.NewMail().Send("Weather Today", formatter.FormatCurrentWeather(currentWeather))
	if err != nil {
		fmt.Println(err)
	}
}
