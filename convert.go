package termcolor

var values = [6]int{
	0x00,
	0x5f, // prev + 95
	0x87, // prev + 40
	0xaf, // prev + 40
	0xd7, // prev + 40
	0xff, // prev + 40
}

var first16Colors = map[int][3]int{
	0:  {0, 0, 0},       // #000000
	1:  {204, 0, 0},     // #cc0000
	2:  {78, 154, 6},    // #4e9a06
	3:  {196, 160, 0},   // #c4a000
	4:  {52, 101, 164},  // #3465a4
	5:  {117, 80, 123},  // #75507b
	6:  {6, 152, 154},   // #06989a
	7:  {211, 215, 207}, // #d3d7cf
	8:  {85, 87, 83},    // #555753
	9:  {239, 41, 41},   // #ef2929
	10: {138, 226, 52},  // #8ae234
	11: {252, 233, 79},  // #fce94f
	12: {114, 159, 207}, // #729fcf
	13: {173, 127, 168}, // #ad7fa8
	14: {52, 226, 226},  // #34e2e2
	15: {238, 238, 236}, // #eeeeec
}

func NumberToRGB(n int) (int, int, int) {
	if n < 16 {
		c := first16Colors[n]
		return c[0], c[1], c[2]
	}
	if n >= 232 {
		v := 8 + 10*(n-232)
		return v, v, v
	}
	m := n - 16

	red := values[int(m/36)]
	green := values[int((m%36)/6)]
	blue := values[m%6]

	return red, green, blue
}
