package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParsesHeightmap(t *testing.T) {
	lavaTubes := &LavaTubes{}
	lavaTubes.Accept("123")
	lavaTubes.Accept("654")
	lavaTubes.Accept("789")

	assert.Equal(t, 2, lavaTubes.Loc(1, 0))
	assert.Equal(t, 4, lavaTubes.Loc(2, 1))
	assert.Equal(t, 7, lavaTubes.Loc(0, 2))
}
