package main

import (
	"fmt"
	"os"
	"time"

	geo "mocking/refactored/geocoder"
	"mocking/refactored/sun"
)

func main() {
	mapsApiKey := os.Getenv("GOOGLE_MAPS_API_KEY")

	geocoder := geo.GeocoderRealImpl{ApiKey: mapsApiKey}

	sun := sun.SunInformer{Geocoder: geocoder}

	rise, set, err := sun.Info(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sunrise: %v; Sunset: %v\n", rise.Local().Format(time.TimeOnly), set.Local().Format(time.TimeOnly))
}
