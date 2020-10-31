package main

import (
	"context"
	"fmt"
	"forecast/ggl"
	"forecast/mail"
	"forecast/openweather"
	"forecast/users"
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

	ctx := context.Background()
	firestoreClient, err := ggl.NewFireStoreClient(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer firestoreClient.Store.Close()

	recipients, err := users.Fetch(ctx, firestoreClient)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, r := range recipients {
		openWeatherRequests := openweather.NewRequests()

		currentWeather, err := openWeatherRequests.GetCurrentWeather(r.Location)
		if err != nil {
			fmt.Println(err)
			return
		}

		oneCallWeather, err := openWeatherRequests.GetOneCallWeather(r.Location)
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
