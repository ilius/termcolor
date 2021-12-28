package termcolor

import (
	"image/color"
	"testing"

	"github.com/ilius/is/v2"
)

func TestRgbToFloat(t *testing.T) {
	is := is.New(t)
	is.Equal(rgbToFloat(&color.RGBA{255, 0, 0, 255}), [3]float64{1, 0, 0})
}
