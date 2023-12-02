package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alperkose/adventofcode2021/input"
)

func main() {
	lanternFishes := &LanternFishes{}
	input.Parse(os.Stdin, lanternFishes)
	// lanternFishes.Accept("6")
	execTime := time.Now()
	pop := Simulate(lanternFishes, 80)
	fmt.Println("Population after 80 days", pop, time.Since(execTime))

	execTime = time.Now()
	pop = Simulate(lanternFishes, 256)
	fmt.Println("Population after 256 days", pop, time.Since(execTime))

	// lanternFishes.Accept("6")
	// for i := 0; i < 128; i += 7 {
	// 	pop := Simulate(lanternFishes, i)
	// 	fmt.Println("Population after", i, "days", pop)
	// }
}
