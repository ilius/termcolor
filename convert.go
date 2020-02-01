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

func validRoundMode(mode RoundMode) bool {
	switch mode {
	case RoundCloser:
		return true
	case RoundDown:
		return true
	case RoundUp:
		return true
	}
	return false
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

	channelSwitch := func(u uint8, ri uint8, gi uint8, bi uint8) *ColorProp {
		switch u {
		case r:
			return Colors[ri]
		case g:
			return Colors[gi]
		case b:
			return Colors[bi]
		}
		return nil
	}

	switch {
	case x <= delta && around(z, 170):
		if y <= delta {
			return channelSwitch(z, 1, 2, 4)
		} else if around(y, 170) {
			return channelSwitch(x, 6, 5, 142)
		} else {
			// {170, 85, 0}
			if around(r, 170) && around(g, 85) && b <= delta {
				return Colors[3]
			}
		}
	case around(x, 85) && around(z, 255):
		if around(y, 85) {
			return channelSwitch(z, 9, 10, 12)
		} else if around(y, 255) {
			return channelSwitch(x, 14, 13, 11)
		}
	}
	return nil
}

// ClosestToRGB finds the closest terminal color to a given full RGB color
func ClosestToRGB(in *ClosestToRGBInput) (*ColorProp, error) {
	if !validRoundMode(in.RoundMode) {
		return nil, fmt.Errorf("invalid RoundMode")
	}
	ri := roundValue(in.Target.R, in.RoundMode)
	gi := roundValue(in.Target.G, in.RoundMode)
	bi := roundValue(in.Target.B, in.RoundMode)
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
		dist := DistanceRGB(&res.RGBA, &in.Target)
		distGray := DistanceRGB(&resGray.RGBA, &in.Target)
		if distGray < dist {
			res = resGray
		}
	}
	resPal := closestFromPalette(&in.Target)
	if resPal != nil {
		dist := DistanceRGB(&res.RGBA, &in.Target)
		distPal := DistanceRGB(&resPal.RGBA, &in.Target)
		if distPal <= dist { // pallete is preferred if same distance
			res = resPal
		}
	}
	return res, nil
}
