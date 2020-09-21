package main

import (
	"fmt"
	"forecast/mail"
	"forecast/openweather"
	"forecast/vacations"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("DEV") == "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	recipients := mail.NewMail().GetRecipients()

	for _, r := range recipients {
		openWeatherRequests := openweather.NewRequests()
		vacationSpot := vacations.GetLocation()
		currentLocation := r.Location

		if vacationSpot != nil {
			currentLocation = vacationSpot.Location
		}

		currentWeather, err := openWeatherRequests.GetCurrentWeather(currentLocation)
		if err != nil {
			fmt.Println(err)
			return
		}

		oneCallWeather, err := openWeatherRequests.GetOneCallWeather(currentLocation)
		if err != nil {
			fmt.Println(err)
			return
		}

		formatter := openweather.NewFormat()

		emailBody := formatter.FormatCurrentWeather(currentWeather) + "\n" + formatter.FormatOneCallWeather(oneCallWeather)

		err = mail.NewMail().Send([]string{r.Email}, "Weather Today", emailBody)
		if err != nil {
			log.Fatal(err)
		}
	}
}
