package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExtendOneByOneCaveBy5(t *testing.T) {

	newCave, newWidth := ExtendCave([]int{8}, 1, 5)

	assert.Equal(t, 5, newWidth)
	assert.Equal(t, []int{8, 9, 1, 2, 3, 9, 1, 2, 3, 4, 1, 2, 3, 4, 5, 2, 3, 4, 5, 6, 3, 4, 5, 6, 7}, newCave)
}
func Test_ExtendTwoByTwoCaveBy5(t *testing.T) {

	newCave, newWidth := ExtendCave([]int{8, 9, 3, 4}, 2, 5)

	assert.Equal(t, 10, newWidth)
	expectedCave := []int{8, 9, 9, 1, 1, 2, 2, 3, 3, 4,
		3, 4, 4, 5, 5, 6, 6, 7, 7, 8,
		9, 1, 1, 2, 2, 3, 3, 4, 4, 5,
		4, 5, 5, 6, 6, 7, 7, 8, 8, 9,
		1, 2, 2, 3, 3, 4, 4, 5, 5, 6,
		5, 6, 6, 7, 7, 8, 8, 9, 9, 1,
		2, 3, 3, 4, 4, 5, 5, 6, 6, 7,
		6, 7, 7, 8, 8, 9, 9, 1, 1, 2,
		3, 4, 4, 5, 5, 6, 6, 7, 7, 8,
		7, 8, 8, 9, 9, 1, 1, 2, 2, 3,
	}
	assert.Equal(t, expectedCave, newCave)
}
