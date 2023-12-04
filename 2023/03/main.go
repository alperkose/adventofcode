package main

import (
	"fmt"
	"os"

	"github.com/alperkose/adventofcode/2023/input"
)

func main() {
	schematic := &Schematic{}
	input.Parse(os.Stdin, schematic)
	fmt.Println("Pt1: ", schematic.Solve())
	fmt.Println("Pt2: ", schematic.SolvePt2())
}

type Schematic struct {
	numbers [][]PartNumber
	symbols [][]Symbol
}

func (s *Schematic) Accept(inp string) {
	numbersInLine := []PartNumber{}
	symbolsInLine := []Symbol{}
	val := 0
	begin := -1
	for ind, ch := range inp {
		isDigit := ch >= '0' && ch <= '9'
		if isDigit {
			val = val*10 + int(ch-'0')
			if begin < 0 {
				begin = ind
			}
		} else if begin >= 0 {
			numbersInLine = append(numbersInLine, PartNumber{len(s.numbers), begin, ind - 1, val})
			begin = -1
			val = 0
		}

		if !isDigit && ch != '.' {
			symbolsInLine = append(symbolsInLine, Symbol{ind, ch})
		}
	}
	if begin >= 0 {
		numbersInLine = append(numbersInLine, PartNumber{len(s.numbers), begin, len(inp) - 1, val})
		begin = -1
		val = 0
	}
	s.numbers = append(s.numbers, numbersInLine)
	s.symbols = append(s.symbols, symbolsInLine)
}

func (s *Schematic) Solve() int {
	runningSum := 0
	for line, symbolsInLine := range s.symbols {
		for _, sym := range symbolsInLine {
			aboveLine := line - 1
			belowLine := line + 1
			if line == 0 {
				aboveLine = line
			}
			if belowLine == len(s.numbers) {
				belowLine = line
			}
			for l := aboveLine; l <= belowLine; l++ {
				for _, num := range s.numbers[l] {
					if num.Adjacent_to(line, sym.pos) {
						runningSum += num.value
					}
				}
			}

		}
	}
	return runningSum
}

func (s *Schematic) SolvePt2() int {
	runningSum := 0
	for line, symbolsInLine := range s.symbols {
		for _, sym := range symbolsInLine {
			if sym.char != '*' {
				continue
			}
			aboveLine := line - 1
			belowLine := line + 1
			if line == 0 {
				aboveLine = line
			}
			if belowLine == len(s.numbers) {
				belowLine = line
			}

			ratio := 1
			adjacencyCount := 0
			for l := aboveLine; l <= belowLine; l++ {
				for _, num := range s.numbers[l] {
					if num.Adjacent_to(line, sym.pos) {
						adjacencyCount++
						ratio *= num.value
					}
				}
			}
			if adjacencyCount == 2 {
				runningSum += ratio
			}

		}
	}
	return runningSum
}

type PartNumber struct {
	line  int
	begin int
	end   int
	value int
}

func (pn PartNumber) Adjacent_to(line, pos int) bool {
	return line >= (pn.line-1) && line <= (pn.line+1) && pos >= pn.begin-1 && pos <= pn.end+1
}

type Symbol struct {
	pos  int
	char rune
}
