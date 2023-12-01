package main

func ExtendCave(cave []int, width, factor int) ([]int, int) {
	factorSq := factor * factor
	newCave := make([]int, len(cave)*factorSq)
	newWidth := width * factor

	for i := 0; i < len(cave); i++ {
		row := i / width
		ind := i%width + row*newWidth
		for f := 0; f < factor; f++ {
			newCave[ind+f*width] = (cave[i]+f)%10 + (cave[i]+f)/10
		}
		for f := 1; f < factor; f++ {
			dind := ind + f*factor*len(cave)
			for ff := 0; ff < factor; ff++ {
				newCave[dind+ff*width] = (newCave[ind+ff*width]+f)%10 + (newCave[ind+ff*width]+f)/10
			}
		}
		// fmt.Println(row, i, ind, newCave)
	}
	return newCave, newWidth
}
