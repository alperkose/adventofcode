package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IncreasingFuelConsumption(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []int
		to       int
		expected int
	}{
		{"single crab single step", []int{1}, 2, 1},
		{"single crab two steps", []int{1}, 3, 3},
		{"single crab three steps", []int{1}, 4, 6},
		{"from sample input", []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, 5, 168},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := IncreasingFuelConsumption(tC.input, tC.to)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
