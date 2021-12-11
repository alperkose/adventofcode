package main

import (
	"fmt"
	"os"

	"github.com/alperkose/adventofcode2021/input"
)

func main() {
	generatedBingo := &GeneratedBingo{}
	input.Parse(os.Stdin, generatedBingo)
	fmt.Println(generatedBingo.FirstBingoScore())
	fmt.Println(generatedBingo.LastBingoScore())
}
