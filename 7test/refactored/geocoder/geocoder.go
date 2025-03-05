package geo

import (
	"github.com/kelvins/geocoder"
)

type GeocoderRealImpl struct {
	ApiKey string
}

func (g GeocoderRealImpl) GetCoordsByName(city string) (latitude, longitude float64, err error) {
	geocoder.ApiKey = g.ApiKey

	address := geocoder.Address{
		City: city,
	}

	// Convert address to location (latitude, longitude)
	location, err := geocoder.Geocoding(address)
	return location.Latitude, location.Longitude, err
}
