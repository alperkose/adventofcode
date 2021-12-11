package main

import (
	"math"
	"strconv"
)

type DumboOctopuses struct {
	octopusEnergies []int
	width           int
	length          int
}

func (lt *DumboOctopuses) Accept(in string) {
	for _, l := range in {
		v, _ := strconv.Atoi(string(l))
		lt.octopusEnergies = append(lt.octopusEnergies, v)
	}
	lt.width = len(in)
	lt.length += 1
}

func (lt *DumboOctopuses) Energy(x, y int) int {
	if x < 0 || y < 0 || x >= lt.width || y >= lt.length {
		return math.MaxInt
	}
	return lt.octopusEnergies[x+y*lt.width]
}

func (lt *DumboOctopuses) Length() int {
	return lt.length
}
func (lt *DumboOctopuses) Width() int {
	return lt.width
}
