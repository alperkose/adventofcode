package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Adjacency(t *testing.T) {
	testCases := []struct {
		desc     string
		line     int
		pos      int
		expected bool
	}{
		{"", 2, 0, false},
		{"", 2, 1, true},
		{"", 2, 2, true},
		{"", 2, 3, true},
		{"", 2, 4, true},
		{"", 2, 5, true},
		{"", 2, 6, false},
		{"", 3, 0, false},
		{"", 3, 1, true},
		{"", 3, 5, true},
		{"", 3, 6, false},
		{"", 4, 0, false},
		{"", 4, 1, true},
		{"", 4, 2, true},
		{"", 4, 3, true},
		{"", 4, 4, true},
		{"", 4, 5, true},
		{"", 4, 6, false},
	}
	part_number := PartNumber{
		line:  3,
		begin: 2,
		end:   4,
		value: 123,
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expected, part_number.Adjacent_to(tC.line, tC.pos))
		})
	}
}
