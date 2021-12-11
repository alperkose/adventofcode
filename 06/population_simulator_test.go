package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cycles(t *testing.T) {
	testCases := []struct {
		desc               string
		cycles             string
		days               int
		expectedPopulation int
	}{
		{"No spawn before the cycle", "5", 4, 1},
		{"spawns after 7 days", "6", 7, 2},
		{"spawns twice after 14 days", "6", 14, 3},
		{"child also spawns after 16 days", "6", 16, 4},
		{"from sample after 18 days", "3,4,3,1,2", 18, 26},
		{"from sample after 80 days", "3,4,3,1,2", 80, 5934},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lf := &LanternFishes{}
			lf.Accept(tC.cycles)
			simulatedPopulation := Simulate(lf, tC.days)
			assert.Equal(t, tC.expectedPopulation, simulatedPopulation)
		})
	}
}
