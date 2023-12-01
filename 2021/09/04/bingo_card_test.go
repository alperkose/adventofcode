package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BingoCardMatching(t *testing.T) {
	testCases := []struct {
		desc          string
		draws         []string
		expectedScore int
	}{
		{
			desc:          "Matches a row in sixth draw",
			draws:         []string{"21", "9", "14", "16", "3", "7"},
			expectedScore: (300 - 70) * 7,
		},
		{
			desc:          "Matches a column in seventh draw",
			draws:         []string{"17", "23", "14", "16", "5", "3", "20"},
			expectedScore: (300 - 98) * 20,
		},
		{
			desc:          "Doesn't match a diagonal",
			draws:         []string{"22", "2", "14", "18", "19"},
			expectedScore: 0,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			card := &BingoCard{
				[]string{"22", "13", "17", "11", "0", "8", "2", "23", "4", "24", "21", "9", "14", "16", "7", "6", "10", "3", "18", "5", "1", "12", "20", "15", "19"},
				5,
				map[string]Cell{},
			}
			score := 0
			for i := 0; i < len(tC.draws); i++ {
				score = card.Bingo(tC.draws[i])
			}
			assert.Equal(t, tC.expectedScore, score)
		})
	}
}
