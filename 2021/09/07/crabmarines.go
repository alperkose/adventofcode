package main

import (
	"strconv"
	"strings"
)

type CrabMarines struct {
	placement []int
}

func (cm *CrabMarines) Accept(in string) {
	placementStr := strings.Split(in, ",")
	for _, pstr := range placementStr {
		v, _ := strconv.Atoi(pstr)
		cm.placement = append(cm.placement, v)
	}
}

func (cm *CrabMarines) Placements() []int {
	return cm.placement
}
