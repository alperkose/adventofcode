package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_VerticalReflection(t *testing.T) {
	mirrorClue := &MirrorClues{}
	mirrorClue.Accept("#.##..##.")
	mirrorClue.Accept("..#.##.#.")
	mirrorClue.Accept("##......#")
	mirrorClue.Accept("##......#")
	mirrorClue.Accept("..#.##.#.")
	mirrorClue.Accept("..##..##.")
	mirrorClue.Accept("#.#.##.#.")
	assert.Equal(t, 5, mirrorClue.SolvePt1())
	assert.Equal(t, 300, mirrorClue.SolvePt2())
}

func Test_HorizontalReflection(t *testing.T) {
	mirrorClue := &MirrorClues{}
	mirrorClue.Accept("#...##..#")
	mirrorClue.Accept("#....#..#")
	mirrorClue.Accept("..##..###")
	mirrorClue.Accept("#####.##.")
	mirrorClue.Accept("#####.##.")
	mirrorClue.Accept("..##..###")
	mirrorClue.Accept("#....#..#")
	assert.Equal(t, 400, mirrorClue.SolvePt1())
	assert.Equal(t, 100, mirrorClue.SolvePt2())
}
func Test_SampleInputForPart1(t *testing.T) {
	mirrorClue := &MirrorClues{}
	mirrorClue.Accept("#.##..##.")
	mirrorClue.Accept("..#.##.#.")
	mirrorClue.Accept("##......#")
	mirrorClue.Accept("##......#")
	mirrorClue.Accept("..#.##.#.")
	mirrorClue.Accept("..##..##.")
	mirrorClue.Accept("#.#.##.#.")
	mirrorClue.Accept("")
	mirrorClue.Accept("#...##..#")
	mirrorClue.Accept("#....#..#")
	mirrorClue.Accept("..##..###")
	mirrorClue.Accept("#####.##.")
	mirrorClue.Accept("#####.##.")
	mirrorClue.Accept("..##..###")
	mirrorClue.Accept("#....#..#")
	assert.Equal(t, 405, mirrorClue.SolvePt1())
	assert.Equal(t, 400, mirrorClue.SolvePt2())
}
