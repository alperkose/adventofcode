package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SampleInputForPart1(t *testing.T) {
	observation := &Observation{}
	observation.Accept("...#......")
	observation.Accept(".......#..")
	observation.Accept("#.........")
	observation.Accept("..........")
	observation.Accept("......#...")
	observation.Accept(".#........")
	observation.Accept(".........#")
	observation.Accept("..........")
	observation.Accept(".......#..")
	observation.Accept("#...#.....")
	assert.Equal(t, 374, observation.SolvePt1())
}
