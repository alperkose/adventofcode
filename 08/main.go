package main

import (
	"fmt"
	"os"

	"github.com/alperkose/adventofcode2021/input"
)

func main() {
	observations := &Observations{}
	input.Parse(os.Stdin, observations)
	sum := 0
	for i, entry := range observations.entries {
		sum += entry.DigitValues()
		fmt.Printf("%3d) +%4d -> %d\n", i, entry.DigitValues(), sum)
	}
}
