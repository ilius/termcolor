// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/ilius/termcolor"
)

func main() {
	file, err := os.Create("colors.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprintf(file, `package termcolor

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
`)
	for code, names := range termcolor.ColorNames {
		red, green, blue := termcolor.NumberToRGB(code)
		htmlColor := termcolor.RGBToHexColor(red, green, blue)
		fmt.Fprintf(
			file,
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
	fmt.Fprintf(file, "}\n")
}
