// +build ignore

package main

import (
	"fmt"
	"go/format"
	"io/ioutil"

	"github.com/ilius/termcolor"
)

var ignoreColors = map[int]bool{
	16:  true, // duplicate of 0 = black
	231: true, // duplicate of 15 = white
}

func main() {
	goCode := `//go:generate go run gen-byname.go
	
package termcolor

var ColorLookup = map[string]uint8{
`
	for code, c := range termcolor.Colors {
		if ignoreColors[code] {
			continue
		}
		for _, name := range c.Names {
			goCode += fmt.Sprintf("\t%#v: %d,\n", name, code)
		}
		goCode += fmt.Sprintf("\t%#v: %d,\n", c.Hex, code)
	}
	goCode += "}"
	goCodeNew, err := format.Source([]byte(goCode))
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("lookup.go", goCodeNew, 0644)
	if err != nil {
		panic(err)
	}
}
