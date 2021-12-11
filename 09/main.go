package main

import (
	"fmt"
	"os"

	"github.com/alperkose/adventofcode2021/input"
)

func main() {
	lavaTubes := &LavaTubes{}
	input.Parse(os.Stdin, lavaTubes)
	fmt.Println(RiskFactorOfHoles(lavaTubes))
	fmt.Println(ThreeLargestBasinsMultiplied(lavaTubes))
}
