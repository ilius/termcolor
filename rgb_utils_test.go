package termcolor

import (
	"github.com/ilius/is"
	"image/color"
	"testing"
)

func TestRgbToFloat(t *testing.T) {
	is := is.New(t)
	is.Equal(rgbToFloat(&color.RGBA{255, 0, 0, 255}), [3]float64{1, 0, 0})
}
