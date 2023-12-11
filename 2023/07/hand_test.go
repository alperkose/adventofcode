package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BeatingHands(t *testing.T) {
	testCases := []struct {
		desc                  string
		higherHand, lowerHand Hand
	}{
		{"Two pair beats one pair", NewHand("KK677", 765), NewHand("T55J3", 684)},
		{"Three of a kind beats a pair", NewHand("T55J5", 684), NewHand("32T3K", 765)},
		{"Three of a kind beats two pairs", NewHand("T55J5", 684), NewHand("32K3K", 765)},
		{"Full house beats one pair", NewHand("23332", 684), NewHand("38T8K", 765)},
		{"Full house beats two pairs", NewHand("AAA22", 684), NewHand("6T6KK", 765)},
		{"Full house beats three of a kind", NewHand("23332", 684), NewHand("33T3K", 765)},
		{"Four of a kind one pair", NewHand("AA8AA", 684), NewHand("27342", 765)},
		{"Four of a kind two pairs", NewHand("AA8AA", 684), NewHand("449TT", 765)},
		{"Four of a kind three of a kind", NewHand("AA8AA", 684), NewHand("Q3Q9Q", 765)},
		{"Four of a kind beats Fullhouse", NewHand("AA8AA", 684), NewHand("JKJKJ", 765)},
		{"Five of a kind beats one pair", NewHand("AAAAA", 684), NewHand("J523J", 765)},
		{"Five of a kind beats two pairs", NewHand("AAAAA", 684), NewHand("6K22K", 765)},
		{"Five of a kind beats three of a kind", NewHand("AAAAA", 684), NewHand("32QQQ", 765)},
		{"Five of a kind beats full house", NewHand("AAAAA", 684), NewHand("5TT5T", 765)},
		{"Five of a kind beats a four of a kind", NewHand("AAAAA", 684), NewHand("3333K", 765)},
		{"Tie, first higher card wins", NewHand("KK657", 765), NewHand("T55J3", 684)},
		{"Tie, second higher card wins", NewHand("Q6343", 765), NewHand("Q55J3", 684)},
		{"Tie, third higher card wins", NewHand("Q6545", 765), NewHand("Q6443", 684)},
		{"Tie, fourth higher card wins", NewHand("44377", 765), NewHand("44366", 684)},
		{"Tie, fifth higher card wins", NewHand("AAAAA", 765), NewHand("KKKKK", 684)},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// fmt.Printf("%v %v %v %v\n", tC.higherHand, tC.lowerHand, tC.higherHand.Beats(tC.lowerHand), tC.lowerHand.Beats(tC.higherHand))
			assert.True(t, tC.higherHand.Beats(tC.lowerHand))
			assert.False(t, tC.lowerHand.Beats(tC.higherHand))
		})
	}
}

func Test_weird(t *testing.T) {
	assert.False(t, NewHand("T55J3", 684).Beats(NewHand("KK657", 765)))
	assert.False(t, NewHand("Q55J3", 684).Beats(NewHand("Q6343", 765)))
}
