//go:build ignore
// +build ignore

// reads names.go (var ColorNames) and generates colors.go

package main

import (
	"fmt"
	"io/ioutil"
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
	return fmt.Sprintf(strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.3f", f), "0"), "."))
}

func formatHSL(h, s, l float64) string {
	return fmt.Sprintf(
		"%s, %s, %s",
		formatFloat(h),
		formatFloat(s),
		formatFloat(l),
	)
}

func main() {
	sampleTitle := "Sample"
	sampleCharN := 59

	headerTpl := `## Terminal / ANSI Colors

| %s | Code | Hex       | RGB           | HSL               |
| %s | ---- | --------- | ------------- | ----------------- |
`
	mdText := fmt.Sprintf(
		headerTpl,
		sampleTitle+strings.Repeat(" ", sampleCharN-len(sampleTitle)),
		strings.Repeat("-", sampleCharN),
	)
	for code, _ := range termcolor.ColorNames {
		red, green, blue := termcolor.CodeToRGB(uint8(code))
		cf, ok := colorful.MakeColor(simpleRGB{red, green, blue})
		if !ok {
			panic("failed to make color")
		}
		H, S, L := cf.Hsl()
		hexColor := termcolor.RGBToHexColor(red, green, blue)
		rgb := fmt.Sprintf("%d, %d, %d", red, green, blue)
		hsl := formatHSL(H, S, L)
		mdText += fmt.Sprintf(
			"| %s | %- 4s | `%- 7s` | %- 13s | %- 17s |\n",
			fmt.Sprintf(
				`![](https://via.placeholder.com/60x30/%s/000000?text=+)`,
				hexColor[1:],
			),
			fmt.Sprintf("%v", code),
			hexColor,
			rgb,
			hsl,
		)
	}
	{
		err := ioutil.WriteFile("colors.md", []byte(mdText), 0644)
		if err != nil {
			panic(err)
		}
	}
}
