package main

import (
	"fmt"
	"os"

	"github.com/alperkose/adventofcode2021/input"
)

func main() {

	diagnosticReport := &ByteContainer{}
	input.Parse(os.Stdin, diagnosticReport)
	mask := uint32(0xFFFFFFFF) >> (32 - diagnosticReport.Digits())
	commonBit := GammaRate(diagnosticReport.values)
	fmt.Printf("%5d (%b), %5d (%b), %5d (%b)\n", commonBit, commonBit, ^commonBit&mask, ^commonBit&mask, commonBit*(^commonBit&mask), commonBit*((^commonBit)&mask))

	o2 := O2GenerationRating(diagnosticReport.values, uint32(diagnosticReport.Digits()))
	co2 := CO2ScrubberRating(diagnosticReport.values, uint32(diagnosticReport.Digits()))

	fmt.Printf("%5d (%b), %5d (%b), %5d (%b)\n", o2, o2, co2, co2, o2*co2, o2*co2)
}
