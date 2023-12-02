package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParsingASingleLine(t *testing.T) {
	parser := &GameParser{}
	parser.Accept("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	assert.Equal(t, 1, len(parser.Games))
	assert.Equal(t, 1, parser.Games[0].id)
	assert.Equal(t, 3, len(parser.Games[0].reveals))
	assert.Equal(t, 3, parser.Games[0].reveals[0]["blue"])
}

func Test_ParsingAMultipleLine(t *testing.T) {
	parser := &GameParser{}
	parser.Accept("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	parser.Accept("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green")
	assert.Equal(t, 2, len(parser.Games))
	assert.Equal(t, 3, parser.Games[1].id)
	assert.Equal(t, 2, len(parser.Games[1].reveals))
	assert.Equal(t, 13, parser.Games[1].reveals[1]["green"])
}
