package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreatesNewBinaryFromCommonBits(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []uint32
		expected uint32
	}{
		{"returns same if single item", []uint32{101}, uint32(101)},
		{"calculates it for 3 items", []uint32{1, 4, 6}, uint32(4)},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := GammaRate(tC.input)
			assert.Equal(t, tC.expected, result)
		})
	}
}
