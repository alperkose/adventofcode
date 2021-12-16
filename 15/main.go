package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alperkose/adventofcode2021/input"
)

func main() {
	cave := &ChitonCave{}
	input.Parse(os.Stdin, cave)
	fmt.Println("Part1", LowestRiskPath(cave))

	execTime := time.Now()
	pop := LowestRiskPathOfExtended(cave, 1)
	fmt.Println("Lowest risk extension by 1", pop, time.Since(execTime))

	execTime = time.Now()
	pop = LowestRiskPathOfExtended(cave, 2)
	fmt.Println("Lowest risk extension by 2", pop, time.Since(execTime))

	execTime = time.Now()
	pop = LowestRiskPathOfExtended(cave, 3)
	fmt.Println("Lowest risk extension by 3", pop, time.Since(execTime))

	execTime = time.Now()
	pop = LowestRiskPathOfExtended(cave, 4)
	fmt.Println("Lowest risk extension by 4", pop, time.Since(execTime))

	execTime = time.Now()
	pop = LowestRiskPathOfExtended(cave, 5)
	fmt.Println("Lowest risk extension by 5", pop, time.Since(execTime))

	// fmt.Println("Part2", LowestRiskPathOfExtended(cave, 5))
}
