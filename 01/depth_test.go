package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DepthIncreases(t *testing.T) {
	testCases := []struct {
		desc          string
		depths        []int
		expectedCount int
	}{
		{"Single depth", []int{100}, 0},
		{"Two depths, increasing", []int{100, 101}, 1},
		{"Two depths, decreasing", []int{100, 99}, 0},
		{"Increasing & decreasing", []int{100, 101, 100}, 1},
		{"Sample provided", []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 7},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cntDepthInc := CountDepthIncrease(tC.depths)
			assert.Equal(t, tC.expectedCount, cntDepthInc)
		})
	}
}
func Test_SlidingWindowOfDepthIncreases(t *testing.T) {
	testCases := []struct {
		desc          string
		depths        []int
		expectedCount int
	}{
		{"Single depth", []int{100}, 0},
		{"Two depths, increasing", []int{100, 101}, 0},
		{"Two depths, decreasing", []int{100, 99, 101}, 0},
		{"Increasing & decreasing", []int{100, 99, 101, 102}, 1},
		{"Increasing & decreasing", []int{100, 99, 101, 99}, 0},
		{"Sample provided", []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 5},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cntDepthInc := SlidingWindowOfDepthIncreases(tC.depths)
			assert.Equal(t, tC.expectedCount, cntDepthInc)
		})
	}
}
