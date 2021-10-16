//go:build ignore
// +build ignore

// reads names.go (var ColorNames) and generates colors.go

package main

import (
	"encoding/json"
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
		"(%s, %s, %s)",
		formatFloat(h),
		formatFloat(s),
		formatFloat(l),
	)
}

func jsonEncode(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func formatStringList(list []string) string {
	formattedList := make([]string, len(list))
	for i, s := range list {
		formattedList[i] = jsonEncode(s)
	}
	return "[" + strings.Join(formattedList, ", ") + "]"
}

func main() {
	pyCode := `#!/usr/bin/env python3

from collections import namedtuple


ColorProp = namedtuple("ColorProp", (
	"code",  # int
	"rgb",  # Tuple[float, float, float]
	"hsl",  # Tuple[Number, Number, Number]
	"hex",  # str
	"names",  # List[str]
))


colors = [
`
	for code, names := range termcolor.ColorNames {
		red, green, blue := termcolor.CodeToRGB(uint8(code))
		cf, ok := colorful.MakeColor(simpleRGB{red, green, blue})
		if !ok {
			panic("failed to make color")
		}
		H, S, L := cf.Hsl()
		htmlColor := termcolor.RGBToHexColor(red, green, blue)
		pyCode += fmt.Sprintf(
			"\tColorProp(\n"+
				"\t\tcode=%d,\n"+
				"\t\trgb=(%d, %d, %d),\n"+
				"\t\thsl="+formatHSL(H, S, L)+",\n"+
				"\t\thex=%#v,\n"+
				"\t\tnames=%s,\n"+
				"\t),\n",
			code,
			red, green, blue,
			htmlColor,
			formatStringList(names),
		)
	}
	pyCode += "]\n"
	{
		err := ioutil.WriteFile("colors.py", []byte(pyCode), 0644)
		if err != nil {
			panic(err)
		}
	}
}
