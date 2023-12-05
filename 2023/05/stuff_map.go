package main

type StuffMapEntry struct {
	destination int
	source      int
	rng         int
}

func (entry StuffMapEntry) Get(src int) int {
	diff := src - entry.source
	if diff < 0 || diff >= entry.rng {
		return -1
	}
	return entry.destination + diff
}

type StuffMap []StuffMapEntry

func (mp StuffMap) Get(src int) int {
	for _, m := range mp {
		val := m.Get(src)
		if val >= 0 {
			return val
		}
	}
	return src
}
