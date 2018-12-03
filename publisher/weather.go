package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Endpoint for Dark Sky REST API
// example endpoint: /forcast/[key]/[latitude],[longitude]
const Endpoint = "https://api.darksky.net/"

// WeatherAPI wrapper struct for darksky api
type WeatherAPI struct {
	endpoint string
	key      string
	lat      float32
	long     float32
}

// CreateWeatherAPI initialize a simple WeatherAPI struct based  on key param
func CreateWeatherAPI(key string) *WeatherAPI {
	return &WeatherAPI{
		endpoint: Endpoint,
		key:      key,
		lat:      0,
		long:     0,
	}
}

// GetForecast Send HTTP Request to Forcast Endpoint Based on Coordinates
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

// SerializeJSON We first decode our http response body into a map[string]interface{}
// then we marshal/serialize this map into a []byte for transfer through redis
func SerializeJSON(res *http.Response) ([]byte, error) {
	var obj interface{}
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	serializedResponse, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return serializedResponse, nil
}
