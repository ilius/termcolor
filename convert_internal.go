package termcolor

import (
	"sort"
)

var values = [6]uint8{
	0,
	95,  // prev + 95
	135, // prev + 40
	175, // prev + 40
	215, // prev + 40
	255, // prev + 40
}

// these colors are compatible with Xfce terminal and LXDE terminal
// but Tango palette (in Gnome Terminal) has brighter colors
var first16Colors = map[uint8][3]uint8{
	0:  {0, 0, 0},       // 000000, tango:2e3436, group 1
	8:  {85, 85, 85},    // 555555, tango:555753, group 1
	7:  {185, 185, 185}, // b9b9b9, tango:d3d7cf, group 1
	15: {255, 255, 255}, // ffffff, tango:eeeeec, group 1
	1:  {170, 0, 0},     // aa0000, tango:cc0000, group 2
	2:  {0, 170, 0},     // 00aa00, tango:4e9a06, group 2
	4:  {0, 0, 170},     // 0000aa, tango:3465a4, group 2
	3:  {170, 85, 0},    // aa5500, tango:c4a000, group 3
	5:  {170, 0, 170},   // aa00aa, tango:75507b, group 4
	6:  {0, 170, 170},   // 00aaaa, tango:06989a, group 4
	9:  {255, 85, 85},   // ff5555, tango:ef2929, group 5
	10: {85, 255, 85},   // 55ff55, tango:8ae234, group 5
	12: {85, 85, 255},   // 5555ff, tango:729fcf, group 5
	11: {255, 255, 85},  // ffff55, tango:fce94f, group 6
	13: {255, 85, 255},  // ff55ff, tango:ad7fa8, group 6
	14: {85, 255, 255},  // 55ffff, tango:34e2e2, group 6
}

// returns round(a/b) based on given RoundMode
func divideRound(a uint8, b uint8, mode RoundMode) uint8 {
	switch mode {
	case RoundDown:
		return a / b
	case RoundUp:
		return divideRoundUp(a, b)
	}
	return divideRoundCloser(a, b)
}

func diffWithMode(target uint8, source uint8, mode RoundMode) uint8 {
	switch mode {
	case RoundUp:
		if source > target {
			return 255
		}
		return target - source
	case RoundDown:
		if source < target {
			return 255
		}
		return source - target
	}
	// assume it's RoundCloser
	return absInt8Diff(target, source)
}

func closerToWhite(v uint8, mode RoundMode) bool {
	switch mode {
	case RoundCloser:
		return v > 246
	case RoundUp:
		return v > 238
	}
	return false
}

func closestGrayCodeToRGB(vv uint8, mode RoundMode) uint8 {
	grayIndex := divideRound(vv-8, 10, mode)
	closeToValue := func(v1 uint8) bool {
		return diffWithMode(v1, vv, mode) <
			diffWithMode(Colors[grayIndex+232].RGBA.R, vv, mode)
	}
	switch grayIndex {
	case 7, 8: // {78, 78, 78}, {88, 88, 88}
		if closeToValue(85) {
			return 8 // {85, 85, 85}
		}
	case 17, 18: // {178, 178, 178}, {188, 188, 188}
		if closeToValue(185) {
			return 7 // {185, 185, 185}
		}
	case 23, 24, 25: // {238, 238, 238}
		if closerToWhite(vv, mode) {
			return 15 // white
		} else {
			return 255 // lighest gray
		}
	}
	return grayIndex + 232
}

func roundValue(v uint8, mode RoundMode) int {
	switch mode {
	case RoundDown:
		return roundValueDown(v)
	case RoundUp:
		return roundValueUp(v)
	}
	// assume RoundCloser
	return roundValueCloser(v)
}

func roundValueCloser(v uint8) int {
	switch {
	case v <= 47:
		return 0
	case v <= 115:
		return 1
	case v <= 155:
		return 2
	case v <= 195:
		return 3
	case v <= 235:
		return 4
	}
	return 5
}

func roundValueDown(v uint8) int {
	return 5 - sort.Search(6, func(i int) bool {
		return values[5-i] <= v
	})
}

func roundValueUp(v uint8) int {
	return sort.Search(6, func(i int) bool {
		return values[i] >= v
	})
}
