package main

import (
	"fmt"
	"sort"
)

func ThreeLargestBasinsMultiplied(lt *LavaTubes) int {
	basins := []int{}
	lenghtBoundary := lt.Length()
	widthBoundary := lt.Width()
	basinFilteredMap := make([]int, lenghtBoundary*widthBoundary)

	for i := 0; i < widthBoundary; i++ {
		for j := 0; j < lenghtBoundary; j++ {
			if lt.Loc(i, j) == 9 {
				basinFilteredMap[i+j*widthBoundary] = -1
			}
			if basinFilteredMap[i+j*widthBoundary] == 0 {
				basinSize := markBasin(lt, basinFilteredMap, i, j, len(basins)+1)
				basins = append(basins, basinSize)
			}

		}
	}
	// fmt.Println("bs map:", basinFilteredMap)
	fmt.Println("basins:", basins)
	product := 1
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	fmt.Println("basins:", basins)
	for i := 0; i < len(basins) && i < 3; i++ {
		product *= basins[i]
	}
	return product
}

func markBasin(lt *LavaTubes, basinMap []int, x, y, basinID int) int {
	pos := x + y*lt.Width()
	// fmt.Println(lt.heightMap)
	// fmt.Println(basinMap)
	// fmt.Println(x, y, pos, basinID)

	stack := []int{pos}

	basinSize := 0
	for stackIndex := 0; stackIndex < len(stack); stackIndex++ {
		pos := stack[stackIndex]
		if basinMap[pos] != 0 {
			continue
		}
		row := pos / lt.Width()
		basinSize++
		// fmt.Println("marking", pos)
		basinMap[pos] = basinID

		posToChk := pos - lt.Width()
		if posToChk > 0 && basinMap[posToChk] == 0 && lt.heightMap[posToChk] != 9 {
			stack = append(stack, posToChk)
		}
		posToChk = pos + lt.Width()
		if posToChk < len(basinMap) && basinMap[posToChk] == 0 && lt.heightMap[posToChk] != 9 {
			stack = append(stack, posToChk)
		}
		posToChk = pos - 1
		if posToChk > 0 && posToChk/lt.Width() == row && basinMap[posToChk] == 0 && lt.heightMap[posToChk] != 9 {
			stack = append(stack, posToChk)
		}
		posToChk = pos + 1
		if posToChk < len(basinMap) && posToChk/lt.Width() == row && basinMap[posToChk] == 0 && lt.heightMap[posToChk] != 9 {
			stack = append(stack, posToChk)
		}
	}
	return basinSize
}
