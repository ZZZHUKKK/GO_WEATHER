package main

import (
	"flag"
	"fmt"
	"go/weather/geo"
	"go/weather/weather"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода")

	flag.Parse()

	geo, err := geo.GetMyGeo(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	weather, err := weather.GetWeather(geo, *format)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(weather)
}
