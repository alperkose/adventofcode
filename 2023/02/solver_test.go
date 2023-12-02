package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc:     "Single Line possible",
			input:    []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			expected: 1,
		}, {
			desc:     "Single Line not possible",
			input:    []string{"Game 1: 20 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			expected: 0,
		}, {
			desc:     "Multiple lines both possible",
			input:    []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"},
			expected: 3,
		}, {
			desc: "Multiple lines some are not possible",
			input: []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			expected: 8,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			p := &GameParser{}
			for _, inp := range tC.input {
				p.Accept(inp)
			}
			assert.Equal(t, tC.expected, SolverForPart1().Solve(p.Games))
		})
	}
}

func Test_Part2(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc:     "Single Line ",
			input:    []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			expected: 48,
		}, {
			desc:     "Single Line ",
			input:    []string{"Game 1: 20 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			expected: 160,
		}, {
			desc:     "Multiple lines with addition",
			input:    []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"},
			expected: 60,
		}, {
			desc: "Sample input",
			input: []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			expected: 2286,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			p := &GameParser{}
			for _, inp := range tC.input {
				p.Accept(inp)
			}
			assert.Equal(t, tC.expected, SolverForPart2().Solve(p.Games))
		})
	}
}
