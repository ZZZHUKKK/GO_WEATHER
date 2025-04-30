package geo_test

import (
	"go/weather/geo"
	"testing"
)

func TestGetMyGeo(t *testing.T) {
	city := "London"
	expected := geo.GeoStruct{
		City: "London",
	}

	res, err := geo.GetMyGeo(city)

	if err != nil {
		t.Error("Не удается задать город")
	}
	if res.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected.City, res.City)
	}
}

func TestGetMyGeoNoCity(t *testing.T) {
	city := "ertwjdh"
	_, err := geo.GetMyGeo(city)
	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v, получено %v", geo.ErrNoCity, err)
	}
}
