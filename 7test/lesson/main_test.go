package main_test

import (
	main "mocking/lesson"
	"mocking/lesson/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	gm := &mocks.Geocoder{}
	//gm.On("GetCoordsByName", "Ekaterinburg", "dsgdsg").Return(45, 45.0, nil)
	gm.EXPECT().GetCoordsByName("Ekaterinburg2").Return(45, 45, nil)

	rise, set, err := main.CalcSunrise("Ekaterinburg2", gm)
	require.NoError(t, err)
	assert.Equal(t, time.Date(2000, time.January, 1, 4, 38, 13, 0, time.UTC), rise)
	assert.Equal(t, time.Date(2000, time.January, 1, 13, 28, 2, 0, time.UTC), set)
}
