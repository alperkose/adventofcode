package main

import (
	"fmt"
	"os"

	"github.com/alperkose/adventofcode2021/input"
)

func main() {
	dumboOctopuses := &DumboOctopuses{}
	input.Parse(os.Stdin, dumboOctopuses)
	// fmt.Println(FlashesAfterSteps(dumboOctopuses, 100))
	fmt.Println(SynchronousFlashStep(dumboOctopuses))
}
