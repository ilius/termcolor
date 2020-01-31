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

func rgbMaxDelta(c *color.RGBA) uint8 {
	return max3uint8(
		absInt8Diff(c.R, c.G),
		absInt8Diff(c.R, c.B),
		absInt8Diff(c.G, c.B),
	)
}

func rgbAverage(c *color.RGBA) uint8 {
	return (c.R + c.G + c.B) / 3
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

func ClosestGrayToRGB(target *color.RGBA, mode RoundMode) *ColorProp {
	v1 := rgbAverage(target)
	if v1 <= 8 {
		return Colors[232]
	}
	return Colors[divideRound(v1-8, 10, mode)+232]
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
		// fmt.Printf("color=%#v, dist=%.1f, distGray=%.1f\n", in.Target, dist, distGray)
		if distGray < dist {
			res = resGray
		}
	}
	return res, nil
}
