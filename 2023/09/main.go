package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alperkose/adventofcode/2023/input"
)

func main() {
	oasis := &Oasis{}
	input.Parse(os.Stdin, oasis)
	fmt.Println("Pt1: ", oasis.SolvePt1())
	fmt.Println("Pt2: ", oasis.SolvePt2())
}

type HistoricalData []int

type Oasis struct {
	measurements []HistoricalData
}

func (gr *Oasis) Accept(inp string) {
	hist := HistoricalData{}
	parts := strings.Split(strings.TrimSpace(inp), " ")
	for _, v := range parts {

		val, _ := strconv.Atoi(v)
		hist = append(hist, val)
	}
	gr.measurements = append(gr.measurements, hist)
}

func (gr *Oasis) SolvePt1() int {
	runningSum := 0
	for _, measurement := range gr.measurements {
		runningSum += measurement.PredictNext()
	}
	return runningSum
}

func (gr *Oasis) SolvePt2() int {
	runningSum := 0
	for _, measurement := range gr.measurements {
		runningSum += measurement.PredictEarlier()
	}
	return runningSum
}

func (hd HistoricalData) PredictNext() int {
	_, last := deriveUntilZeroAndPredict(hd)
	return last
}

func (hd HistoricalData) PredictEarlier() int {
	first, _ := deriveUntilZeroAndPredict(hd)
	return first
}

func deriveUntilZeroAndPredict(hd HistoricalData) (int, int) {
	hd_prime := make(HistoricalData, len(hd)-1)
	allZero := true
	for i := 0; i < len(hd)-1; i++ {
		hd_prime[i] = hd[i+1] - hd[i]
		allZero = allZero && hd_prime[i] == 0
	}

	if allZero {
		return hd[0], hd[len(hd)-1]
	}

	dd_first, dd_last := deriveUntilZeroAndPredict(hd_prime)
	return hd[0] - dd_first, hd[len(hd)-1] + dd_last

}
