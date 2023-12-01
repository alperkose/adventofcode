package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LowestRiskInTwoByTwo(t *testing.T) {
	cave := &ChitonCave{}
	cave.Accept("15")
	cave.Accept("21")

	assert.Equal(t, 3, LowestRiskPath(cave))
}
func Test_LowestRiskInThreeByThree(t *testing.T) {
	cave := &ChitonCave{}
	cave.Accept("159")
	cave.Accept("213")
	cave.Accept("213")

	assert.Equal(t, 7, LowestRiskPath(cave))
}
func Test_LowestRiskThatGoesOverAllDirections(t *testing.T) {
	cave := &ChitonCave{}
	cave.Accept("119111")
	cave.Accept("911191")
	cave.Accept("999991")
	cave.Accept("999991")
	cave.Accept("999111")
	cave.Accept("999199")
	cave.Accept("999111")

	assert.Equal(t, 17, LowestRiskPath(cave))
}

func Test_LowestRiskInSample(t *testing.T) {
	cave := &ChitonCave{}
	cave.Accept("1163751742")
	cave.Accept("1381373672")
	cave.Accept("2136511328")
	cave.Accept("3694931569")
	cave.Accept("7463417111")
	cave.Accept("1319128137")
	cave.Accept("1359912421")
	cave.Accept("3125421639")
	cave.Accept("1293138521")
	cave.Accept("2311944581")
	assert.Equal(t, 40, LowestRiskPath(cave))
}

func Test_LowestRiskWithExtensionInSample(t *testing.T) {
	cave := &ChitonCave{}
	cave.Accept("1163751742")
	cave.Accept("1381373672")
	cave.Accept("2136511328")
	cave.Accept("3694931569")
	cave.Accept("7463417111")
	cave.Accept("1319128137")
	cave.Accept("1359912421")
	cave.Accept("3125421639")
	cave.Accept("1293138521")
	cave.Accept("2311944581")
	assert.Equal(t, 315, LowestRiskPathOfExtended(cave, 5))
}

func Benchmark_Sample_ExtendByFactorOf1(b *testing.B) {
	cave := &ChitonCave{}
	cave.Accept("1163751742")
	cave.Accept("1381373672")
	cave.Accept("2136511328")
	cave.Accept("3694931569")
	cave.Accept("7463417111")
	cave.Accept("1319128137")
	cave.Accept("1359912421")
	cave.Accept("3125421639")
	cave.Accept("1293138521")
	cave.Accept("2311944581")

	for i := 1; i < b.N; i++ {
		LowestRiskPathOfExtended(cave, 1)
	}
}
func Benchmark_Sample_ExtendByFactorOf3(b *testing.B) {
	cave := &ChitonCave{}
	cave.Accept("1163751742")
	cave.Accept("1381373672")
	cave.Accept("2136511328")
	cave.Accept("3694931569")
	cave.Accept("7463417111")
	cave.Accept("1319128137")
	cave.Accept("1359912421")
	cave.Accept("3125421639")
	cave.Accept("1293138521")
	cave.Accept("2311944581")

	for i := 1; i < b.N; i++ {
		LowestRiskPathOfExtended(cave, 3)
	}
}
func Benchmark_Sample_ExtendByFactorOf5(b *testing.B) {
	cave := &ChitonCave{}
	cave.Accept("1163751742")
	cave.Accept("1381373672")
	cave.Accept("2136511328")
	cave.Accept("3694931569")
	cave.Accept("7463417111")
	cave.Accept("1319128137")
	cave.Accept("1359912421")
	cave.Accept("3125421639")
	cave.Accept("1293138521")
	cave.Accept("2311944581")

	for i := 1; i < b.N; i++ {
		LowestRiskPathOfExtended(cave, 5)
	}
}
