package main

import (
	"fmt"
	"os"

	"github.com/alperkose/adventofcode2021/input"
)

func main() {
	depths := &input.IntContainer{}
	input.Parse(os.Stdin, depths)
	fmt.Println(CountDepthIncrease(depths.Values()))
	fmt.Println(SlidingWindowOfDepthIncreases(depths.Values()))
}
