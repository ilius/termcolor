package termcolor

import (
	"image/color"
	"sort"
)

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

func rgbToFloat(c *color.RGBA) [3]float64 {
	return [3]float64{
		float64(c.R) / 255,
		float64(c.G) / 255,
		float64(c.B) / 255,
	}
}
