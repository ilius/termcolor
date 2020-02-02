// +build ignore

package main

import (
	"fmt"
	"go/format"
	"io/ioutil"

	"github.com/ilius/termcolor"
)

func main() {
	goCode := `package termcolor

import (
	"image/color"
)

type ColorProp struct {
	RGBA color.RGBA
	Code uint8
	Hex string
	Names []string
}

var Colors = [256]*ColorProp{
`
	for code, names := range termcolor.ColorNames {
		red, green, blue := termcolor.CodeToRGB(uint8(code))
		htmlColor := termcolor.RGBToHexColor(red, green, blue)
		goCode += fmt.Sprintf(
			"\t&ColorProp{\n"+
				"\t\tCode: %d,\n"+
				"\t\tRGBA: color.RGBA{%d, %d, %d, 255},\n"+
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
	err = ioutil.WriteFile("colors.go", goCodeNew, 0644)
	if err != nil {
		panic(err)
	}
}
