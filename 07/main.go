package main

import (
	"fmt"
	"os"

	"github.com/alperkose/adventofcode2021/input"
)

func main() {
	crabMarines := &CrabMarines{}
	input.Parse(os.Stdin, crabMarines)
	alignment := LocalMinima(crabMarines.Placements(), cached(fuelConsumption))
	fmt.Println(alignment, fuelConsumption(crabMarines.Placements(), alignment))

	alignment = LocalMinima(crabMarines.Placements(), cached(IncreasingFuelConsumption))
	fmt.Println(alignment, IncreasingFuelConsumption(crabMarines.Placements(), alignment))

	// for i := 0; i <= 17; i++ {
	// 	fmt.Println(i, IncreasingFuelConsumption(crabMarines.Placements(), i))
	// }
}
