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
	Num uint8
	HTML string
	Names []string
}

var Colors = [256]*ColorProp{
`)
	for num, names := range termcolor.ColorNames {
		red, green, blue := termcolor.NumberToRGB(num)
		htmlColor := termcolor.RGBToHtml(red, green, blue)
		fmt.Fprintf(
			file,
			"\t&ColorProp{\n"+
				"\t\tNum: %d,\n"+
				"\t\tRGBA: color.RGBA{0x%02x, 0x%02x, 0x%02x, 0xff},\n"+
				"\t\tHTML: %#v,\n"+
				"\t\tNames: %#v,\n"+
				"\t},\n",
			num,
			red, green, blue,
			htmlColor,
			names,
		)
	}
	fmt.Fprintf(file, "}\n")
}
