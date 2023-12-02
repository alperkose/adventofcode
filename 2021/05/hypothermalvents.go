package main

import (
	"strconv"
	"strings"
)

type HypoThermalVentMap struct {
	clouds []Cloud
}

func (htvm *HypoThermalVentMap) Accept(in string) {
	points := strings.Split(in, " -> ")
	htvm.clouds = append(htvm.clouds, Cloud{toPoint(points[0]), toPoint(points[1])})
}

func (htvm *HypoThermalVentMap) Clouds() Clouds {
	return htvm.clouds
}

func toPoint(commaSeparatedPoint string) Point {
	point := strings.Split(commaSeparatedPoint, ",")
	x, _ := strconv.Atoi(point[0])
	y, _ := strconv.Atoi(point[1])
	return Point{x, y}
}
