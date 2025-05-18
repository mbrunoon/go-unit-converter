package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeterToKilometer(t *testing.T) {
	conversor, err := NewConversor("meters", 1000., "kilometers")
	assert.Empty(t, err)

	result, err := conversor.Result()

	assert.Empty(t, err)
	assert.Equal(t, 1., result)
}

func TestMetersToKilograms(t *testing.T) {
	_, err := NewConversor("grams", 1000, "meters")
	assert.NotEmpty(t, err)
}

func TestInvalidConversion(t *testing.T) {
	conversor, err := NewConversor("invalid", 1000., "invalid")
	assert.NotEmpty(t, err)

	_, err = conversor.Result()
	assert.NotEmpty(t, err)
}
