package main

func ChunkScore(chunk string) (int, int) {
	fifo := []rune{}
	for _, r := range chunk {
		expectedOpener := openerOf(r)
		// fmt.Printf("current %s %s %v\n", string(r), string(expectedOpener), fifo)
		if expectedOpener == '.' {
			fifo = append(fifo, r)
			continue
		}
		lastOpen := fifo[len(fifo)-1]
		// fmt.Println("lastopen", string(lastOpen), lastOpen)
		if lastOpen != expectedOpener {
			return scoreOf(expectedOpener), 0
		}
		// fmt.Println("closing ", string(lastOpen), lastOpen)
		fifo = fifo[:len(fifo)-1]
	}
	return 0, completetionScore(fifo)
}

func completetionScore(fifo []rune) int {
	sum := 0
	// fmt.Println("completetionScore", fifo)
	for i := len(fifo) - 1; i >= 0; i-- {
		// fmt.Println("compl", string(fifo[i]))
		sum = sum*5 + completionScoreOf(fifo[i])
	}
	return sum
}

func completionScoreOf(closer rune) int {

	switch closer {
	case '(':
		return 1
	case '[':
		return 2
	case '{':
		return 3
	case '<':
		return 4
	}
	return 0
}

func scoreOf(closer rune) int {
	switch closer {
	case '(':
		return 3
	case '[':
		return 57
	case '{':
		return 1197
	case '<':
		return 25137
	}
	return 0

}

func openerOf(open rune) rune {
	switch open {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	case '>':
		return '<'
	}
	return '.'
}
