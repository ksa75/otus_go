package sun_test

import (
	"mocking/refactored/sun"
	"mocking/refactored/sun/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSun(t *testing.T) {
	mockedGeo := mocks.NewGeocoder(t)
	//mockedGeo.On("GetCoordsByName", "London").Return(51.5074, 0.1278, nil)
	mockedGeo.EXPECT().GetCoordsByName("London").Return(51.5074, 0.1278, nil)

	s := sun.SunInformer{mockedGeo}

	rise, set, err := s.Info("London")
	require.NoError(t, err)

	assert.Equal(t, "08:05:10", rise.Format("15:04:05"))
	assert.Equal(t, "16:00:11", set.Format("15:04:05"))
}
