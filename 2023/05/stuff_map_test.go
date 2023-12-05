package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StuffMapping(t *testing.T) {
	testCases := []struct {
		desc     string
		input    StuffMapEntry
		check    int
		expected int
	}{
		{"", StuffMapEntry{50, 98, 2}, 98, 50},
		{"", StuffMapEntry{50, 98, 2}, 99, 51},
		{"", StuffMapEntry{50, 98, 2}, 97, 0},
		{"", StuffMapEntry{50, 98, 2}, 100, 0},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expected, tC.input.Get(tC.check))
		})
	}
}

func Test_SuffMappingEntry(t *testing.T) {
	stuffMap := StuffMap{StuffMapEntry{50, 98, 2}, StuffMapEntry{52, 50, 48}}
	assert.Equal(t, 81, stuffMap.Get(79))
}
