package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParsesDumboOctopuses(t *testing.T) {
	dumboOctopuses := &DumboOctopuses{}
	dumboOctopuses.Accept("123")
	dumboOctopuses.Accept("654")
	dumboOctopuses.Accept("789")

	assert.Equal(t, 2, dumboOctopuses.Energy(1, 0))
	assert.Equal(t, 4, dumboOctopuses.Energy(2, 1))
	assert.Equal(t, 7, dumboOctopuses.Energy(0, 2))
}
