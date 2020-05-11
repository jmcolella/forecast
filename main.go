package main

import (
	"fmt"
	"forecast/mail"
	"forecast/openweather"
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
		openWeatherRequests := openweather.NewRequests(r.CityID)

		currentWeather, err := openWeatherRequests.GetCurrentWeather()
		if err != nil {
			fmt.Println(err)
			return
		}

		formatter := openweather.NewFormat()

		err = mail.NewMail().Send([]string{r.Email}, "Weather Today", formatter.FormatCurrentWeather(currentWeather))
		if err != nil {
			log.Fatal(err)
		}
	}
}
