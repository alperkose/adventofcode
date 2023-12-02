package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alperkose/adventofcode/2023/input"
)

type GameSession []Game

type GameParser struct {
	Games GameSession
}

type Game struct {
	id      int
	reveals []Reveal
}

type Reveal map[string]int

func (gp *GameParser) Accept(inp string) {
	parts := strings.Split(inp, ":")
	gameId, _ := strconv.Atoi(parts[0][5:])

	reveals := []Reveal{}

	revealStrings := strings.Split(strings.TrimSpace(parts[1]), ";")
	for _, rstr := range revealStrings {
		groups := strings.Split(strings.TrimSpace(rstr), ",")
		reveal := Reveal{}
		reveals = append(reveals, reveal)
		for _, grp := range groups {
			set := strings.Split(strings.TrimSpace(grp), " ")
			cnt, _ := strconv.Atoi(set[0])
			reveal[set[1]] = cnt
		}
	}

	gp.Games = append(gp.Games, Game{gameId, reveals})
}

type Solver interface {
	Solve(GameSession) int
}

func main() {
	parser := &GameParser{}
	input.Parse(os.Stdin, parser)
	fmt.Println("Pt1: ", SolverForPart1().Solve(parser.Games))
	fmt.Println("Pt2: ", SolverForPart2().Solve(parser.Games))
}
