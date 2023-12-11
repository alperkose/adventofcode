package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/alperkose/adventofcode/2023/input"
)

func main() {
	camelCards := &CamelCards{}
	input.Parse(os.Stdin, camelCards)
	fmt.Println("Pt1: ", camelCards.SolvePt1())
}

type Hand struct {
	cards string
	bid   int
	v     int
}

func NewHand(cards string, bid int) Hand {

	groups := make([]int, 13)
	for i := 0; i < len(cards); i++ {
		groups[cardOrder(cards[i])] += 1
	}

	v := 1
	for _, c := range groups {
		switch c {
		case 2:
			v += 10
		case 3:
			v += 100
		case 4:
			v += 1000
		case 5:
			v += 10000
		}

	}

	hand := Hand{cards, bid, v}
	return hand
}

func (my Hand) Beats(that Hand) bool {
	// fmt.Printf("%v\t%v\n", my, that)
	if my.v == that.v {
		var myNthCard, thatNthCard int
		for i := 0; i < len(my.cards); i++ {
			myNthCard = cardOrder(my.cards[i])
			thatNthCard = cardOrder(that.cards[i])
			if myNthCard != thatNthCard {
				return myNthCard > thatNthCard
			}
		}
	}
	return my.v > that.v
}

func cardOrder(c byte) int {
	switch {
	case c == 'A':
		return 12
	case c == 'K':
		return 11
	case c == 'Q':
		return 10
	case c == 'J':
		return 9
	case c == 'T':
		return 8
	case c >= '2' || c <= 9:
		return int(c - '2')
	}
	return -1
}

type CamelCards struct {
	hands []Hand
}

func (gr *CamelCards) Accept(inp string) {
	cards := inp[0:5]
	bid, _ := strconv.Atoi(inp[6:])

	gr.hands = append(gr.hands, NewHand(cards, bid))
}

func (gr *CamelCards) SolvePt1() int {
	sort.Slice(gr.hands, func(i, j int) bool {
		return gr.hands[j].Beats(gr.hands[i])
	})
	runningValue := 0
	for ind, hand := range gr.hands {
		runningValue += (ind + 1) * hand.bid
	}
	// fmt.Printf("%+v\n", gr.hands)
	return runningValue
}
