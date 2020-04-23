package termcolor

import (
	"image/color"
	"testing"

	"github.com/ilius/is/v2"
)

func TestRGBAToHexColor(t *testing.T) {
	is := is.New(t)
	is.Equal(RGBAToHexColor(color.RGBA{0, 0, 0, 0}), "#000000")
	is.Equal(RGBAToHexColor(color.RGBA{0, 255, 0, 0}), "#00ff00")
}

func TestParseHexColor(t *testing.T) {
	is := is.New(t)
	{
		c, err := ParseHexColor("#d7005f")
		is.NotErr(err)
		is.Equal(c, &color.RGBA{215, 0, 95, 255})
	}
	{
		c, err := ParseHexColor("#0f0")
		is.NotErr(err)
		is.Equal(c, &color.RGBA{0, 255, 0, 255})
	}
	{
		c, err := ParseHexColor("#00F")
		is.NotErr(err)
		is.Equal(c, &color.RGBA{0, 0, 255, 255})
	}
	{
		c, err := ParseHexColor("#00g")
		is.ErrMsg(err, "invalid hex color")
		is.Nil(c)
	}
	{
		c, err := ParseHexColor("#00000g")
		is.ErrMsg(err, "invalid hex color")
		is.Nil(c)
	}
	{
		c, err := ParseHexColor("0f0")
		is.ErrMsg(err, "invalid hex color")
		is.Nil(c)
	}
	{
		c, err := ParseHexColor("#0f00")
		is.ErrMsg(err, "invalid hex color")
		is.Nil(c)
	}
}
