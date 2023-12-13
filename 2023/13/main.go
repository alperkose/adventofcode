package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alperkose/adventofcode/2023/input"
)

func main() {
	observation := &MirrorClues{}
	input.Parse(os.Stdin, observation)
	fmt.Println("Pt1: ", observation.SolvePt1())
	fmt.Println("Pt2: ", observation.SolvePt2())
}

type MirrorClues struct {
	clues []MirrorClue
}

type MirrorClue struct {
	rows []string
}

func (cr *MirrorClues) Accept(inp string) {
	trimmedInput := strings.TrimSpace(inp)
	if len(cr.clues) == 0 && len(trimmedInput) > 0 {
		cr.clues = append(cr.clues, MirrorClue{[]string{inp}})
		return
	}
	if len(trimmedInput) == 0 {
		cr.clues = append(cr.clues, MirrorClue{[]string{}})
	} else {
		cr.clues[len(cr.clues)-1].rows = append(cr.clues[len(cr.clues)-1].rows, trimmedInput)
	}
}

func (cr *MirrorClues) SolvePt1() int {
	runningSum := 0

	for _, clue := range cr.clues {
		col := findColReflection(clue.rows, 0)
		if col > 0 {
			runningSum += col
		}
		row := findRowReflection(clue.rows, 0)
		if row > 0 {
			runningSum += row * 100
		}

	}
	return runningSum
}

func findColReflection(clue []string, expectedApproximation int) int {
	approximation := 0
	for i := 1; i < len(clue[0]); i++ {
		approximation = compareCols(clue, i, i-1)
		for j := 1; approximation <= expectedApproximation && i-j-1 >= 0 && i+j < len(clue[0]); j++ {
			approximation += compareCols(clue, i-j-1, i+j)
		}
		if approximation == expectedApproximation {
			return i
		}
	}
	return -1
}

func compareCols(clue []string, col1, col2 int) int {
	approximation := 0
	for _, row := range clue {
		if row[col1] != row[col2] {
			approximation++
		}
	}
	return approximation
}

func findRowReflection(clue []string, expectedApproximation int) int {
	approximation := 0
	for i := 1; i < len(clue); i++ {
		approximation = compareRows(clue[i], clue[i-1])
		for j := 1; approximation <= expectedApproximation && i-j-1 >= 0 && i+j < len(clue); j++ {
			approximation += compareRows(clue[i+j], clue[i-j-1])
		}
		if approximation == expectedApproximation {
			return i
		}
	}
	return -1
}

func compareRows(row1, row2 string) int {
	approximation := 0
	for i := 0; i < len(row1); i++ {
		if row1[i] != row2[i] {
			approximation++
		}
	}
	return approximation
}

func (cr *MirrorClues) SolvePt2() int {
	runningSum := 0

	for _, clue := range cr.clues {
		col := findColReflection(clue.rows, 1)
		if col > 0 {
			runningSum += col
		} else {
			row := findRowReflection(clue.rows, 1)
			if row > 0 {
				runningSum += row * 100
			}
		}
	}
	return runningSum
}
