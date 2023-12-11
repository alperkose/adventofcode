package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Prediction(t *testing.T) {
	testCases := []struct {
		desc     string
		history  HistoricalData
		expected int
	}{
		{"0 3 6 9 12 15", HistoricalData{0, 3, 6, 9, 12, 15}, 18},
		{"1 3 6 10 15 21", HistoricalData{1, 3, 6, 10, 15, 21}, 28},
		{"10 13 16 21 30 45", HistoricalData{10, 13, 16, 21, 30, 45}, 68},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.history.PredictNext(), tC.expected)
		})
	}
}
