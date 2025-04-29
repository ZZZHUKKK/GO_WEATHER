package weather

import (
	"fmt"
	"go/weather/geo"
	"io"
	"net/http"
	"net/url"
)

func GetWeather(geo *geo.GeoStruct, format int) (string, error) {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %v", err)
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	// fmt.Println(baseUrl)
	response, err := http.Get(baseUrl.String())
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read body: %v", err)
	}

	return string(body), nil

}
