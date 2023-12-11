package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alperkose/adventofcode/2023/input"
)

func main() {
	raceRecords := &RaceRecords{}
	input.Parse(os.Stdin, raceRecords)
	fmt.Println("Pt1: ", raceRecords.SolvePt1())
}

type RaceRecords struct {
	duration []int
	record   []int
}

func (gr *RaceRecords) Accept(inp string) {
	if strings.HasPrefix(inp, "Time:") {
		parts := strings.Split(strings.TrimSpace(inp[5:]), " ")
		for _, str := range parts {
			val, err := strconv.Atoi(str)
			if err == nil {
				gr.duration = append(gr.duration, val)
			}
		}
	} else if strings.HasPrefix(inp, "Distance:") {
		parts := strings.Split(strings.TrimSpace(inp[9:]), " ")
		for _, str := range parts {
			val, err := strconv.Atoi(str)
			if err == nil {
				gr.record = append(gr.record, val)
			}
		}
	}
}

func (gr *RaceRecords) SolvePt1() int {
	runningSum := 1
	for i := 0; i < len(gr.duration); i++ {
		passingCount := 0
		fmt.Printf("%d - %+v\n", i, gr)
		for j := 1; j < gr.duration[i]; j++ {
			// fmt.Printf("  > %d - %2d/%2d-%2d | %v >= %2d\n", j, gr.record[i], gr.duration[i], j, (j*(gr.duration[i]-j)) > gr.record[i], j)
			if (j * (gr.duration[i] - j)) > gr.record[i] {
				passingCount++
			}
		}
		if passingCount > 0 {
			runningSum *= passingCount
		}
	}
	return runningSum
}
