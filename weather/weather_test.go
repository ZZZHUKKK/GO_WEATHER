package weather_test

import (
	"go/weather/geo"
	"go/weather/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "London"
	city := &geo.GeoStruct{
		City: expected,
	}
	format := 4

	res, err := weather.GetWeather(city, format)
	if err != nil {
		t.Error("Не удается задать город")
	}
	if !strings.Contains(res, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, res)
	}

}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big Format", format: 125},
	{name: "0 Format", format: 0},
	{name: "Minus Format", format: -3},
}

func TestGetWeatherIncorrectFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "London"
			city := &geo.GeoStruct{
				City: expected,
			}
			_, err := weather.GetWeather(city, tc.format)
			if err != weather.ErrFormat {
				t.Errorf("Ожидалось %v, получено %v", weather.ErrFormat, err)
			}
		})
	}
}
