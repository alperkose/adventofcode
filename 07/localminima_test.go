package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindLocalMinima(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []int
		expected int
	}{
		{"single point returns itself", []int{1}, 1},
		{"three point converges on the middle", []int{3, 1, 2}, 2},
		{"converges to 2", []int{2, 1, 1, 3, 3}, 2},
		{"from sample input", []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, 2},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := LocalMinima(tC.input, fuelConsumption)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
func Test_FindLocalMinimaWithIncreasingFuelConsumption(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []int
		expected int
	}{
		{"from sample input", []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, 5},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := LocalMinima(tC.input, IncreasingFuelConsumption)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
