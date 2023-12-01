package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Chunks(t *testing.T) {
	testCases := []struct {
		chunk                   string
		expectedSyntaxScore     int
		expectedCompletionScore int
	}{
		{"[]", 0, 0},
		{"[({})]", 0, 0},
		{"{([(<{}[<>[]}>{[]{[(<()>", 1197, 0},
		{"[[<[([]))<([[{}[[()]]]", 3, 0},
		{"[{[{({}]{}}([{[{{{}}([]", 57, 0},
		{"[<(<(<(<{}))><([]([]()", 3, 0},
		{"<{([([[(<>()){}]>(<<{{", 25137, 0},
		{"[({(<(())[]>[[{[]{<()<>>", 0, 288957},
		{"[(()[<>])]({[<{<<[]>>(", 0, 5566},
		{"(((({<>}<{<{<>}{[]{[]{}", 0, 1480781},
		{"{<[[]]>}<{[{[{[]{()[[[]", 0, 995444},
		{"<{([{{}}[<[[[<>{}]]]>[]]", 0, 294},
	}
	for _, tC := range testCases {
		t.Run(tC.chunk, func(t *testing.T) {
			syntaxScore, completionScore := ChunkScore(tC.chunk)
			assert.Equal(t, tC.expectedSyntaxScore, syntaxScore)
			assert.Equal(t, tC.expectedCompletionScore, completionScore)
		})
	}
}
