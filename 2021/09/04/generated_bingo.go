package main

import (
	"strings"
)

type GeneratedBingo struct {
	draws []string
	cards []*BingoCard
}

func (gb *GeneratedBingo) Accept(in string) {
	if len(gb.draws) == 0 {
		gb.draws = strings.Split(in, ",")
		return
	}

	if len(strings.Trim(in, " ")) == 0 {
		gb.cards = append(gb.cards, &BingoCard{[]string{}, 0, map[string]Cell{}})
		return
	}

	currentCard := gb.cards[len(gb.cards)-1]
	newRow := strings.Split(in, " ")
	rowSize := 0
	for _, v := range newRow {
		if len(v) > 0 {
			currentCard.content = append(currentCard.content, v)
			rowSize++
		}
	}
	currentCard.rowsize = rowSize
}

func (gb *GeneratedBingo) Draws() []string {
	return gb.draws
}

func (gb *GeneratedBingo) BingoCards() []*BingoCard {
	return gb.cards
}

func (gb *GeneratedBingo) FirstBingoScore() int {
	for i := 0; i < len(gb.draws); i++ {
		for j := 0; j < len(gb.cards); j++ {
			score := gb.cards[j].Bingo(gb.draws[i])
			if score > 0 {
				return score
			}
		}
	}
	return 0
}

func (gb *GeneratedBingo) LastBingoScore() int {
	lastScore := 0
	bingoedCard := make([]bool, len(gb.cards))
	for i := 0; i < len(gb.draws); i++ {
		for j := 0; j < len(gb.cards); j++ {
			if bingoedCard[j] {
				continue
			}
			score := gb.cards[j].Bingo(gb.draws[i])
			if score > 0 {
				bingoedCard[j] = true
				lastScore = score
			}
		}
	}
	return lastScore
}
