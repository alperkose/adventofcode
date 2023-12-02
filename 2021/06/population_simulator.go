package main

const CYCLE_DURATION = 7
const NEW_SPAWN_CYCLE = 9

func Simulate_loop(lf *LanternFishes, days int) int {
	population := lf.Cycles()
	daysIncrement := days % 7
	for i := daysIncrement; i <= days; i += daysIncrement {
		popSize := len(population)
		// fmt.Println("cycle", i, popSize, population)
		for j := 0; j < popSize; j++ {
			diff := population[j] - daysIncrement
			// fmt.Println("ind  ", j, population[j], diff, daysIncrement)
			if diff < 0 {
				population = append(population, NEW_SPAWN_CYCLE+diff)
				population[j] = CYCLE_DURATION + diff
			} else {
				population[j] = diff
			}
		}
		daysIncrement = CYCLE_DURATION
	}
	// fmt.Println("cycle", days, len(population), population)
	return len(population)
}

func Simulate(lf *LanternFishes, days int) int {

	pg := make([]int, 9)
	for _, c := range lf.Cycles() {
		pg[c]++
	}

	for i := 0; i < days; i++ {
		pg[0], pg[1], pg[2], pg[3], pg[4], pg[5], pg[6], pg[7], pg[8] = pg[1], pg[2], pg[3], pg[4], pg[5], pg[6], pg[7]+pg[0], pg[8], pg[0]
	}

	return pg[0] + pg[1] + pg[2] + pg[3] + pg[4] + pg[5] + pg[6] + pg[7] + pg[8]
}
