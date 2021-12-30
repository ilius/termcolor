//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/ilius/termcolor"
)

func formatFloat(f float64) string {
	return fmt.Sprintf(strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.3f", f), "0"), "."))
}

func exportTableToHTML(filename string, colorTable [][]*termcolor.ColorProp) {
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	bgColor := "#252525"
	fgColor := "#ffffff"

	tableWidth := "600px"
	cellWidth := "40px"

	fmt.Fprintf(fp, "<html>\n<head><title>8-bit Colors (Compact)</title></head>\n")
	fmt.Fprintf(fp, "<body style=\"background-color:%s;color:%s;\">\n", bgColor, fgColor)
	fmt.Fprintf(fp, "<table width=\"%s\">\n", tableWidth)
	for _, row := range colorTable {
		fmt.Fprintf(fp, "\t<tr>\n")
		for _, c := range row {
			code := c.Code
			var textColor string
			if c.HSL[2] < 0.5 {
				textColor = "#ffffff"
			} else {
				textColor = "#000000"
			}
			fmt.Fprintf(
				fp,
				"\t\t"+`<td width="%s" align="center" style="background-color:%s;color:%s">%03d</td>`+"\n",
				cellWidth,
				c.Hex,
				textColor,
				code,
			)
		}
		fmt.Fprintf(fp, "\t</tr>\n")
	}

	fmt.Fprintf(fp, "</table>\n</body>\n</html>")
}

func main() {
	colors := termcolor.Colors[:]

	maxDeltaSat := 0.1
	maxDeltaHue := 5.0

	// sort by Saturation (with delta <= 0.1 allowed) in reverse order
	// then by Hue (with delta <= 5 allowed)
	// then by Lightness in reverse order
	sort.Slice(colors, func(i, j int) bool {
		i_hsl := colors[i].HSL
		j_hsl := colors[j].HSL

		if math.Abs(i_hsl[1]-j_hsl[1]) > maxDeltaSat {
			return i_hsl[1] > j_hsl[1]
		}

		if math.Abs(i_hsl[0]-j_hsl[0]) > maxDeltaHue {
			return i_hsl[0] < j_hsl[0]
		}

		if i_hsl[2] != j_hsl[2] {
			return i_hsl[2] > j_hsl[2]
		}

		// if colors[i].HSV[2] != colors[j].HSV[2] {
		//	return colors[i].HSV[2] > colors[j].HSV[2]
		// }

		return colors[i].Code < colors[j].Code
	})

	table := [][]*termcolor.ColorProp{}

	newRow := func(c *termcolor.ColorProp) {
		table = append(table, []*termcolor.ColorProp{c})
	}

	for i, c := range colors {
		if i == 0 {
			newRow(c)
			continue
		}
		if math.Abs(c.HSL[1]-colors[i-1].HSL[1]) > maxDeltaSat {
			newRow(c)
			continue
		}
		if math.Abs(c.HSL[0]-colors[i-1].HSL[0]) > maxDeltaHue {
			newRow(c)
			continue
		}
		n := len(table)
		table[n-1] = append(table[n-1], c)
	}

	exportTableToHTML("colors-compact.html", table)
}
