package main

type Entry struct {
	signals []string
	digits  []string
}

func (e Entry) DigitValues() int {
	byteToDigit := make(ByteSlice, 10)

	loopCount := 0
	for i := 0; !isDigitized(byteToDigit) && loopCount < 100; i = (i + 1) % 10 {
		signal := e.signals[i]
		signalAsBye := asByte(signal)
		// fmt.Println("digits", byteToDigit, signal, signalAsBye)
		switch len(signal) {
		case 2:
			byteToDigit[1] = signalAsBye
		case 3:
			byteToDigit[7] = signalAsBye
		case 4:
			byteToDigit[4] = signalAsBye
		case 5:
			if byteToDigit[3] == 0 && byteToDigit[1] > 0 && byteToDigit[1]&signalAsBye == byteToDigit[1] {
				byteToDigit[3] = signalAsBye
			} else if byteToDigit[5] == 0 && byteToDigit[6] > 0 && byteToDigit[6]&signalAsBye == signalAsBye {
				byteToDigit[5] = signalAsBye
			} else if byteToDigit[5] > 0 && byteToDigit[5] != signalAsBye && byteToDigit[3] > 0 && byteToDigit[3] != signalAsBye {
				byteToDigit[2] = signalAsBye
			}
		case 6:
			if byteToDigit[6] == 0 && byteToDigit[1] > 0 && byteToDigit[1]&signalAsBye != byteToDigit[1] {
				byteToDigit[6] = signalAsBye
			} else if byteToDigit[6] == 0 && byteToDigit[7] > 0 && byteToDigit[7]&signalAsBye != byteToDigit[7] {
				byteToDigit[6] = signalAsBye
			} else if byteToDigit[9] == 0 && byteToDigit[4] > 0 && byteToDigit[4]&signalAsBye == byteToDigit[4] {
				byteToDigit[9] = signalAsBye
			} else if byteToDigit[9] > 0 && byteToDigit[9] != signalAsBye && byteToDigit[6] > 0 && byteToDigit[6] != signalAsBye {
				byteToDigit[0] = signalAsBye
			}
		case 7:
			byteToDigit[8] = signalAsBye
		}
		loopCount++
	}
	return byteToDigit.indexOf(asByte(e.digits[0]))*1000 +
		byteToDigit.indexOf(asByte(e.digits[1]))*100 +
		byteToDigit.indexOf(asByte(e.digits[2]))*10 +
		byteToDigit.indexOf(asByte(e.digits[3]))
}

func isDigitized(digits []byte) bool {
	for _, d := range digits {
		if d == 0 {
			return false
		}
	}
	return true
}

func asByte(signal string) byte {
	b := byte(0)
	for _, c := range signal {
		mask := charToByte[c]
		b |= mask
	}
	return b
}

var charToByte = map[rune]byte{
	'a': 1 << 7,
	'b': 1 << 6,
	'c': 1 << 5,
	'd': 1 << 4,
	'e': 1 << 3,
	'f': 1 << 2,
	'g': 1,
}

type ByteSlice []byte

func (byteSlice ByteSlice) indexOf(term byte) int {
	for i, v := range byteSlice {
		if v == term {
			return i
		}
	}
	return -1
}
