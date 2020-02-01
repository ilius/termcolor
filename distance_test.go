package termcolor

import (
	"image/color"
	"testing"

	"github.com/ilius/is"
)

func TestDistanceRGB(t *testing.T) {
	test := func(c1 color.RGBA, c2 color.RGBA, dist float64) {
		is := is.New(t).AddMsg("c1=%v, c2=%v", c1, c2)
		actual := DistanceRGB(&c1, &c2)
		isFloatBetween(is, actual, dist, dist+0.1)
	}
	test(
		color.RGBA{0, 0, 0, 0},
		color.RGBA{0, 0, 0, 0},
		0,
	)
	test(
		color.RGBA{0, 0, 0, 0},
		color.RGBA{255, 255, 255, 0},
		441.6,
	)
	test(
		color.RGBA{255, 0, 0, 0},
		color.RGBA{0, 255, 0, 0},
		360.6,
	)
	test(
		color.RGBA{0, 100, 0, 0},
		color.RGBA{0, 0, 100, 0},
		141.4,
	)
}
