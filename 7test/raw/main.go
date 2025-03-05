package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kelvins/geocoder"
	"github.com/nathan-osman/go-sunrise"
)

// Это приложение по имени места, переданном в аргументе, показывает время восхода и заката солнца там.
// TODO: Что будем менять, чтобы это получше протестировать?
func main() {
	lat, long, err := getCoordsByName(os.Args[0])
	if err != nil {
		panic(err)
	}

	rise, set := sunrise.SunriseSunset(
		lat, long,
		2000, time.January, 1, // TODO: set date
	)

	fmt.Printf("Sunrise: %v; Sunset: %v\n", rise.Local().Format(time.TimeOnly), set.Local().Format(time.TimeOnly))
}

func getCoordsByName(city string) (latitude, longitude float64, err error) {
	geocoder.ApiKey = os.Getenv("GOOGLE_MAPS_API_KEY")

	address := geocoder.Address{
		City: city,
	}

	// Convert address to location (latitude, longitude)
	location, err := geocoder.Geocoding(address)
	return location.Latitude, location.Longitude, err
}
