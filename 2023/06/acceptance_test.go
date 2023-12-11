package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SampeleInputForPart1(t *testing.T) {
	records := &RaceRecords{}
	records.Accept("Time:      7  15   30")
	records.Accept("Distance:  9  40  200")

	assert.Equal(t, 288, records.SolvePt1())
}
