package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/alperkose/adventofcode/2023/input"
)

func main() {
	scratchCards := &ScratchCards{}
	input.Parse(os.Stdin, scratchCards)
	fmt.Println("Pt1: ", scratchCards.SolvePt1())
	fmt.Println("Pt2: ", scratchCards.SolvePt2())
}

type ScratchCards struct {
	cards []ScratchCard
}

type ScratchCard struct {
	numbers []int
}

func (s *ScratchCards) Accept(inp string) {
	nums := []int{}
	parts := strings.Split(inp[8:], " ")
	for _, p := range parts {
		v, err := strconv.Atoi(p)
		if err == nil {
			nums = append(nums, v)
		}
	}
	s.cards = append(s.cards, ScratchCard{nums})
}

func (s ScratchCards) SolvePt1() int {
	runningSum := 0
	for _, sc := range s.cards {
		nums := sc.numbers
		sort.Ints(nums)
		points := 1
		for i := 1; i < len(nums); i++ {
			if nums[i-1] == nums[i] {
				points = points << 1
			}
		}
		runningSum += points >> 1
	}

	return runningSum
}

func (s ScratchCards) SolvePt2() int {
	cards := make([]int, len(s.cards))
	runningSum := 0
	for ind, sc := range s.cards {
		cards[ind] += 1
		runningSum += cards[ind]
		// fmt.Println(runningSum, ind, cards[ind])
		nums := sc.numbers
		sort.Ints(nums)
		count := 0
		for i := 1; i < len(nums); i++ {
			if nums[i-1] == nums[i] {
				count++
				if ind+count < len(cards) {
					cards[ind+count] += cards[ind]
				}
			}
		}
	}

	return runningSum
}
