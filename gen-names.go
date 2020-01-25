// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/ilius/termcolor"
)

func main() {
	file, err := os.Create("names.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprintf(file, "package termcolor\n\nvar ColorNames = [][]string{\n")
	for _, c := range termcolor.Colors {
		fmt.Fprintf(file, "\t%#v,\n", c.Names)
	}
	fmt.Fprintf(file, "}\n")
}
