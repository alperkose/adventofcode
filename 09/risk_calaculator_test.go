package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculatesTheRiskOfTheHoleInTheMiddle(t *testing.T) {
	lavaTubes := &LavaTubes{}
	lavaTubes.Accept("423")
	lavaTubes.Accept("614")
	lavaTubes.Accept("789")

	risk := RiskFactorOfHoles(lavaTubes)
	assert.Equal(t, 2, risk)
}
func Test_CalculatesTheRiskOfTheHoleInTheCorner(t *testing.T) {
	lavaTubes := &LavaTubes{}
	lavaTubes.Accept("676")
	lavaTubes.Accept("697")
	lavaTubes.Accept("789")

	risk := RiskFactorOfHoles(lavaTubes)
	assert.Equal(t, 7, risk)
}

func Test_CalculatesTheRiskOfMultipleHoles(t *testing.T) {
	lavaTubes := &LavaTubes{}
	lavaTubes.Accept("6774")
	lavaTubes.Accept("5975")
	lavaTubes.Accept("7838")
	lavaTubes.Accept("7858")

	risk := RiskFactorOfHoles(lavaTubes)
	assert.Equal(t, 15, risk)
}
func Test_CalculatesTheRiskOfSampleInput(t *testing.T) {
	lavaTubes := &LavaTubes{}
	lavaTubes.Accept("2199943210")
	lavaTubes.Accept("3987894921")
	lavaTubes.Accept("9856789892")
	lavaTubes.Accept("8767896789")
	lavaTubes.Accept("9899965678")

	risk := RiskFactorOfHoles(lavaTubes)
	assert.Equal(t, 15, risk)
}
