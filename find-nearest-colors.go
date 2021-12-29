//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"sort"

	. "github.com/ilius/termcolor"
)

var skipColorCodes = map[uint8]uint8{
	246: 1, // 'mountain mist' (#959396) is close but not gray
	247: 1, // 'star dust' (#9f9f9c) is close but not gray
	241: 1, // 'storm dust' (#646463') is close but not gray
	188: 1, // 'iron' (#d4d7d9) is close but not gray
	102: 1, // 'suva gray' (#888387) is close but not gray
	245: 1, // 'stack' (#8a8f8a) is close but not gray
	249: 1, // 'nobel' (#b7b1b1) is close but not gray

	0:   1,
	15:  1,
	16:  1,
	21:  1,
	23:  1,
	24:  1,
	46:  1,
	51:  1,
	65:  1,
	73:  1,
	93:  1,
	101: 1,
	124: 1,
	194: 1,
	196: 1,
	201: 1,
	210: 1,
	217: 1,
	220: 1,
	221: 1,
	223: 1,
	224: 1,
	226: 1,
	229: 1,
	230: 1,
	231: 1,
	236: 1,
	242: 1,
	244: 1,
	248: 1,
	253: 1,
	254: 1,
	255: 1,
}

func parseColorNamesJsonFile(fpath string) map[string]*color.RGBA {
	origMap := map[string]string{}
	jsonB, err := ioutil.ReadFile(fpath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonB, &origMap)
	if err != nil {
		panic(err)
	}
	res := map[string]*color.RGBA{}
	for name, hex := range origMap {
		color, err := ParseHexColor(hex)
		if err != nil {
			panic(fmt.Errorf("invalid hex color %#v: %v", hex, err))
		}
		res[name] = color
	}
	return res
}

type DistItem struct {
	Color    *color.RGBA
	Name     string
	Distance float64
}

func (x *DistItem) String() string {
	return fmt.Sprintf(
		"distance: %.1f, '%s' (%s)",
		x.Distance,
		x.Name,
		RGBAToHexColor(*x.Color),
	)
}

func main() {
	inputFileName := os.Args[1]
	outputFileName := "nearest-colors.json"
	colorNames := parseColorNamesJsonFile(inputFileName)
	data := map[string][]string{}
	for _, c := range Colors {
		if skipColorCodes[c.Code] > 0 {
			continue
		}

		currentNames := map[string]bool{}
		for _, name := range c.Names {
			currentNames[name] = true
		}
		if len(currentNames) >= 2 {
			continue
		}

		items := []*DistItem{}
		for name, cc := range colorNames {
			distance := DistanceRGB(&c.RGBA, cc)
			if distance > 10 {
				continue
			}
			items = append(items, &DistItem{
				Color:    cc,
				Name:     name,
				Distance: distance,
			})
		}
		if len(items) == 0 {
			continue
		}
		sort.Slice(items, func(i int, j int) bool {
			return items[i].Distance < items[j].Distance
		})
		if currentNames[items[0].Name] {
			continue
		}
		minDistance := items[0].Distance
		end := 3
		// end := 1
		// for ; end < 4; end++ {
		//	if items[end].Distance > 0 {
		//		break
		//	}
		// }
		if len(items) > end {
			items = items[:end]
		}
		strItems := make([]string, len(items))
		for i, item := range items {
			strItems[i] = item.String()
		}
		data[fmt.Sprintf(
			"dist=%04.1f - currentNames=%d - code=%.3d - %s",
			minDistance,
			len(currentNames),
			c.Code,
			c.Hex,
		)] = strItems
	}
	jsonBytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(outputFileName, jsonBytes, 0644)
	if err != nil {
		panic(err)
	}
}
