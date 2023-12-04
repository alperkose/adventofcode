package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SampleInput(t *testing.T) {
	schematic := &Schematic{}
	schematic.Accept("467..114..")
	schematic.Accept("...*......")
	schematic.Accept("..35..633.")
	schematic.Accept("......#...")
	schematic.Accept("617*......")
	schematic.Accept(".....+.58.")
	schematic.Accept("..592.....")
	schematic.Accept("......755.")
	schematic.Accept("...$.*....")
	schematic.Accept(".664.598..")

	assert.Equal(t, 4361, schematic.Solve())
	assert.Equal(t, 467835, schematic.SolvePt2())
}

func Test_Acceptance_Part1(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc:     "left of symbol",
			input:    []string{"123*"},
			expected: 123,
		}, {
			desc:     "right of another symbol",
			input:    []string{"+432"},
			expected: 432,
		}, {
			desc:     "left of another symbol with redundancy",
			input:    []string{"12+.32"},
			expected: 12,
		}, {
			desc:     "right of another symbol with redundancy",
			input:    []string{"12.+54"},
			expected: 54,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			schematic := &Schematic{}
			for _, v := range tC.input {
				schematic.Accept(v)
			}
			assert.Equal(t, tC.expected, schematic.Solve())
		})
	}
}

func Test_Acceptance_Part2(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc:     "left of symbol",
			input:    []string{"123*"},
			expected: 0,
		}, {
			desc:     "right of another symbol",
			input:    []string{"+432"},
			expected: 0,
		}, {
			desc:     "left of another symbol with redundancy",
			input:    []string{"12+.32"},
			expected: 0,
		}, {
			desc:     "right of another symbol with redundancy",
			input:    []string{"12.+54"},
			expected: 0,
		}, {
			desc:     "right of another symbol with redundancy",
			input:    []string{"12*54"},
			expected: 648,
		}, {
			desc:     "right of another symbol with redundancy",
			input:    []string{"12*..", "...34"},
			expected: 408,
		}, {
			desc:     "right of another symbol with redundancy",
			input:    []string{"12...", ".*...", "42..."},
			expected: 504,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			schematic := &Schematic{}
			for _, v := range tC.input {
				schematic.Accept(v)
			}
			assert.Equal(t, tC.expected, schematic.SolvePt2())
		})
	}
}
