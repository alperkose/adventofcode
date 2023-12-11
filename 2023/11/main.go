package main

import (
	"fmt"
	"math"
	"os"

	"github.com/alperkose/adventofcode/2023/input"
)

func main() {
	observation := &Observation{}
	input.Parse(os.Stdin, observation)
	fmt.Println("Pt1: ", observation.SolvePt1())
	fmt.Println("Pt2: ", observation.SolvePt2())
}

type Observation struct {
	image    []string
	galaxies []RowCol
}

func (ob *Observation) Accept(inp string) {
	// for i, sp := range inp {
	// 	if sp == '#' {
	// 		ob.galaxies = append(ob.galaxies, RowCol{len(ob.image), i})
	// 	}
	// }
	ob.image = append(ob.image, inp)
}

func (ob *Observation) SolvePt1() int {
	runningSum := 0
	emptyCol := make([]bool, len(ob.image[0]))
	emptyRow := make([]bool, len(ob.image))
	for i, _ := range emptyCol {
		emptyCol[i] = true
	}
	for i := 0; i < len(ob.image); i++ {
		emptyRow[i] = true
		for j := 0; j < len(ob.image[0]); j++ {
			emptyCol[j] = emptyCol[j] && ob.image[i][j] == '.'
			emptyRow[i] = emptyRow[i] && ob.image[i][j] == '.'
		}
	}

	fmt.Println(emptyRow)
	fmt.Println(emptyCol)

	galaxies := []RowCol{}
	deltaRow, deltaCol := 0, 0
	for i := 0; i < len(ob.image); i++ {
		if emptyRow[i] {
			deltaRow += 999999
		}
		deltaCol = 0
		for j := 0; j < len(ob.image[0]); j++ {
			if emptyCol[j] {
				deltaCol += 999999
			}

			if ob.image[i][j] == '#' {
				galaxies = append(galaxies, RowCol{i + deltaRow, j + deltaCol})
			}
		}
	}

	for i := 0; i < len(galaxies); i++ {
		for j := i; j < len(galaxies); j++ {
			runningSum += int(math.Abs(float64(galaxies[i].col-galaxies[j].col)) + math.Abs(float64(galaxies[i].row-galaxies[j].row)))
		}
	}

	return runningSum
}

func (ob *Observation) SolvePt2() int {
	runningSum := 0

	return runningSum
}

type RowCol struct {
	row, col int
}

func (rc RowCol) Equal(that RowCol) bool {
	return rc.col == that.col && rc.row == that.row
}
