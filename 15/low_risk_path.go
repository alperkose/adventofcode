package main

import (
	"container/heap"
	"math"
)

type Path struct {
	lastStep int
	value    int
}

func LowestRiskPath(cave *ChitonCave) int {
	return lowestRiskPath(cave.Risks(), cave.Width())
}

func LowestRiskPathOfExtended(cave *ChitonCave, factor int) int {
	return lowestRiskPath(ExtendCave(cave.Risks(), cave.Width(), factor))
}

func lowestRiskPath(cave []int, width int) int {
	paths := PriorityQueue{{0, cave[0]}}
	heap.Init(&paths)
	riskOfThatStep := make([]int, len(cave))
	for i := 0; i < len(riskOfThatStep); i++ {
		riskOfThatStep[i] = math.MaxInt
	}
	// pathToCheck := []Path{{0, cave[0]}}

	// for pathIndex := 0; pathIndex < len(pathToCheck); pathIndex++ {
	for len(paths) > 0 {
		currentPath := heap.Pop(&paths).(Path)
		// currentPath := pathToCheck[pathIndex]
		if currentPath.value > riskOfThatStep[len(riskOfThatStep)-1] {
			continue
		}

		posToChk := currentPath.lastStep - width
		if posToChk > 0 && (riskOfThatStep[posToChk] > (cave[posToChk] + currentPath.value)) {
			riskOfThatStep[posToChk] = currentPath.value + cave[posToChk]
			// pathToCheck = append(pathToCheck, Path{posToChk, currentPath.value + cave[posToChk]})
			heap.Push(&paths, Path{posToChk, currentPath.value + cave[posToChk]})
		}
		posToChk = currentPath.lastStep + width
		if posToChk < len(riskOfThatStep) && (riskOfThatStep[posToChk] > (cave[posToChk] + currentPath.value)) {
			riskOfThatStep[posToChk] = currentPath.value + cave[posToChk]
			// pathToCheck = append(pathToCheck, Path{posToChk, currentPath.value + cave[posToChk]})
			heap.Push(&paths, Path{posToChk, currentPath.value + cave[posToChk]})
		}
		posToChk = currentPath.lastStep - 1
		row := currentPath.lastStep / width
		if posToChk > 0 && posToChk/width == row && (riskOfThatStep[posToChk] > (cave[posToChk] + currentPath.value)) {
			riskOfThatStep[posToChk] = currentPath.value + cave[posToChk]
			// pathToCheck = append(pathToCheck, Path{posToChk, currentPath.value + cave[posToChk]})
			heap.Push(&paths, Path{posToChk, currentPath.value + cave[posToChk]})
		}
		posToChk = currentPath.lastStep + 1
		if posToChk < len(riskOfThatStep) && posToChk/width == row && (riskOfThatStep[posToChk] > (cave[posToChk] + currentPath.value)) {
			riskOfThatStep[posToChk] = currentPath.value + cave[posToChk]
			// pathToCheck = append(pathToCheck, Path{posToChk, currentPath.value + cave[posToChk]})
			heap.Push(&paths, Path{posToChk, currentPath.value + cave[posToChk]})
		}
	}
	return riskOfThatStep[len(riskOfThatStep)-1] - cave[0]
}
