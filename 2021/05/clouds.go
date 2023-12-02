package main

type Point struct {
	X, Y int
}

type Cloud struct {
	From, To Point
}

type Clouds []Cloud

func (c Clouds) NumberOfIntersections() int {
	cloudLocations := map[Point]int{}
	countIntersections := 0
	for _, cloud := range c {
		// fmt.Println("iterating over", cloud)
		for _, currentLoc := range cloud.LocationsWithin() {
			// fmt.Println("loc", currentLoc)
			p := cloudLocations[currentLoc]
			cloudLocations[currentLoc] = p + 1
			if p == 1 {
				countIntersections++
			}
		}
	}
	// fmt.Println("cloudLocations", cloudLocations)
	return countIntersections

}

func (c Cloud) LocationsWithin() []Point {
	verticalIncr := c.To.Y - c.From.Y
	verticalIncrAbs := intAbs(verticalIncr)
	if verticalIncr != 0 {
		verticalIncr = verticalIncr / verticalIncrAbs
	}
	horizontalIncr := c.To.X - c.From.X
	horizontalIncrAbs := intAbs(horizontalIncr)
	if horizontalIncr != 0 {
		horizontalIncr = horizontalIncr / horizontalIncrAbs
	}

	iterationLimit := intMax(verticalIncrAbs, horizontalIncrAbs)

	locations := []Point{}
	for i := 0; i <= iterationLimit; i++ {
		locations = append(locations, Point{c.From.X + horizontalIncr*i, c.From.Y + verticalIncr*i})
	}
	return locations
}

func intAbs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
