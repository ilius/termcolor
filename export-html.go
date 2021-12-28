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

	// tableWidth := "%50"
	tableWidth := "600px"
	codeColWidth := "60px"
	bgColor := "#252525"
	fgColor := "#ffffff"

	fmt.Fprintf(fp, "<html>\n<head><title>8-bit Colors</title></head>\n")
	fmt.Fprintf(fp, "<body style=\"background-color:%s;color:%s;\">\n", bgColor, fgColor)
	fmt.Fprintf(fp, "<table width=\"%s\">\n", tableWidth)
	fmt.Fprintf(fp, "\t<tr>\n\t\t"+strings.Join([]string{
		`<th width="%%30">Color</th>`,
		`<th width="%s" align="center">#</th>`,
		`<th width="%%30">HTML</th>`,
	}, "\n\t\t"), codeColWidth)

	for code, c := range termcolor.Colors {
		if code != int(c.Code) {
			panic("bad color code")
		}
		fmt.Fprintf(fp, "\t<tr>\n")
		fmt.Fprintf(fp, "\t\t"+`<td width="%%30" style="background-color:%s">&nbsp;</th>`, c.Hex)
		fmt.Fprintf(fp, "\t\t"+`<td width="%s" align="center">%d</th>`, codeColWidth, code)
		fmt.Fprintf(fp, "\t\t"+`<td width="%%30">%s</th>`, c.Hex)
	}

	fmt.Fprintf(fp, "\n</table>\n</body>\n</html>")
}
