package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alperkose/adventofcode/2023/input"
)

func main() {
	calibration := &CalibrationInput{}
	input.Parse(os.Stdin, calibration)
	fmt.Println("Pt1: ", calibration.Resolve())
}

type CalibrationInput struct {
	runningValue int
}

var num_alpha = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func (cal *CalibrationInput) Accept(in string) {

	first := 0
	firstIndex := len(in)
	last := 0
	lastIndex := -1

	for val, digit := range num_alpha {
		loc := strings.Index(in, digit)
		if loc >= 0 && loc < firstIndex {
			firstIndex = loc
			first = val % 10
		}
		loc = strings.LastIndex(in, digit)
		if loc > lastIndex {
			lastIndex = loc
			last = val % 10
		}

	}

	cal.runningValue += first*10 + last
	fmt.Println(in, first*10+last)
}

func (cal *CalibrationInput) Accept_pt1(in string) {
	last := 0
	first := 0

	for i := 0; i < len(in); i++ {
		v, err := strconv.Atoi(string(in[i]))
		if err == nil {
			first = v
			break
		}
	}
	for i := len(in) - 1; i >= 0; i-- {
		v, err := strconv.Atoi(string(in[i]))
		if err == nil {
			last = v
			break
		}
	}

	cal.runningValue += first*10 + last
}

func (cal *CalibrationInput) Resolve() int {
	return cal.runningValue
}
