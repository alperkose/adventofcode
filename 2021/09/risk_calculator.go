package main

func RiskFactorOfHoles(lt *LavaTubes) int {
	// fmt.Println(lt)
	risk := 0
	lenghtBoundary := lt.Length()
	widthBoundary := lt.Width()
	for i := 0; i < widthBoundary; i++ {
		for j := 0; j < lenghtBoundary; j++ {
			r, low := compare4Neighbors(lt, i, j)
			if low {
				// println("found a hole", i, j, r)
				risk += r + 1
			}
		}
	}
	return risk
}

func compare4Neighbors(lt *LavaTubes, i, j int) (int, bool) {
	currentDepth := lt.Loc(i, j)
	// fmt.Println(currentDepth, i, j, width, lenght)
	return currentDepth, !(lt.Loc(i-1, j) <= currentDepth ||
		lt.Loc(i+1, j) <= currentDepth ||
		lt.Loc(i, j-1) <= currentDepth ||
		lt.Loc(i, j+1) <= currentDepth)
}
