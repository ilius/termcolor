//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ilius/termcolor"
)

func main() {
	fp, err := os.Create("colors.html")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	colorColWidth := "%50"
	codeColWidth := "%15"
	hexColWidth := "%30"

	// tableWidth := "260em"
	// tableWidth := "%50"
	// tableWidth := "50vw"
	tableWidth := "300px"

	bgColor := "#252525"
	fgColor := "#ffffff"

	fmt.Fprintf(fp, "<html>\n<head><title>8-bit Colors</title></head>\n")
	fmt.Fprintf(fp, "<body style=\"background-color:%s;color:%s;\">\n", bgColor, fgColor)
	fmt.Fprintf(fp, "<table width=\"%s\">\n", tableWidth)
	fmt.Fprintf(fp, "\t<tr>\n\t\t"+strings.Join([]string{
		`<th width="%s" align="center">Color</th>`,
		`<th width="%s" align="center">#</th>`,
		`<th width="%s" align="center">Hex</th>`,
	}, "\n\t\t"), colorColWidth, codeColWidth, hexColWidth)

	for code, c := range termcolor.Colors {
		if code != int(c.Code) {
			panic("bad color code")
		}
		fmt.Fprintf(fp, "\t<tr>\n")
		fmt.Fprintf(fp, "\t\t"+`<td width="%s" style="background-color:%s">&nbsp;</td>`, colorColWidth, c.Hex)
		fmt.Fprintf(fp, "\t\t"+`<td width="%s" align="center">%d</td>`, codeColWidth, code)
		fmt.Fprintf(fp, "\t\t"+`<td width="%s"><pre>%s</pre></td>`, hexColWidth, c.Hex)
	}

	fmt.Fprintf(fp, "\n</table>\n</body>\n</html>")
}
