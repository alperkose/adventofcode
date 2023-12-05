package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/alperkose/adventofcode/2023/input"
)

func main() {
	stuff := &SeedsAndMappedStuff{}
	input.Parse(os.Stdin, stuff)
	fmt.Println("Pt1: ", stuff.SolvePt1())
	fmt.Println("Pt2: ", stuff.SolvePt2())
}

const (
	SeedToFertilizerMap   = 0
	SeedToSoil            = 1
	SoilToFertilizer      = 2
	FertilizerToWater     = 3
	WaterToLight          = 4
	LightToTemperature    = 5
	TemperatureToHumidity = 6
	HumidityToLocation    = 7
)

type SeedsAndMappedStuff struct {
	seeds       []int
	maps        []StuffMap
	updatingMap int
}

func (stuff *SeedsAndMappedStuff) Accept(inp string) {
	if strings.HasPrefix(inp, "seeds: ") {
		seedStrs := strings.Split(strings.TrimSpace(inp[7:]), " ")
		for _, ss := range seedStrs {
			seed, _ := strconv.Atoi(ss)
			stuff.seeds = append(stuff.seeds, seed)
		}
		return
	}
	if len(strings.TrimSpace(inp)) == 0 {
		return
	}
	if strings.Index(inp, "map:") > 0 {
		stuff.maps = append(stuff.maps, StuffMap{})
	} else {
		mapStrs := strings.Split(strings.TrimSpace(inp), " ")
		dest, _ := strconv.Atoi(mapStrs[0])
		source, _ := strconv.Atoi(mapStrs[1])
		rng, _ := strconv.Atoi(mapStrs[2])
		stuff.maps[len(stuff.maps)-1] = append(stuff.maps[len(stuff.maps)-1], StuffMapEntry{dest, source, rng})
	}
}

func (stuff *SeedsAndMappedStuff) SolvePt1() int {
	// fmt.Printf("%+v\n", stuff)
	minLoc := math.MaxInt
	for _, seed := range stuff.seeds {
		tempVal := seed
		for _, smap := range stuff.maps {
			tempVal = smap.Get(tempVal)
			// fmt.Printf("smap(%4d) -> %v\n", ind, tempVal)
		}
		if tempVal < minLoc {
			minLoc = tempVal
		}
		// fmt.Printf("seed(%4d) -> %v\n", seed, tempVal)
	}
	return minLoc
}

var wg sync.WaitGroup

func (stuff *SeedsAndMappedStuff) SolvePt2() int {
	c := make(chan int)

	for i := 0; i < len(stuff.seeds); i += 2 {

		wg.Add(1)
		go func(c chan int, wg *sync.WaitGroup, index int) {
			defer wg.Done() // 3
			minLoc := math.MaxInt
			for j := 0; j < stuff.seeds[index+1]; j++ {
				seed := stuff.seeds[index] + j
				tempVal := seed
				for _, smap := range stuff.maps {
					tempVal = smap.Get(tempVal)
					// fmt.Printf("smap(%4d) -> %v\n", ind, tempVal)

				}
				if tempVal < minLoc {
					minLoc = tempVal
				}
				// fmt.Printf("seed(%4d) -> %v\n", seed, tempVal)
			}
			c <- minLoc
		}(c, &wg, i)
	}
	minLoc := math.MaxInt
	for i := range c {
		fmt.Println(i)
		if i < minLoc {
			minLoc = i
		}
	}
	fmt.Printf("minimum location: %d", minLoc)
	close(c)
	return minLoc
}
