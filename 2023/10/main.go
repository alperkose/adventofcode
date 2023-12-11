package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alperkose/adventofcode/2023/input"
)

func main() {
	oasis := &Schematics{}
	input.Parse(os.Stdin, oasis)
	fmt.Println("Pt1: ", oasis.SolvePt1())
	fmt.Println("Pt2: ", oasis.SolvePt2())
}

type Direction byte

const (
	North Direction = 0
	East  Direction = 1
	South Direction = 2
	West  Direction = 3
)

func (d Direction) String() string {
	switch d {
	case North:
		return "▲"
	case East:
		return "►"
	case South:
		return "▼"
	case West:
		return "◀︎"
	default:
		return "."
	}
}

type Schematics struct {
	loc   []string
	start RowCol
}

func (gr *Schematics) Accept(inp string) {
	start := strings.Index(inp, "S")

	if start >= 0 {
		gr.start = RowCol{len(gr.loc), start}
	}
	gr.loc = append(gr.loc, inp)
}

func (gr *Schematics) SolvePt1() int {
	return len(gr.findThePipe()) / 2
}

func (gr *Schematics) findThePipe() []RowCol {

	var path []RowCol

	startVal := gr.loc[gr.start.row][gr.start.col]
	candidates, directions := Connections(startVal, gr.start)

	for i, cnd := range candidates {
		if cnd.WithinBounds(len(gr.loc), len(gr.loc[0])) && ConnectsTo(startVal, gr.loc[cnd.row][cnd.col], directions[i]) {
			path = gr.followThePipe(gr.start, cnd)
		}
		if len(path) > 0 {
			return path
		}
	}
	return path
}

func (gr *Schematics) followThePipe(start, next RowCol) []RowCol {
	// maxRow, maxCol := len(gr.loc)-1, len(gr.loc[0])-1
	path := []RowCol{start, next}
	var curr, prev RowCol
	for {
		curr = path[len(path)-1]
		prev = path[len(path)-2]
		currVal := gr.loc[curr.row][curr.col]
		// fmt.Printf("%v %v %v - %v\n", prev, curr, string(currVal), path)
		candidates, directions := Connections(currVal, curr)

		for i, cnd := range candidates {
			if cnd.WithinBounds(len(gr.loc), len(gr.loc[0])) && !prev.Equal(cnd) && ConnectsTo(currVal, gr.loc[cnd.row][cnd.col], directions[i]) {
				path = append(path, cnd)
			}
		}

		if path[len(path)-1].Equal(curr) {
			return []RowCol{}
		}
		if path[len(path)-1].Equal(start) {
			return path
		}
	}
}

