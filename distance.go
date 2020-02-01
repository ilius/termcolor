package termcolor

import (
	"image/color"
	"math"
)

func chanDiffSq(x1 uint8, x2 uint8) int32 {
	d := int32(x1) - int32(x2)
	return d * d
}

func DistanceRGB(c1 *color.RGBA, c2 *color.RGBA) float64 {
	return math.Sqrt(float64(
		chanDiffSq(c1.R, c2.R) +
			chanDiffSq(c1.G, c2.G) +
			chanDiffSq(c1.B, c2.B),
	))
}
