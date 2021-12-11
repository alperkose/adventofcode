package main

func GammaRate(values []uint32) uint32 {
	gammaRate := uint32(0)
	for mask := uint32(1); mask != 0; mask = mask << 1 {
		gammaRate |= commonBitForMask(values, mask)
	}
	return gammaRate
}

func commonBitForMask(values []uint32, mask uint32) uint32 {
	cnt := 0
	for i := 0; i < len(values); i++ {
		if mask&values[i] == 0 {
			continue
		}
		cnt++
	}
	if cnt*2 >= len(values) {
		return mask
	}
	return 0
}

// func EpsilonRate
