package sun

import (
	"fmt"
	"time"

	"github.com/nathan-osman/go-sunrise"
)

//go:generate mockery --name Geocoder --with-expecter
type Geocoder interface {
	GetCoordsByName(city string) (latitude, longitude float64, err error)
}

type SunInformer struct {
	Geocoder Geocoder
}

func (s SunInformer) Info(city string) (rise, set time.Time, err error) {
	geocoder := s.Geocoder

	lat, long, err := geocoder.GetCoordsByName(city)
	if err != nil {
		return rise, set, fmt.Errorf("get coords by name: %w", err)
	}

	rise, set = sunrise.SunriseSunset(
		lat, long,
		2000, time.January, 1, // TODO: set date
	)

	return
}
