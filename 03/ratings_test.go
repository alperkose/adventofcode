package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_O2GenerationRating(t *testing.T) {
	values := []uint32{4, 30, 22, 23, 21, 15, 7, 28, 16, 25, 2, 10}
	digits := 5

	actual := O2GenerationRating(values, uint32(digits))
	assert.Equal(t, uint32(23), actual)
}

func Test_CO2ScrubberRating(t *testing.T) {
	values := []uint32{4, 30, 22, 23, 21, 15, 7, 28, 16, 25, 2, 10}
	digits := 5

	actual := CO2ScrubberRating(values, uint32(digits))
	assert.Equal(t, uint32(10), actual)
}
