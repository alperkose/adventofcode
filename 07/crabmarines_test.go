package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CrabMarinesPlacement(t *testing.T) {
	testCases := []struct {
		input               string
		expectedCrabMarines []int
	}{
		{"16,1,2,0,4,2,7,1,2,14", []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			crabmarines := &CrabMarines{}
			crabmarines.Accept(tC.input)
			assert.Equal(t, tC.expectedCrabMarines, crabmarines.Placements())
		})
	}
}
