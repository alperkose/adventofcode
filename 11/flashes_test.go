package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CountsFlashingOfAnOctopusInTheMiddle(t *testing.T) {
	lavaTubes := &DumboOctopuses{}
	lavaTubes.Accept("111")
	lavaTubes.Accept("191")
	lavaTubes.Accept("111")

	flashCount := CountFlashes(lavaTubes)
	assert.Equal(t, 1, flashCount)
}

func Test_FlashingOfAnOctopusInTheMiddleIncreasesTheOtherAround(t *testing.T) {
	lavaTubes := &DumboOctopuses{}
	lavaTubes.Accept("888")
	lavaTubes.Accept("898")
	lavaTubes.Accept("888")

	flashCount := CountFlashes(lavaTubes)
	assert.Equal(t, 9, flashCount)
}

func Test_ChainFlashingIncreasesTheLowerEnergies(t *testing.T) {
	lavaTubes := &DumboOctopuses{}
	lavaTubes.Accept("788")
	lavaTubes.Accept("898")
	lavaTubes.Accept("888")

	flashCount := CountFlashes(lavaTubes)
	assert.Equal(t, 9, flashCount)
}

func Test_StepsFromSampleInput(t *testing.T) {
	testCases := []struct {
		desc            string
		step            int
		expectedFlashes int
	}{
		{"step 1", 1, 0},
		{"step 2", 2, 35},
		{"step 3", 3, 80},
		{"step 4", 4, 96},
		{"step 5", 5, 104},
		{"step 6", 6, 105},
		{"step 7", 7, 112},
		{"step 8", 8, 136},
		{"step 9", 9, 175},
		{"step 10", 10, 204},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			dumboOctopuses := &DumboOctopuses{}
			dumboOctopuses.Accept("5483143223")
			dumboOctopuses.Accept("2745854711")
			dumboOctopuses.Accept("5264556173")
			dumboOctopuses.Accept("6141336146")
			dumboOctopuses.Accept("6357385478")
			dumboOctopuses.Accept("4167524645")
			dumboOctopuses.Accept("2176841721")
			dumboOctopuses.Accept("6882881134")
			dumboOctopuses.Accept("4846848554")
			dumboOctopuses.Accept("5283751526")

			count := FlashesAfterSteps(dumboOctopuses, tC.step)
			assert.Equal(t, tC.expectedFlashes, count)
		})
	}
}

func Test_CountUntilSynchronousFlash(t *testing.T) {
	dumboOctopuses := &DumboOctopuses{}
	dumboOctopuses.Accept("5483143223")
	dumboOctopuses.Accept("2745854711")
	dumboOctopuses.Accept("5264556173")
	dumboOctopuses.Accept("6141336146")
	dumboOctopuses.Accept("6357385478")
	dumboOctopuses.Accept("4167524645")
	dumboOctopuses.Accept("2176841721")
	dumboOctopuses.Accept("6882881134")
	dumboOctopuses.Accept("4846848554")
	dumboOctopuses.Accept("5283751526")

	step := SynchronousFlashStep(dumboOctopuses)
	assert.Equal(t, 195, step)
}
