package termcolor

var values = [6]int{
	0x00,
	0x5f, // prev + 95
	0x87, // prev + 40
	0xaf, // prev + 40
	0xd7, // prev + 40
	0xff, // prev + 40
}

// these colors are compatible with Xfce terminal and LXDE terminal
// but Gnome terminal shows brighter colors that are commented
var first16Colors = map[int][3]int{
	0:  {0, 0, 0},       // 000000, gnome:2e3436
	1:  {170, 0, 0},     // aa0000, gnome:cc0000
	2:  {0, 170, 0},     // 00aa00, gnome:4e9a06
	3:  {170, 85, 0},    // aa5500, gnome:c4a000
	4:  {0, 0, 170},     // 0000aa, gnome:3465a4
	5:  {170, 0, 170},   // aa00aa, gnome:75507b
	6:  {0, 170, 170},   // 00aaaa, gnome:06989a
	7:  {185, 185, 185}, // b9b9b9, gnome:d3d7cf
	8:  {85, 85, 85},    // 555555, gnome:555753
	9:  {255, 85, 85},   // ff5555, gnome:ef2929
	10: {85, 255, 85},   // 55ff55, gnome:8ae234
	11: {255, 255, 85},  // ffff55, gnome:fce94f
	12: {85, 85, 255},   // 5555ff, gnome:729fcf
	13: {255, 85, 255},  // ff55ff, gnome:ad7fa8
	14: {85, 255, 255},  // 55ffff, gnome:34e2e2
	15: {255, 255, 255}, // ffffff, gnome:eeeeec
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
