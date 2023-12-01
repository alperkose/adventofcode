package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HypoThermalVentMap(t *testing.T) {
	testCases := []struct {
		desc          string
		input         string
		expectedCloud Cloud
	}{
		{
			desc:          "HorizontalCloud",
			input:         "0,9 -> 5,9",
			expectedCloud: Cloud{Point{0, 9}, Point{5, 9}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			htvMap := &HypoThermalVentMap{}
			htvMap.Accept(tC.input)
			assert.Equal(t, tC.expectedCloud, htvMap.Clouds()[0])
		})
	}
}
