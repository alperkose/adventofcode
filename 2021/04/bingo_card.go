package main

import (
	"strconv"
)

type Cell struct {
	col int
	row int
}
type BingoCard struct {
	content []string
	rowsize int
	matches map[string]Cell
}

func (card *BingoCard) Bingo(number string) int {
	loc := indexOf(card.content, number)
	if loc < 0 {
		return 0
	}
	cell := Cell{loc % card.rowsize, loc / card.rowsize}
	card.matches[number] = cell

	if card.bingoForCell(cell) {
		return card.score(number)
	}
	return 0
}

func (card *BingoCard) bingoForCell(cell Cell) bool {
	// fmt.Println("=================================")
	// fmt.Printf("checking bingo for cell %v\n", cell)
	// fmt.Printf("state of matching cells %v\n", card.matches)
	// fmt.Printf("contents %v\n", card.content)
	found := true

	for i := 0; i < card.rowsize; i++ {
		_, ok := card.matches[card.content[cell.row*card.rowsize+i]]
		found = found && ok
		// fmt.Printf("[row] %d, matched:%v, found:%v\n", cell.row*card.rowsize+i, ok, found)
		if !found {
			break
		}
	}
	if found {
		return true
	}
	found = true
	for i := 0; i < card.rowsize; i++ {
		_, ok := card.matches[card.content[i*card.rowsize+cell.col]]
		found = found && ok
		// fmt.Printf("[col] %d, matched:%v, found:%v\n", i*card.rowsize+cell.col, ok, found)
	}
	return found
}

func (card *BingoCard) score(multiplier string) int {
	m, _ := strconv.Atoi(multiplier)
	sum := 0
	sumMatched := 0
	for _, v := range card.content {
		nv, _ := strconv.Atoi(v)
		sum += nv
		if _, ok := card.matches[v]; ok {
			sumMatched += nv
		}
	}
	return (sum - sumMatched) * m
}

func indexOf(c []string, e string) int {
	for i, v := range c {
		if v == e {
			return i
		}
	}
	return -1
}
