package main

import (
	"fmt"
	"os"

	"github.com/alperkose/adventofcode2021/input"
)

func main() {
	htvMap := &HypoThermalVentMap{}
	input.Parse(os.Stdin, htvMap)
	fmt.Println(htvMap.Clouds().NumberOfIntersections())
}
