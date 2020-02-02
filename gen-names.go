// +build ignore

package main

import (
	"fmt"
	"go/format"
	"io/ioutil"

	"github.com/ilius/termcolor"
)

func main() {
	goCode := `//go:generate go run gen-names.go

package termcolor

var ColorNames = [][]string{
`
	for _, c := range termcolor.Colors {
		goCode += fmt.Sprintf("\t%#v,\n", c.Names)
	}
	goCode += "}"
	goCodeNew, err := format.Source([]byte(goCode))
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("names.go", goCodeNew, 0644)
	if err != nil {
		panic(err)
	}
}
