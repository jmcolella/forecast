package openweather

import (
	"fmt"
	"time"
)

/*
Format is an empty struct with methods
that aid in formatting OpenWeather requests
into human-readable text
*/
type Format struct{}

/*
FormatCurrentWeather uses a currentWeather struct
to format the data into human-readable text
*/
func (f *Format) FormatCurrentWeather(currentWeather *CurrentWeather) string {
	name := currentWeather.Name
	mainDesc := currentWeather.Weather[0].Main
	temp := currentWeather.Temp.Temp
	feelsLike := currentWeather.Temp.FeelsLike
	tempMax := currentWeather.Temp.TempMax

	open := fmt.Sprintf("In %s today:", name)
	desc := fmt.Sprintf("It is %v", mainDesc)
	currentDesc := fmt.Sprintf("With a current temp of %vF degrees, and a high of %vF degrees.", temp, tempMax)
	feelsDesc := fmt.Sprintf("It feels like %vF degrees.", feelsLike)

	return fmt.Sprintf("\n%v\n%v\n%v\n%v\n", open, desc, currentDesc, feelsDesc)
}

/*
FormatOneCallWeather uses a oneCallWeather struct
to format the data into human-readable text
*/
func (f *Format) FormatOneCallWeather(oneCallWeather *OneCallWeather) string {
	var hourly string
	var daily string

	now := time.Now().Hour()
	hoursToEOD := 24 - now
	hourlyCount := 0

	for hourlyCount < hoursToEOD {
		h := oneCallWeather.Hourly[hourlyCount]

		hour := formatHour(h.DT)
		currentHour := fmt.Sprintf("%v: %vF degrees, %v", hour, h.Temp, h.Weather[0].Description)

		hourly = fmt.Sprintf("%s\n", hourly+currentHour)

		hourlyCount++
	}

	dailyCount := 1
	dailyCutOff := 5

	for dailyCount <= dailyCutOff {
		d := oneCallWeather.Daily[dailyCount]

		day := formatDay(d.DT)
		currentDay := fmt.Sprintf("%s: %vF degrees, %v", day, d.Temp.Day, d.Weather[0].Description)

		daily = fmt.Sprintf("%s\n", daily+currentDay)

		dailyCount++
	}

	return fmt.Sprintf("The rest of the day:\n%s\nThe next 5 days:\n%s", hourly, daily)
}

func formatHour(unixTime int64) string {
	hour := time.Unix(unixTime, 0).Hour()

	if hour <= 12 {
		return fmt.Sprintf("%vam", hour)
	}

	return fmt.Sprintf("%vpm", hour-12)
}

func formatDay(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	day := t.Day()
	month := t.Month()
	weekday := t.Weekday()

	return fmt.Sprintf("%s %s %v", weekday, month, day)
}

/*
NewFormat returns the Format struct
which has methods to aid in formatting OpenWeather requests
*/
func NewFormat() *Format {
	return &Format{}
}
