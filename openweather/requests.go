package openweather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

/*
Requests is the struct that maintains state and methods
for interacting with the OpenWeather API
*/
type Requests struct {
	APIKEY string
	CityID string
}

/*
CurrentWeather is a struct that has the same shape
as the JSON returned by the OpenWeather current weather API
*/
type CurrentWeather struct {
	Name    string    `json:"name"`
	Weather []weather `json:"weather"`
	Temp    temp      `json:"main"`
}

type weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

type temp struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
}

/*
GetCurrentWeather uses the OpenWeather current weather API
to get the current weather for the given CityID
*/
func (r *Requests) GetCurrentWeather() (*CurrentWeather, error) {
	url := "https://api.openweathermap.org/data/2.5/weather?id=" + r.CityID + "&appid=" + r.APIKEY + "&units=imperial"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cw := CurrentWeather{}
	err = json.Unmarshal(body, &cw)
	if err != nil {
		return nil, err
	}

	return &cw, nil
}

/*
NewRequests sets up a new OpenWeather request struct
to make requests with that API
*/
func NewRequests(cityID string) *Requests {
	return &Requests{
		APIKEY: os.Getenv("OPENWEATHER_API_KEY"),
		CityID: cityID,
	}
}
