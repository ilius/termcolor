//go:build ignore
// +build ignore

// reads colors.go (var Colors) and generates colors-hex.go

package main

import (
	"fmt"
	"go/format"
	"io/ioutil"

	"github.com/ilius/termcolor"
)

func main() {
	goCode := `package termcolor

var ColorsHex = [256]string{
`
	for _, color := range termcolor.Colors {
		goCode += fmt.Sprintf(
			"\t%#v,\n",
			color.Hex,
		)
	}
	goCode += "}"
	goCodeNew, err := format.Source([]byte(goCode))
	if err != nil {
		panic(err)
	}
	{
		err := ioutil.WriteFile("colors-hex.go", goCodeNew, 0644)
		if err != nil {
			panic(err)
		}
	}
}
