// +build ignore

package main

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"strconv"

	"github.com/ilius/termcolor"
	"github.com/lucasb-eyer/go-colorful"
)

func uint8to32(x uint8) uint32 {
	return uint32(float64(x) * 257.0)
}

type simpleRGB [3]uint8

func (c simpleRGB) RGBA() (r, g, b, a uint32) {
	return uint8to32(c[0]), uint8to32(c[1]), uint8to32(c[2]), 0xffff
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func formatHSL(h, s, l float64) string {
	return fmt.Sprintf(
		"[3]float64{%s, %s, %s}",
		formatFloat(h),
		formatFloat(s),
		formatFloat(l),
	)
}

func main() {
	goCode := `package termcolor

import (
	"image/color"
)

type ColorProp struct {
	Code uint8
	RGBA color.RGBA
	HSL [3]float64
	Hex string
	Names []string
}

var Colors = [256]*ColorProp{
`
	for code, names := range termcolor.ColorNames {
		red, green, blue := termcolor.CodeToRGB(uint8(code))
		cf := colorful.MakeColor(simpleRGB{red, green, blue})
		H, S, L := cf.Hsl()
		htmlColor := termcolor.RGBToHexColor(red, green, blue)
		goCode += fmt.Sprintf(
			"\t&ColorProp{\n"+
				"\t\tCode: %d,\n"+
				"\t\tRGBA: color.RGBA{%d, %d, %d, 255},\n"+
				"\t\tHSL: "+formatHSL(H, S, L)+",\n"+
				"\t\tHex: %#v,\n"+
				"\t\tNames: %#v,\n"+
				"\t},\n",
			code,
			red, green, blue,
			htmlColor,
			names,
		)
	}
	goCode += "}"
	goCodeNew, err := format.Source([]byte(goCode))
	if err != nil {
		panic(err)
	}
	{
		err := ioutil.WriteFile("colors.go", goCodeNew, 0644)
		if err != nil {
			panic(err)
		}
	}
}
