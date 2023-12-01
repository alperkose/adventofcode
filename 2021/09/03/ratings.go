package main

func O2GenerationRating(values []uint32, digit uint32) uint32 {
	o2contestants := make([]uint32, len(values))
	copy(o2contestants, values)
	for mask := uint32(1) << (digit - 1); mask != 0; mask >>= 1 {
		o2contestants = filterOnCurrentBit(o2contestants, commonBitForMask(o2contestants, mask), mask)
	}
	return o2contestants[0]
}

func CO2ScrubberRating(values []uint32, digit uint32) uint32 {
	co2contestants := make([]uint32, len(values))
	copy(co2contestants, values)
	for mask := uint32(1) << (digit - 1); mask != 0; mask >>= 1 {
		co2contestants = filterOnCurrentBit(co2contestants, ^commonBitForMask(co2contestants, mask)&mask, mask)
	}
	return co2contestants[0]
}

func filterOnCurrentBit(contestants []uint32, currentBit, mask uint32) []uint32 {
	survivors := []uint32{}
	if len(contestants) > 1 {
		for i := 0; i < len(contestants); i++ {
			// fmt.Printf("contestant %12b - %12b -> %12b\n", contestants[i], currentBit, ^(currentBit^contestants[i])&mask)
			if ^(currentBit^contestants[i])&mask >= 1 {
				survivors = append(survivors, contestants[i])
			}
		}
		// fmt.Printf("survivors %v\n", survivors)
		return survivors
	}
	return contestants
}
