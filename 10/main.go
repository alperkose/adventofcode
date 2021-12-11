package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/alperkose/adventofcode2021/input"
)

func main() {
	nav := &NavigationSubSystem{}
	input.Parse(os.Stdin, nav)

	sumSyntaxScore := 0
	complScores := []int{}
	for _, c := range nav.chunks {
		syn, compl := ChunkScore(c)
		sumSyntaxScore += syn
		// fmt.Println(c, syn)
		if compl > 0 {
			complScores = append(complScores, compl)
		}
	}
	fmt.Println("sum", sumSyntaxScore)
	sort.Ints(complScores)
	fmt.Println(complScores[len(complScores)/2])
}
