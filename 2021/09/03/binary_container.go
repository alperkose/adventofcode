package main

type ByteContainer struct {
	digits int
	values []uint32
}

func (bc *ByteContainer) Accept(in string) {
	val := uint32(0)
	mask := uint32(1)
	for i := len(in) - 1; i >= 0; i-- {
		if in[i] == '1' {
			val |= mask
		}
		mask = mask << 1
	}
	bc.values = append(bc.values, val)
	if len(in) > bc.digits {
		bc.digits = len(in)
	}
}

func (bc *ByteContainer) Values() []uint32 {
	return bc.values
}

func (bc *ByteContainer) Digits() int {
	return bc.digits
}
