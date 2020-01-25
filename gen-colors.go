// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/ilius/termcolor"
)

var values = [6]int{
	0x00,
	0x5f, // prev + 95
	0x87, // prev + 40
	0xaf, // prev + 40
	0xd7, // prev + 40
	0xff, // prev + 40
}

var colors = map[int][3]int{
	0:  {0, 0, 0},       // #000000
	1:  {204, 0, 0},     // #cc0000
	2:  {78, 154, 6},    // #4e9a06
	3:  {196, 160, 0},   // #c4a000
	4:  {52, 101, 164},  // #3465a4
	5:  {117, 80, 123},  // #75507b
	6:  {6, 152, 154},   // #06989a
	7:  {211, 215, 207}, // #d3d7cf
	8:  {85, 87, 83},    // #555753
	9:  {239, 41, 41},   // #ef2929
	10: {138, 226, 52},  // #8ae234
	11: {252, 233, 79},  // #fce94f
	12: {114, 159, 207}, // #729fcf
	13: {173, 127, 168}, // #ad7fa8
	14: {52, 226, 226},  // #34e2e2
	15: {238, 238, 236}, // #eeeeec
}

func getRGB(n int) (int, int, int) {
	if n < 16 {
		c := colors[n]
		return c[0], c[1], c[2]
	}
	if n >= 232 {
		v := 8 + 10*(n-232)
		return v, v, v
	}
	m := n - 16

	red := values[int(m/36)]
	green := values[int((m%36)/6)]
	blue := values[m%6]

	return red, green, blue
}

func rgbToHtml(r int, g int, b int) string {
	return fmt.Sprintf("#%.2x%.2x%.2x", r, g, b)
}

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
		red, green, blue := getRGB(num)
		htmlColor := rgbToHtml(red, green, blue)
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
