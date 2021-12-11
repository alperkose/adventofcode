package main

import (
	"math"
	"strconv"
)

type LavaTubes struct {
	heightMap []int
	width     int
	length    int
}

func (lt *LavaTubes) Accept(in string) {
	for _, l := range in {
		v, _ := strconv.Atoi(string(l))
		lt.heightMap = append(lt.heightMap, v)
	}
	lt.width = len(in)
	lt.length += 1
}

func (lt *LavaTubes) Loc(x, y int) int {
	if x < 0 || y < 0 || x >= lt.width || y >= lt.length {
		return math.MaxInt
	}
	return lt.heightMap[x+y*lt.width]
}

func (lt *LavaTubes) Length() int {
	return lt.length
}
func (lt *LavaTubes) Width() int {
	return lt.width
}
