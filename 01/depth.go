package main

func CountDepthIncrease(depth []int) int {
	cnt := 0
	for i := 1; i < len(depth); i++ {
		if depth[i] > depth[i-1] {
			cnt++
		}
	}
	return cnt
}

func SlidingWindowOfDepthIncreases(depth []int) int {
	if len(depth) < 4 {
		return 0
	}
	cnt := 0
	prev := depth[0] + depth[1] + depth[2]
	for i := 3; i < len(depth); i++ {
		curr := depth[i] + depth[i-1] + depth[i-2]
		if curr > prev {
			cnt++
		}
		prev = curr
	}
	return cnt
}
