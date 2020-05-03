# :sunny: forecast :sunny:

A small Golang app to send the local weather info!

## Context

I wanted to play around a bit more with Golang, in particular handling pointers and HTTP requests. I also wanted to finally check out Github Actions and utilize the cron job feature to run actions on an automated schedule. I thought it might be fun and mildly useful to use Golang for querying weather data and then Github Actions to send that data every morning to my email. From there, `forecast` was born.

## Integrations

Utilizes the [OpenWeather current weather API](https://openweathermap.org/current) for gathering local weather data for my location.