package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Maybe change to just api.darksky.net
// REST Forcast Endpoint: /forcast/[key]/[latitude],[longitude]
const Endpoint = "https://api.darksky.net/"

type WeatherAPI struct {
	endpoint string
	key      string
	lat      float32
	long     float32
}

// initialize a simple WeatherAPI struct based  on key param
func CreateWeatherAPI(key string) *WeatherAPI {
	return &WeatherAPI{
		endpoint: Endpoint,
		key:      key,
		lat:      0,
		long:     0,
	}
}

// Send HTTP Request to Forcast Endpoint Based on Coordinates
func (w *WeatherAPI) GetForecast() (*http.Response, error) {
	var sb strings.Builder
	sb.WriteString(w.endpoint)
	sb.WriteString("forecast/")
	sb.WriteString(w.key)
	sb.WriteString("/")
	sb.WriteString(fmt.Sprint(w.lat))
	sb.WriteString(",")
	sb.WriteString(fmt.Sprint(w.long))
	
	res, err := http.Get(sb.String())
	if err != nil {
		return nil, err
	}
	return res, err
}
