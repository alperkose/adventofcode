package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BasinIsFoundOnTheMiddle(t *testing.T) {

	lavaTubes := &LavaTubes{}
	lavaTubes.Accept("9999")
	lavaTubes.Accept("9149")
	lavaTubes.Accept("9999")

	factor := ThreeLargestBasinsMultiplied(lavaTubes)
	assert.Equal(t, 2, factor)
}
func Test_BasinIsFoundOnTheCorners(t *testing.T) {

	lavaTubes := &LavaTubes{}
	lavaTubes.Accept("1299")
	lavaTubes.Accept("9921")
	lavaTubes.Accept("9932")

	factor := ThreeLargestBasinsMultiplied(lavaTubes)
	assert.Equal(t, 8, factor)
}

func Test_SampleInputContaningMoreThaThreeBasins(t *testing.T) {

	lavaTubes := &LavaTubes{}
	lavaTubes.Accept("2199943210")
	lavaTubes.Accept("3987894921")
	lavaTubes.Accept("9856789892")
	lavaTubes.Accept("8767896789")
	lavaTubes.Accept("9899965678")

	factor := ThreeLargestBasinsMultiplied(lavaTubes)
	assert.Equal(t, 1134, factor)
}