func (gr *Schematics) SolvePt2() int {
	runningSum := 0
	orderMap := make([][]int, len(gr.loc))
	for i := range orderMap {
		orderMap[i] = make([]int, len(gr.loc[0]))
	}

	thePipe := gr.findThePipe()
	countCCW, countCW := 0, 0
	for i := 1; i < len(thePipe); i++ {
		curr := thePipe[i]
		prev := thePipe[i-1]
		tile := gr.loc[curr.row][curr.col]
		dir := direction(curr, prev, tile)
		if tile == '7' || tile == 'F' {
			if dir == West {
				countCCW++
			} else if dir == East {
				countCW++
			}
		}
		if tile == 'L' || tile == 'J' {
			if dir == East {
				countCCW++
			} else if dir == West {
				countCW++
			}
		}
		orderMap[curr.row][curr.col] = i + 1
	}
	turnCCW := countCCW > countCW
	fmt.Println("left", countCCW, "right", countCW, "CCW", turnCCW)

	fmt.Println("Input and the path:")
	for i, r := range orderMap {
		fmt.Print(gr.loc[i], "  |  ")
		for _, c := range r {
			if c > 0 {
				fmt.Print(direction(thePipe[c-1], thePipe[c-2], gr.loc[thePipe[c-1].row][thePipe[c-1].col]))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i < len(orderMap); i++ {
		row := orderMap[i]
		for j := 0; j < len(row); j++ {
			// fmt.Println("checking", i, j, orderMap[i][j])
			if orderMap[i][j] > 0 {
				continue
			}
			dr := i + 1
			for ; dr < len(orderMap) && orderMap[dr][j] == 0; dr++ {
			}
			if dr == len(orderMap) {
				continue
			}
			foundPipe := thePipe[orderMap[dr][j]-1]
			beforeFoundPipe := thePipe[orderMap[dr][j]-2]
			dir := direction(foundPipe, beforeFoundPipe, gr.loc[dr][j])
			// fmt.Println("for", i, j, foundPipe, beforeFoundPipe, dir)
			if (turnCCW && dir == East) || (!turnCCW && dir == West) {
				// fmt.Println("inside")
				runningSum++
			}
		}
	}

	return runningSum
}

func direction(pipe, prevpipe RowCol, tile byte) Direction {
	if pipe.row == prevpipe.row {
		if pipe.col < prevpipe.col {
			return West
		}
		return East
	} else if pipe.row < prevpipe.row {
		if tile == 'F' {
			return East
		} else if tile == '7' {
			return West
		} else {
			return North
		}

	} else if pipe.row > prevpipe.row {
		if tile == 'L' {
			return East
		} else if tile == 'J' {
			return West
		} else {
			return South
		}
	}
	return South
}

func ConnectsTo(from, to byte, direction Direction) bool {

	switch from {
	case '|':
		return (direction == North && (to == 'S' || to == '7' || to == '|' || to == 'F')) || (direction == South && (to == 'S' || to == 'J' || to == '|' || to == 'L'))
	case '-':
		return (direction == East && (to == 'S' || to == 'J' || to == '-' || to == '7')) || (direction == West && (to == 'S' || to == 'L' || to == '-' || to == 'F'))
	case '7':
		return (direction == South && (to == 'S' || to == 'J' || to == '|' || to == 'L')) || (direction == West && (to == 'S' || to == 'L' || to == '-' || to == 'F'))
	case 'J':
		return (direction == North && (to == 'S' || to == '7' || to == '|' || to == 'F')) || (direction == West && (to == 'S' || to == 'L' || to == '-' || to == 'F'))
	case 'L':
		return (direction == North && (to == 'S' || to == '7' || to == '|' || to == 'F')) || (direction == East && (to == 'S' || to == 'J' || to == '-' || to == '7'))
	case 'F':
		return (direction == South && (to == 'S' || to == 'J' || to == '|' || to == 'L')) || (direction == East && (to == 'S' || to == 'J' || to == '-' || to == '7'))
	case 'S':
		return (direction == North && (to == '7' || to == '|' || to == 'F')) || (direction == East && (to == 'J' || to == '-' || to == '7')) || (direction == South && (to == 'J' || to == '|' || to == 'L')) || (direction == West && (to == 'L' || to == '-' || to == 'F'))
	default:
		return false
	}
}

func Connections(tile byte, rc RowCol) ([]RowCol, []Direction) {
	switch tile {
	case 'S':
		return []RowCol{{rc.row - 1, rc.col}, {rc.row, rc.col - 1}, {rc.row + 1, rc.col}, {rc.row, rc.col + 1}}, []Direction{North, West, South, East}
	case '|':
		return []RowCol{{rc.row - 1, rc.col}, {rc.row + 1, rc.col}}, []Direction{North, South}
	case '-':
		return []RowCol{{rc.row, rc.col - 1}, {rc.row, rc.col + 1}}, []Direction{West, East}
	case '7':
		return []RowCol{{rc.row, rc.col - 1}, {rc.row + 1, rc.col}}, []Direction{West, South}
	case 'J':
		return []RowCol{{rc.row - 1, rc.col}, {rc.row, rc.col - 1}}, []Direction{North, West}
	case 'L':
		return []RowCol{{rc.row - 1, rc.col}, {rc.row, rc.col + 1}}, []Direction{North, East}
	case 'F':
		return []RowCol{{rc.row + 1, rc.col}, {rc.row, rc.col + 1}}, []Direction{South, East}
	default:
		return []RowCol{}, []Direction{}
	}
}

type RowCol struct {
	row, col int
}

func (rc RowCol) Equal(that RowCol) bool {
	return rc.col == that.col && rc.row == that.row
}

func (rc RowCol) WithinBounds(rowBoundary, colBoundary int) bool {
	return rc.row >= 0 && rc.row < rowBoundary && rc.col >= 0 && rc.col < colBoundary
}
