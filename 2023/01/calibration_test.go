package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Calibration(t *testing.T) {
	testCases := []struct {
		desc                string
		input               []string
		expectedCalibration int
	}{
		{"two digit number", []string{"14"}, 14},
		{"three digit number", []string{"234"}, 24},
		{"alphanumeric input with 2 digits placed", []string{"a1b2"}, 12},
		{"alphanumeric input with single digit placed", []string{"a1b"}, 11},
		{"summing multiple lines", []string{"c4t", "3ch0"}, 74},
		{"numbers in written form", []string{"3ch0one"}, 31},
		{"sample input", []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}, 281},
		{"broken", []string{"75447"}, 77},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cal := &CalibrationInput{}
			for i := 0; i < len(tC.input); i++ {
				cal.Accept(tC.input[i])
			}
			assert.Equal(t, tC.expectedCalibration, cal.Resolve())
		})
	}
}
