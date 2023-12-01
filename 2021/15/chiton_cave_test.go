package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParsesChitonCave(t *testing.T) {
	chitonCave := &ChitonCave{}
	chitonCave.Accept("123")
	chitonCave.Accept("654")
	chitonCave.Accept("789")

	assert.Equal(t, 2, chitonCave.Risk(1, 0))
	assert.Equal(t, 4, chitonCave.Risk(2, 1))
	assert.Equal(t, 7, chitonCave.Risk(0, 2))
}
