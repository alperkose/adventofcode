package main

import "math"

func IncreasingFuelConsumption(locs []int, to int) int {
	sum := 0
	for _, v := range locs {
		dist := int(math.Abs(float64(v - to)))
		sum += dist * (dist + 1) / 2
	}
	// fmt.Println("Consumption", locs, to, sum)
	return sum
}
