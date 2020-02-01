package termcolor

import (
	"fmt"
	"image/color"
	"sort"
)

type RoundMode uint8

const (
	RoundCloser RoundMode = iota
	RoundDown
	RoundUp
)

// DefaultGrayMaxDelta is default value for GrayMaxDelta as option for ClosestToRGB func
var DefaultGrayMaxDelta uint8 = 20

var values = [6]uint8{
	0,
	95,  // prev + 95
	135, // prev + 40
	175, // prev + 40
	215, // prev + 40
	255, // prev + 40
}

// these colors are compatible with Xfce terminal and LXDE terminal
// but Gnome terminal shows brighter colors that are commented
var first16Colors = map[uint8][3]uint8{
	0:  {0, 0, 0},       // 000000, gnome:2e3436
	8:  {85, 85, 85},    // 555555, gnome:555753
	7:  {185, 185, 185}, // b9b9b9, gnome:d3d7cf
	15: {255, 255, 255}, // ffffff, gnome:eeeeec

	1: {170, 0, 0}, // aa0000, gnome:cc0000
	2: {0, 170, 0}, // 00aa00, gnome:4e9a06
	4: {0, 0, 170}, // 0000aa, gnome:3465a4

	3: {170, 85, 0}, // aa5500, gnome:c4a000

	5: {170, 0, 170}, // aa00aa, gnome:75507b
	6: {0, 170, 170}, // 00aaaa, gnome:06989a

	9:  {255, 85, 85}, // ff5555, gnome:ef2929
	10: {85, 255, 85}, // 55ff55, gnome:8ae234
	12: {85, 85, 255}, // 5555ff, gnome:729fcf

	11: {255, 255, 85}, // ffff55, gnome:fce94f
	13: {255, 85, 255}, // ff55ff, gnome:ad7fa8
	14: {85, 255, 255}, // 55ffff, gnome:34e2e2
}

func CodeToRGB(n uint8) (uint8, uint8, uint8) {
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

func roundValue(v uint8, mode RoundMode) (int, error) {
	switch mode {
	case RoundCloser:
		return roundValueCloser(v), nil
	case RoundDown:
		return roundValueDown(v), nil
	case RoundUp:
		return roundValueUp(v), nil
	}
	return 0, fmt.Errorf("roundValue: invalid mode")
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

func sortedRGB(c *color.RGBA) (uint8, uint8, uint8) {
	s := []uint8{c.R, c.G, c.B}
	sort.Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
	return s[0], s[1], s[2]
}

func rgbMaxDelta(c *color.RGBA) uint8 {
	return max3uint8(
		absInt8Diff(c.R, c.G),
		absInt8Diff(c.R, c.B),
		absInt8Diff(c.G, c.B),
	)
}

func rgbAverage(c *color.RGBA) uint8 {
	return uint8((uint16(c.R) + uint16(c.G) + uint16(c.B)) / 3)
}

// ClosestToRGBInput is input struct for ClosestToRGB
type ClosestToRGBInput struct {
	Target       color.RGBA
	RoundMode    RoundMode // optional, default: RoundCloser
	GrayMaxDelta *uint8    // optional, default: DefaultGrayMaxDelta
}

// getGrayMaxDelta returns GrayMaxDelta if non-nil, and default value (DefaultGrayMaxDelta) if nil
func (in *ClosestToRGBInput) getGrayMaxDelta() uint8 {
	if in.GrayMaxDelta != nil {
		return *in.GrayMaxDelta
	}
	return DefaultGrayMaxDelta
}

func (in *ClosestToRGBInput) isCloseToGray() bool {
	return rgbMaxDelta(&in.Target) <= in.getGrayMaxDelta()
}

func divideRoundUp(a uint8, b uint8) uint8 {
	return (a + b - 1) / b
}

func divideRoundCloser(a uint8, b uint8) uint8 {
	if a%b > b/2 {
		return (a + b - 1) / b
	}
	return a / b
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

func ClosestGrayToRGB(target *color.RGBA, mode RoundMode) *ColorProp {
	vv := rgbAverage(target)
	if vv <= 8 {
		return Colors[232]
	}
	// TODO: better detection of near-black based on mode
	if vv == 255 {
		return Colors[15]
	}
	code := closestGrayCodeToRGB(vv, mode)
	return Colors[code]
}

func closestFromPalette(target *color.RGBA) *ColorProp {
	const delta = 20
	around := func(v uint8, t uint8) bool {
		return absInt8Diff(v, t) <= delta
	}
	r, g, b := target.R, target.G, target.B
	x, y, z := sortedRGB(target)
	// x <= y <= z   and   set(x, y, z) == set(r, g, b)
	if x >= 235 { // close to white
		return Colors[15]
	}
	switch {
	case x <= delta && around(z, 170):
		if y <= delta {
			switch z {
			case r: // {170, 0, 0}
				return Colors[1]
			case g: // {0, 170, 0}
				return Colors[2]
			case b: // {0, 0, 170}
				return Colors[4]
			}
		} else if around(y, 170) {
			switch x {
			case r: // {0, 170, 170}
				return Colors[6]
			case g: // {170, 0, 170}
				return Colors[5]
			case b: // {170, 170, 0} does not exist
				break
			}
		} else {
			// {170, 85, 0}
			if around(r, 170) && around(g, 85) && b <= delta {
				return Colors[3]
			}
		}
		break
	case around(x, 85) && around(z, 255):
		if around(y, 85) {
			switch z {
			case r: // {255, 85, 85}
				return Colors[9]
			case g: // {85, 255, 85}
				return Colors[10]
			case b: // {85, 85, 255}
				return Colors[12]
			}
		} else if around(y, 255) {
			switch x {
			case r: // {85, 255, 255}
				return Colors[14]
			case g: // {255, 85, 255}
				return Colors[13]
			case b: // {255, 255, 85}
				return Colors[11]
			}
		}
		break
	}
	return nil
}

// ClosestToRGB finds the closest terminal color to a given full RGB color
func ClosestToRGB(in *ClosestToRGBInput) (*ColorProp, error) {
	ri, err := roundValue(in.Target.R, in.RoundMode)
	if err != nil {
		return nil, err
	}
	gi, err := roundValue(in.Target.G, in.RoundMode)
	if err != nil {
		return nil, err
	}
	bi, err := roundValue(in.Target.B, in.RoundMode)
	if err != nil {
		return nil, err
	}
	code := ri*36 + gi*6 + bi + 16
	if code == 16 {
		code = 0
	}
	/*
		m := n - 16
		ri = int(m/36)
		gi = int((m%36)/6)
		bi = m%6
		red := values[ri]
		green := values[gi]
		blue := values[bi]
	*/
	res := Colors[code]
	if in.isCloseToGray() {
		resGray := ClosestGrayToRGB(&in.Target, in.RoundMode)
		dist := DistanceRGB(res.RGBA, in.Target)
		distGray := DistanceRGB(resGray.RGBA, in.Target)
		if distGray < dist {
			res = resGray
		}
	}
	resPal := closestFromPalette(&in.Target)
	if resPal != nil {
		dist := DistanceRGB(res.RGBA, in.Target)
		distPal := DistanceRGB(resPal.RGBA, in.Target)
		if distPal <= dist { // pallete is preferred if same distance
			res = resPal
		}
	}
	return res, nil
}
