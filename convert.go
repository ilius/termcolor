package termcolor

import (
	"fmt"
	"image/color"
)

// DefaultGrayMaxDelta is default value for GrayMaxDelta as option for ClosestToRGB func
var DefaultGrayMaxDelta uint8 = 20

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

// ClosestToRGBInput is input struct for ClosestToRGB
type ClosestToRGBInput struct {
	GrayMaxDelta *uint8 // optional, default: DefaultGrayMaxDelta
	Target       color.RGBA
	RoundMode    RoundMode // optional, default: RoundCloser
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

func ClosestToRGBGray(target *color.RGBA, mode RoundMode) *ColorProp {
	vv := rgbAverage(target)
	if vv <= 8 {
		return Colors[232]
	}
	// TODO: better detection of near-black based on mode
	if vv == 255 {
		return Colors[15]
	}
	code := closestToRGBGrayCode(vv, mode)
	return Colors[code]
}

func ClosestToRGBFromPalette(target *color.RGBA) *ColorProp {
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
	if !in.RoundMode.isValid() {
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
		resGray := ClosestToRGBGray(&in.Target, in.RoundMode)
		dist := DistanceRGB(&res.RGBA, &in.Target)
		distGray := DistanceRGB(&resGray.RGBA, &in.Target)
		if distGray < dist {
			res = resGray
		}
	}
	resPal := ClosestToRGBFromPalette(&in.Target)
	if resPal != nil {
		dist := DistanceRGB(&res.RGBA, &in.Target)
		distPal := DistanceRGB(&resPal.RGBA, &in.Target)
		if distPal <= dist { // pallete is preferred if same distance
			res = resPal
		}
	}
	return res, nil
}
