package main

import (
	"math"
	"strconv"
)

type ChitonCave struct {
	risk   []int
	width  int
	length int
}

func (lt *ChitonCave) Accept(in string) {
	for _, l := range in {
		v, _ := strconv.Atoi(string(l))
		lt.risk = append(lt.risk, v)
	}
	lt.width = len(in)
	lt.length += 1
}

func (lt *ChitonCave) Risks() []int {
	return lt.risk
}

func (lt *ChitonCave) Risk(x, y int) int {
	if x < 0 || y < 0 || x >= lt.width || y >= lt.length {
		return math.MaxInt
	}
	return lt.risk[x+y*lt.width]
}

func (lt *ChitonCave) Width() int {
	return lt.width
}
