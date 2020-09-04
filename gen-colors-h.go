// +build ignore
// reads names.go (var ColorNames) and generates colors.h

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

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
		"{%s, %s, %s}",
		formatFloat(h),
		formatFloat(s),
		formatFloat(l),
	)
}

func main() {
	cCode := `#include <stdio.h>
#include <stdint.h>


typedef struct ColorProp {
	uint8_t code;
	uint8_t rgba[4];
	double hsl[3];
	char *hex;
	char *names; // separated by ","
	uint8_t nameCount; 
} ColorProp;


ColorProp colors[] = {
`
	for code, names := range termcolor.ColorNames {
		red, green, blue := termcolor.CodeToRGB(uint8(code))
		cf, ok := colorful.MakeColor(simpleRGB{red, green, blue})
		if !ok {
			panic("failed to make color")
		}
		H, S, L := cf.Hsl()
		htmlColor := termcolor.RGBToHexColor(red, green, blue)
		cCode += fmt.Sprintf(
			"\t{\n"+
				"\t\t.code = %d,\n"+
				"\t\t.rgba = {%d, %d, %d, 255},\n"+
				"\t\t.hsl = "+formatHSL(H, S, L)+",\n"+
				"\t\t.hex = %#v,\n"+
				"\t\t.names = %#v,\n"+
				"\t},\n",
			code,
			red, green, blue,
			htmlColor,
			strings.Join(names, ","),
		)
	}
	cCode += "};"
	{
		err := ioutil.WriteFile("colors.h", []byte(cCode), 0644)
		if err != nil {
			panic(err)
		}
	}
}
