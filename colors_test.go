package termcolor

import (
	"github.com/ilius/is"
	"testing"
)

var testColors = map[int]string{
	0:  "000000",
	1:  "aa0000",
	2:  "00aa00",
	3:  "aa5500",
	4:  "0000aa",
	5:  "aa00aa",
	6:  "00aaaa",
	7:  "b9b9b9",
	8:  "555555",
	9:  "ff5555",
	10: "55ff55",
	11: "ffff55",
	12: "5555ff",
	13: "ff55ff",
	14: "55ffff",
	15: "ffffff",

	16: "000000",
	90: "870087",
	99: "875fff",

	110: "87afd7",
	120: "87ff87",

	232: "080808",
	233: "121212",
	234: "1c1c1c",
	235: "262626",
	236: "303030",
}

func TestColors(t *testing.T) {
	for num, hex := range testColors {
		hex = "#" + hex
		actual := Colors[num]
		if actual.HTML != hex {
			t.Errorf("Num=%d, expected %#v, got %#v", num, hex, actual.HTML)
		}
	}
}

func TestNumberToRGB(t *testing.T) {
	is := is.New(t).Lax()
	for num, hex := range testColors {
		is := is.AddMsg("num=%d", num)
		hex = "#" + hex
		r, g, b := NumberToRGB(num)
		actualHex := RGBToHexColor(r, g, b)
		if !is.Equal(hex, actualHex) {
			c, err := ParseHexColor(hex)
			if err != nil {
				panic(err)
			}
			t.Logf("%s = {%d, %d, %d}", hex, c.R, c.G, c.B)
		}
	}

}
