package termcolor

func divideRoundUp(a uint8, b uint8) uint8 {
	return (a + b - 1) / b
}

func divideRoundCloser(a uint8, b uint8) uint8 {
	if a%b > b/2 {
		return (a + b - 1) / b
	}
	return a / b
}

func absInt8Diff(a uint8, b uint8) uint8 {
	d := int16(a) - int16(b)
	if d < 0 {
		d = -d
	}
	return uint8(d)
}

func max3uint8(a uint8, b uint8, c uint8) uint8 {
	if a < b {
		a = b
	}
	// ^ in Python words: a = max(a, b)
	if a < c {
		a = c
	}
	// ^ in Python words: a = max(a, c)
	return a
}
