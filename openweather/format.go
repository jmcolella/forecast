package openweather

import "fmt"

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

	return fmt.Sprintf(`
		%v
		%v
		%v
		%v
	`, open, desc, currentDesc, feelsDesc)
}

/*
NewFormat returns the Format struct
which has methods to aid in formatting OpenWeather requests
*/
func NewFormat() *Format {
	return &Format{}
}
