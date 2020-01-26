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

	fmt.Fprintf(fp, "<html>\n<head><title>8-bit Colors</title></head>\n<body>\n")
	fmt.Fprintf(fp, "<table width=\"%s\">\n", tableWidth)
	fmt.Fprintf(fp, "\t<tr>\n\t\t"+strings.Join([]string{
		`<th width="%%30">Color</th>`,
		`<th width="%%10">Code</th>`,
		`<th width="%%30">HTML</th>`,
	}, "\n\t\t"))

	for num, c := range termcolor.Colors {
		if num != int(c.Num) {
			panic("bad number")
		}
		fmt.Fprintf(fp, "\t<tr>\n")
		fmt.Fprintf(fp, "\t\t"+`<td width="%%30" style="background-color:%s">&nbsp;</th>`, c.HTML)
		fmt.Fprintf(fp, "\t\t"+`<td width="%%10">%d</th>`, num)
		fmt.Fprintf(fp, "\t\t"+`<td width="%%30">%s</th>`, c.HTML)
	}

	fmt.Fprintf(fp, "\n</table>\n</body>\n</html>")

}
