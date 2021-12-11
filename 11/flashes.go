package main

func SynchronousFlashStep(lt *DumboOctopuses) int {
	flashCount := 0
	cnt := 0
	for cnt != len(lt.octopusEnergies) {
		flashCount++
		cnt = CountFlashes(lt)
	}
	return flashCount
}

func FlashesAfterSteps(lt *DumboOctopuses, steps int) int {
	flashCount := 0
	for s := 0; s < steps; s++ {
		flashCount += CountFlashes(lt)
	}
	return flashCount
}

func CountFlashes(lt *DumboOctopuses) int {
	// fmt.Println("=== Counting ===")
	energies := lt.octopusEnergies
	flashingOctopuses := map[int]bool{}
	for i := 0; i < len(energies); i++ {
		energies[i] += 1
		if energies[i] > 9 {
			flashOctopus(lt, i, flashingOctopuses)
		}
	}

	for i := range flashingOctopuses {
		energies[i] = 0
	}
	return len(flashingOctopuses)
}

func flashOctopus(lt *DumboOctopuses, octoLoc int, flashingOctopuses map[int]bool) {

	if _, ok := flashingOctopuses[octoLoc]; ok {
		return
	}
	flashingOctopuses[octoLoc] = true
	row := octoLoc / lt.Width()
	// fmt.Println("Flashing Neighbors", octoLoc, row)
	for i := -1; i <= 1; i++ {
		neighborRow := row + i
		for j := -1; j <= 1; j++ {
			neighbor := octoLoc + i*lt.Width() + j
			// fmt.Println("neighbor", neighbor, neighborRow)
			if neighbor >= 0 && neighborRow == neighbor/lt.Width() && neighbor < len(lt.octopusEnergies) {
				lt.octopusEnergies[neighbor] += 1
				// fmt.Println("incr", neighbor, lt.octopusEnergies[neighbor])
				if lt.octopusEnergies[neighbor] == 10 {
					// fmt.Println("flashed", flashingOctopuses)
					flashOctopus(lt, neighbor, flashingOctopuses)
				}
			}
		}
	}
}
