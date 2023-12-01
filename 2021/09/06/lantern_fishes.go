package main

import (
	"strconv"
	"strings"
)

type LanternFishes struct {
	cycles []int
}

func (lf *LanternFishes) Accept(in string) {

	cycleStr := strings.Split(in, ",")
	for _, pstr := range cycleStr {
		v, _ := strconv.Atoi(pstr)
		lf.cycles = append(lf.cycles, v)
	}
}

func (lf *LanternFishes) Cycles() []int {
	return lf.cycles
}
