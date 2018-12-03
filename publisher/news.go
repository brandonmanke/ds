package main

import (
	"net/http"
	"strings"
)

// TopStoriesEndpoint RESTful endpoint for top stories of NYTimes API
const TopStoriesEndpoint = "https://api.nytimes.com/svc/topstories/v2/"

// NYTimesAPI wrapper struct around NYTimes API Endpoints
type NYTimesAPI struct {
	endpoint string
	key      string
}

// CreateNYTimesAPI creates a new struct wrapper for making API calls
func CreateNYTimesAPI(key string) *NYTimesAPI {
	return &NYTimesAPI{
		endpoint: TopStoriesEndpoint,
		key:      key,
	}
}

// GetHome api request to get latest home page stories from NYTimes
func (api *NYTimesAPI) GetHome() (*http.Response, error) {
	var sb strings.Builder
	sb.WriteString(api.endpoint)
	sb.WriteString("home.json")
	sb.WriteString("?api-key=")
	sb.WriteString(api.key)

	res, err := http.Get(sb.String())
	if err != nil {
		return nil, err
	}
	return res, err
}
