package main

type pt1 map[string]int

func SolverForPart1() Solver {
	return pt1{"red": 12, "green": 13, "blue": 14}
}

func (conf pt1) Solve(session GameSession) int {
	runningSum := 0
	failed := false
	for _, game := range session {
		failed = false
		for _, rvl := range game.reveals {
			if conf["red"] < rvl["red"] || conf["green"] < rvl["green"] || conf["blue"] < rvl["blue"] {
				failed = true
				break
			}
		}

		if !failed {
			runningSum += game.id
		}
	}
	return runningSum
}

type pt2 map[string]int

func SolverForPart2() Solver {
	return pt2{}
}

func (conf pt2) Solve(session GameSession) int {
	runningSum := 0
	for _, game := range session {
		minRed := 0
		minGreen := 0
		minBlue := 0
		for _, rvl := range game.reveals {
			c, ok := rvl["red"]
			if ok && c > minRed {
				minRed = c
			}
			c, ok = rvl["green"]
			if ok && c > minGreen {
				minGreen = c
			}
			c, ok = rvl["blue"]
			if ok && c > minBlue {
				minBlue = c
			}
		}
		runningSum += minBlue * minGreen * minRed
	}
	return runningSum
}
