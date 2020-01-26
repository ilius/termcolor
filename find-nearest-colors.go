// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"sort"

	. "github.com/ilius/termcolor"
)

func parseCssColorNamesFile(fpath string) map[string]*color.RGBA {
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
			panic(err)
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
		"distance: %.1f, name: %s, color: %s",
		x.Distance,
		x.Name,
		RGBAToHexColor(*x.Color),
	)
}

func main() {
	cssColors := parseCssColorNamesFile("css-color-names.json")
	data := map[string][]string{}
	for _, c := range Colors {
		items := []*DistItem{}
		for name, cc := range cssColors {
			items = append(items, &DistItem{
				Color:    cc,
				Name:     name,
				Distance: DistanceRGB(c.RGBA, *cc),
			})
		}
		sort.Slice(items, func(i int, j int) bool {
			return items[i].Distance < items[j].Distance
		})
		if items[0].Distance >= 30 {
			continue
		}
		end := 1
		for ;end<4; end++ {
			if items[end].Distance > 0 {
				break
			}
		}
		items = items[:end]
		strItems := make([]string, len(items))
		for i, item := range items {
			strItems[i] = item.String()
		}
		data[fmt.Sprintf("dist=%04.1f - num=%d - %s", items[0].Distance, c.Num, c.Hex)] = strItems
	}
	jsonBytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("nearest-colors.json", jsonBytes, 0644)
	if err != nil {
		panic(err)
	}
}
