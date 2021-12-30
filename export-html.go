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

func exportToHTML(filename string, colors []*termcolor.ColorProp) {
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	colorColWidth := "%30"
	codeColWidth := "%10"
	hexColWidth := "%20"

	satColWidth := "%10"
	hueColWidth := "%10"
	lightColWidth := "%10"
	valueColWidth := "%10"

	tableWidth := "500px"

	bgColor := "#252525"
	fgColor := "#ffffff"

	fmt.Fprintf(fp, "<html>\n<head><title>8-bit Colors</title></head>\n")
	fmt.Fprintf(fp, "<body style=\"background-color:%s;color:%s;\">\n", bgColor, fgColor)
	fmt.Fprintf(fp, "<table width=\"%s\">\n", tableWidth)
	fmt.Fprintf(fp, "\t<tr>\n\t\t"+strings.Join([]string{
		`<th width="%s" align="center">Color</th>`,
		`<th width="%s" align="center">Code</th>`,
		`<th width="%s" align="center">Hex</th>`,
		`<th width="%s" align="center">Sat</th>`,
		`<th width="%s" align="center">Hue</th>`,
		`<th width="%s" align="center">Light</th>`,
		`<th width="%s" align="center">Value</th>`,
	}, "\n\t\t"), colorColWidth, codeColWidth, hexColWidth, hueColWidth, satColWidth, lightColWidth, valueColWidth)

	for _, c := range colors {
		code := c.Code
		fmt.Fprintf(fp, "\t<tr>\n")
		fmt.Fprintf(fp, "\t\t"+`<td width="%s" style="background-color:%s">&nbsp;</td>`, colorColWidth, c.Hex)
		fmt.Fprintf(fp, "\t\t"+`<td width="%s" align="center">%d</td>`, codeColWidth, code)
		fmt.Fprintf(fp, "\t\t"+`<td width="%s" align="center"><pre>%s</pre></td>`, hexColWidth, c.Hex)
		fmt.Fprintf(fp, "\t\t"+`<td width="%s" align="center">%s</td>`, satColWidth, formatFloat(c.HSL[1]))
		fmt.Fprintf(fp, "\t\t"+`<td width="%s" align="center">%05.1f</td>`, hueColWidth, c.HSL[0])
		fmt.Fprintf(fp, "\t\t"+`<td width="%s" align="center">%.4f</td>`, lightColWidth, c.HSL[2])
		fmt.Fprintf(fp, "\t\t"+`<td width="%s" align="center">%.4f</td>`, lightColWidth, c.HSV[2])
	}

	fmt.Fprintf(fp, "\n</table>\n</body>\n</html>")
}

func main() {
	colors := termcolor.Colors[:]

	exportToHTML("colors.html", colors)

	// sort by Saturation (with delta <= 0.1 allowed) in reverse order
	// then by Hue (with delta <= 5 allowed)
	// then by Lightness in reverse order
	sort.Slice(colors, func(i, j int) bool {
		i_hsl := colors[i].HSL
		j_hsl := colors[j].HSL

		if math.Abs(i_hsl[1]-j_hsl[1]) > 0.1 {
			return i_hsl[1] > j_hsl[1]
		}

		if math.Abs(i_hsl[0]-j_hsl[0]) > 5 {
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
	exportToHTML("colors-by-saturation.html", colors)
}
