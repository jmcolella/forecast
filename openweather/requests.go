package openweather

import (
	"encoding/json"
	"forecast/mail"
	"io/ioutil"
	"net/http"
	"os"
)

/*
Requests is the struct that maintains state and methods
for interacting with the OpenWeather API
*/
type Requests struct {
	BaseURL string
	APIKEY  string
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
OneCallWeather is a struct that has the same shape
as the JSON returned by the OpenWeather OneCall weather API
*/
type OneCallWeather struct {
	Hourly []hourlyWeather `json:"hourly"`
	Daily  []dailyWeather  `json:"daily"`
}

type hourlyWeather struct {
	DT        int64     `json:"dt"`
	Temp      float32   `json:"temp"`
	FeelsLike float32   `json:"feels_like"`
	Weather   []weather `json:"weather"`
}

type dailyWeather struct {
	DT        int64     `json:"dt"`
	Sunrise   int64     `json:"sunrise"`
	Sunset    int64     `json:"sunset"`
	Temp      dailyTemp `json:"temp"`
	FeelsLike dailyTemp `json:"feels_like"`
	Weather   []weather `json:"weather"`
}

type dailyTemp struct {
	Day   float32 `json:"day"`
	Night float32 `json:"night"`
	Morn  float32 `json:"morn"`
	Eve   float32 `json:"eve"`
}

/*
GetCurrentWeather uses the OpenWeather current weather API
to get the current weather for the given Long and Lat
*/
func (r *Requests) GetCurrentWeather(location mail.RecipientLocation) (*CurrentWeather, error) {
	url := r.BaseURL + "/weather?appid=" + r.APIKEY + "&lat=" + location.Lat + "&lon=" + location.Long + "&units=imperial"

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
GetOneCallWeather uses the OpenWeather OneCall weather API
to get the hourly and daily weather for the given Long and Lat
*/
func (r *Requests) GetOneCallWeather(location mail.RecipientLocation) (*OneCallWeather, error) {
	url := r.BaseURL + "/onecall?appid=" + r.APIKEY + "&lat=" + location.Lat + "&lon=" + location.Long + "&units=imperial"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	oneCallWeather := OneCallWeather{}
	err = json.Unmarshal(body, &oneCallWeather)
	if err != nil {
		return nil, err
	}

	return &oneCallWeather, nil
}

/*
NewRequests sets up a new OpenWeather request struct
to make requests with that API
*/
func NewRequests() *Requests {
	return &Requests{
		BaseURL: os.Getenv("OPENWEATHER_BASE_URL"),
		APIKEY:  os.Getenv("OPENWEATHER_API_KEY"),
	}
}
