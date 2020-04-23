package termcolor

import (
	"testing"

	"github.com/ilius/is/v2"
)

var testColors = map[int]string{
	0:   "000000",
	1:   "aa0000",
	2:   "00aa00",
	3:   "aa5500",
	4:   "0000aa",
	5:   "aa00aa",
	6:   "00aaaa",
	7:   "b9b9b9",
	8:   "555555",
	9:   "ff5555",
	10:  "55ff55",
	11:  "ffff55",
	12:  "5555ff",
	13:  "ff55ff",
	14:  "55ffff",
	15:  "ffffff",
	16:  "000000",
	20:  "0000d7",
	30:  "008787",
	40:  "00d700",
	50:  "00ffd7",
	60:  "5f5f87",
	70:  "5faf00",
	80:  "5fd7d7",
	90:  "870087",
	99:  "875fff",
	100: "878700",
	110: "87afd7",
	120: "87ff87",
	130: "af5f00",
	140: "af87d7",
	150: "afd787",
	160: "d70000",
	170: "d75fd7",
	180: "d7af87",
	190: "d7ff00",
	200: "ff00d7",
	210: "ff8787",
	220: "ffd700",
	230: "ffffd7",
	232: "080808",
	233: "121212",
	234: "1c1c1c",
	235: "262626",
	236: "303030",
	240: "585858",
	245: "8a8a8a",
	250: "bcbcbc",
	255: "eeeeee",
}

func TestColors(t *testing.T) {
	for code, hex := range testColors {
		hex = "#" + hex
		actual := Colors[code]
		if actual.Hex != hex {
			t.Errorf("code=%d, expected %#v, got %#v", code, hex, actual.Hex)
		}
	}
}

func TestCodeToRGB(t *testing.T) {
	is := is.New(t).Lax()
	for code, hex := range testColors {
		is := is.AddMsg("code=%d", code)
		hex = "#" + hex
		r, g, b := CodeToRGB(uint8(code))
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
