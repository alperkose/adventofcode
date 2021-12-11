package main

import (
	"fmt"
	"math"
	"sort"
)

type MinimaFitness func([]int, int) int

func LocalMinima(values []int, minima MinimaFitness) int {
	sortedValues := make([]int, len(values))
	copy(sortedValues, values)
	sort.Ints(sortedValues)

	alignment := sortedValues[len(sortedValues)/2]
	for {
		mLowEnd := minima(values, alignment-1)
		mAlignment := minima(values, alignment)
		mHighEnd := minima(values, alignment+1)
		fmt.Println("iteration", alignment, mLowEnd, mAlignment, mHighEnd)
		if mAlignment < mLowEnd && mAlignment < mHighEnd {
			return alignment
		}

		if mLowEnd < mAlignment {
			alignment -= 1
			if alignment < sortedValues[0] {
				return alignment + 1
			}
		}
		if mHighEnd < mAlignment {
			alignment += 1
			if alignment > sortedValues[len(sortedValues)-1] {
				return alignment - 1
			}

		}

	}
}

func fuelConsumption(locs []int, to int) int {
	sum := 0
	for _, v := range locs {
		sum += int(math.Abs(float64(v - to)))
	}
	// fmt.Println("Consumption", locs, to, sum)
	return sum
}

func cached(mf MinimaFitness) MinimaFitness {
	cache := map[int]int{}
	return func(v []int, to int) int {
		m, ok := cache[to]
		if ok {
			return m
		}
		m = mf(v, to)
		cache[to] = m
		return m
	}
}
