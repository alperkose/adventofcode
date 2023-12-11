package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SampeleInputForPart1(t *testing.T) {
	camelCards := &CamelCards{}
	camelCards.Accept("32T3K 765")
	camelCards.Accept("T55J5 684")
	camelCards.Accept("KK677 28")
	camelCards.Accept("KTJJT 220")
	camelCards.Accept("QQQJA 483")

	assert.Equal(t, 6440, camelCards.SolvePt1())
}

func Test_Deletethis(t *testing.T) {
	camelCards := &CamelCards{}
	camelCards.Accept("23455 765")
	camelCards.Accept("234KJ 684")

	assert.Equal(t, 765+2*684, camelCards.SolvePt1())
}
