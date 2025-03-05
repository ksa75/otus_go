package main

import (
	"fmt"
	"mocking/lesson/geocoding"
	"os"
	"time"

	"github.com/nathan-osman/go-sunrise"
)

// Это приложение по имени места, переданном в аргументе, показывает время восхода и заката солнца там.
// TODO: Что будем менять, чтобы это получше протестировать?
func main() {
	placeName := os.Args[0]

	gg := geocoding.GoogleGeocoder{
		ApiKey: os.Getenv("GOOGLE_MAPS_API_KEY"),
	}

	rise, set, err := CalcSunrise(placeName, gg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sunrise: %v; Sunset: %v\n", rise.Local().Format(time.TimeOnly), set.Local().Format(time.TimeOnly))
}

func CalcSunrise(placeName string, geocoder Geocoder) (rise, set time.Time, err error) {
	lat, long, err := geocoder.GetCoordsByName(placeName)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("get coords: %w", err)
	}

	rise, set = sunrise.SunriseSunset(
		lat, long,
		2000, time.January, 1, // TODO: set date
	)

	return
}

type Geocoder interface {
	GetCoordsByName(city string) (latitude, longitude float64, err error)
}
