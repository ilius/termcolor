package termcolor

import (
	"image/color"
)

type ColorProp struct {
	RGBA color.RGBA
	Name string
}

var Colors = [256]*ColorProp{
	{color.RGBA{0x00, 0x00, 0x00, 0xff}, "black"},          // 0, #000000
	{color.RGBA{0xcc, 0x00, 0x00, 0xff}, "dark red 1"},     // 1, #cc0000
	{color.RGBA{0x4e, 0x9a, 0x06, 0xff}, "mixed green 1"},  // 2, #4e9a06
	{color.RGBA{0xc4, 0xa0, 0x00, 0xff}, "mixed orange 1"}, // 3, #c4a000
	{color.RGBA{0x34, 0x65, 0xa4, 0xff}, "mixed blue 1"},   // 4, #3465a4
	{color.RGBA{0x75, 0x50, 0x7b, 0xff}, "purple 1"},       // 5, #75507b
	{color.RGBA{0x06, 0x98, 0x9a, 0xff}, "cyan 1"},         // 6, #06989a
	{color.RGBA{0xd3, 0xd7, 0xcf, 0xff}, "light gray"},     // 7, #d3d7cf
	{color.RGBA{0x55, 0x57, 0x53, 0xff}, "dark gray"},      // 8, #555753
	{color.RGBA{0xef, 0x29, 0x29, 0xff}, "light red"},      // 9, #ef2929
	{color.RGBA{0x8a, 0xe2, 0x34, 0xff}, "light green"},    // 10, #8ae234
	{color.RGBA{0xfc, 0xe9, 0x4f, 0xff}, "yellow"},         // 11, #fce94f
	{color.RGBA{0x72, 0x9f, 0xcf, 0xff}, "light blue"},     // 12, #729fcf
	{color.RGBA{0xad, 0x7f, 0xa8, 0xff}, "light purple"},   // 13, #ad7fa8
	{color.RGBA{0x34, 0xe2, 0xe2, 0xff}, "light cyan"},     // 14, #34e2e2
	{color.RGBA{0xee, 0xee, 0xec, 0xff}, "light gray 2"},   // 15, #eeeeec
	{color.RGBA{0x00, 0x00, 0x00, 0xff}, "black"},          // 16, #000000
	{color.RGBA{0x00, 0x00, 0x5f, 0xff}, "blue 4"},         // 17, #00005f
	{color.RGBA{0x00, 0x00, 0x87, 0xff}, "blue 3"},         // 18, #000087
	{color.RGBA{0x00, 0x00, 0xaf, 0xff}, "blue 2"},         // 19, #0000af
	{color.RGBA{0x00, 0x00, 0xd7, 0xff}, "blue 1"},         // 20, #0000d7
	{color.RGBA{0x00, 0x00, 0xff, 0xff}, "blue"},           // 21, #0000ff
	{color.RGBA{0x00, 0x5f, 0x00, 0xff}, "green 4"},        // 22, #005f00
	{color.RGBA{0x00, 0x5f, 0x5f, 0xff}, "blue stone"},     // 23, #005f5f
	{color.RGBA{0x00, 0x5f, 0x87, 0xff}, ""},               // 24, #005f87
	{color.RGBA{0x00, 0x5f, 0xaf, 0xff}, ""},               // 25, #005faf
	{color.RGBA{0x00, 0x5f, 0xd7, 0xff}, ""},               // 26, #005fd7
	{color.RGBA{0x00, 0x5f, 0xff, 0xff}, ""},               // 27, #005fff
	{color.RGBA{0x00, 0x87, 0x00, 0xff}, "green 3"},        // 28, #008700
	{color.RGBA{0x00, 0x87, 0x5f, 0xff}, ""},               // 29, #00875f
	{color.RGBA{0x00, 0x87, 0x87, 0xff}, ""},               // 30, #008787
	{color.RGBA{0x00, 0x87, 0xaf, 0xff}, ""},               // 31, #0087af
	{color.RGBA{0x00, 0x87, 0xd7, 0xff}, ""},               // 32, #0087d7
	{color.RGBA{0x00, 0x87, 0xff, 0xff}, ""},               // 33, #0087ff
	{color.RGBA{0x00, 0xaf, 0x00, 0xff}, "green 2"},        // 34, #00af00
	{color.RGBA{0x00, 0xaf, 0x5f, 0xff}, ""},               // 35, #00af5f
	{color.RGBA{0x00, 0xaf, 0x87, 0xff}, ""},               // 36, #00af87
	{color.RGBA{0x00, 0xaf, 0xaf, 0xff}, ""},               // 37, #00afaf
	{color.RGBA{0x00, 0xaf, 0xd7, 0xff}, ""},               // 38, #00afd7
	{color.RGBA{0x00, 0xaf, 0xff, 0xff}, ""},               // 39, #00afff
	{color.RGBA{0x00, 0xd7, 0x00, 0xff}, "green 1"},        // 40, #00d700
	{color.RGBA{0x00, 0xd7, 0x5f, 0xff}, ""},               // 41, #00d75f
	{color.RGBA{0x00, 0xd7, 0x87, 0xff}, ""},               // 42, #00d787
	{color.RGBA{0x00, 0xd7, 0xaf, 0xff}, ""},               // 43, #00d7af
	{color.RGBA{0x00, 0xd7, 0xd7, 0xff}, ""},               // 44, #00d7d7
	{color.RGBA{0x00, 0xd7, 0xff, 0xff}, ""},               // 45, #00d7ff
	{color.RGBA{0x00, 0xff, 0x00, 0xff}, "green"},          // 46, #00ff00
	{color.RGBA{0x00, 0xff, 0x5f, 0xff}, ""},               // 47, #00ff5f
	{color.RGBA{0x00, 0xff, 0x87, 0xff}, ""},               // 48, #00ff87
	{color.RGBA{0x00, 0xff, 0xaf, 0xff}, ""},               // 49, #00ffaf
	{color.RGBA{0x00, 0xff, 0xd7, 0xff}, ""},               // 50, #00ffd7
	{color.RGBA{0x00, 0xff, 0xff, 0xff}, ""},               // 51, #00ffff
	{color.RGBA{0x5f, 0x00, 0x00, 0xff}, "red 5"},          // 52, #5f0000
	{color.RGBA{0x5f, 0x00, 0x5f, 0xff}, ""},               // 53, #5f005f
	{color.RGBA{0x5f, 0x00, 0x87, 0xff}, ""},               // 54, #5f0087
	{color.RGBA{0x5f, 0x00, 0xaf, 0xff}, ""},               // 55, #5f00af
	{color.RGBA{0x5f, 0x00, 0xd7, 0xff}, ""},               // 56, #5f00d7
	{color.RGBA{0x5f, 0x00, 0xff, 0xff}, ""},               // 57, #5f00ff
	{color.RGBA{0x5f, 0x5f, 0x00, 0xff}, ""},               // 58, #5f5f00
	{color.RGBA{0x5f, 0x5f, 0x5f, 0xff}, ""},               // 59, #5f5f5f
	{color.RGBA{0x5f, 0x5f, 0x87, 0xff}, ""},               // 60, #5f5f87
	{color.RGBA{0x5f, 0x5f, 0xaf, 0xff}, ""},               // 61, #5f5faf
	{color.RGBA{0x5f, 0x5f, 0xd7, 0xff}, ""},               // 62, #5f5fd7
	{color.RGBA{0x5f, 0x5f, 0xff, 0xff}, ""},               // 63, #5f5fff
	{color.RGBA{0x5f, 0x87, 0x00, 0xff}, ""},               // 64, #5f8700
	{color.RGBA{0x5f, 0x87, 0x5f, 0xff}, ""},               // 65, #5f875f
	{color.RGBA{0x5f, 0x87, 0x87, 0xff}, ""},               // 66, #5f8787
	{color.RGBA{0x5f, 0x87, 0xaf, 0xff}, ""},               // 67, #5f87af
	{color.RGBA{0x5f, 0x87, 0xd7, 0xff}, ""},               // 68, #5f87d7
	{color.RGBA{0x5f, 0x87, 0xff, 0xff}, ""},               // 69, #5f87ff
	{color.RGBA{0x5f, 0xaf, 0x00, 0xff}, ""},               // 70, #5faf00
	{color.RGBA{0x5f, 0xaf, 0x5f, 0xff}, ""},               // 71, #5faf5f
	{color.RGBA{0x5f, 0xaf, 0x87, 0xff}, ""},               // 72, #5faf87
	{color.RGBA{0x5f, 0xaf, 0xaf, 0xff}, ""},               // 73, #5fafaf
	{color.RGBA{0x5f, 0xaf, 0xd7, 0xff}, ""},               // 74, #5fafd7
	{color.RGBA{0x5f, 0xaf, 0xff, 0xff}, ""},               // 75, #5fafff
	{color.RGBA{0x5f, 0xd7, 0x00, 0xff}, ""},               // 76, #5fd700
	{color.RGBA{0x5f, 0xd7, 0x5f, 0xff}, ""},               // 77, #5fd75f
	{color.RGBA{0x5f, 0xd7, 0x87, 0xff}, ""},               // 78, #5fd787
	{color.RGBA{0x5f, 0xd7, 0xaf, 0xff}, ""},               // 79, #5fd7af
	{color.RGBA{0x5f, 0xd7, 0xd7, 0xff}, ""},               // 80, #5fd7d7
	{color.RGBA{0x5f, 0xd7, 0xff, 0xff}, ""},               // 81, #5fd7ff
	{color.RGBA{0x5f, 0xff, 0x00, 0xff}, ""},               // 82, #5fff00
	{color.RGBA{0x5f, 0xff, 0x5f, 0xff}, ""},               // 83, #5fff5f
	{color.RGBA{0x5f, 0xff, 0x87, 0xff}, ""},               // 84, #5fff87
	{color.RGBA{0x5f, 0xff, 0xaf, 0xff}, ""},               // 85, #5fffaf
	{color.RGBA{0x5f, 0xff, 0xd7, 0xff}, ""},               // 86, #5fffd7
	{color.RGBA{0x5f, 0xff, 0xff, 0xff}, ""},               // 87, #5fffff
	{color.RGBA{0x87, 0x00, 0x00, 0xff}, "red 4"},          // 88, #870000
	{color.RGBA{0x87, 0x00, 0x5f, 0xff}, ""},               // 89, #87005f
	{color.RGBA{0x87, 0x00, 0x87, 0xff}, ""},               // 90, #870087
	{color.RGBA{0x87, 0x00, 0xaf, 0xff}, ""},               // 91, #8700af
	{color.RGBA{0x87, 0x00, 0xd7, 0xff}, ""},               // 92, #8700d7
	{color.RGBA{0x87, 0x00, 0xff, 0xff}, ""},               // 93, #8700ff
	{color.RGBA{0x87, 0x5f, 0x00, 0xff}, ""},               // 94, #875f00
	{color.RGBA{0x87, 0x5f, 0x5f, 0xff}, ""},               // 95, #875f5f
	{color.RGBA{0x87, 0x5f, 0x87, 0xff}, ""},               // 96, #875f87
	{color.RGBA{0x87, 0x5f, 0xaf, 0xff}, ""},               // 97, #875faf
	{color.RGBA{0x87, 0x5f, 0xd7, 0xff}, ""},               // 98, #875fd7
	{color.RGBA{0x87, 0x5f, 0xff, 0xff}, ""},               // 99, #875fff
	{color.RGBA{0x87, 0x87, 0x00, 0xff}, ""},               // 100, #878700
	{color.RGBA{0x87, 0x87, 0x5f, 0xff}, ""},               // 101, #87875f
	{color.RGBA{0x87, 0x87, 0x87, 0xff}, ""},               // 102, #878787
	{color.RGBA{0x87, 0x87, 0xaf, 0xff}, ""},               // 103, #8787af
	{color.RGBA{0x87, 0x87, 0xd7, 0xff}, ""},               // 104, #8787d7
	{color.RGBA{0x87, 0x87, 0xff, 0xff}, ""},               // 105, #8787ff
	{color.RGBA{0x87, 0xaf, 0x00, 0xff}, ""},               // 106, #87af00
	{color.RGBA{0x87, 0xaf, 0x5f, 0xff}, ""},               // 107, #87af5f
	{color.RGBA{0x87, 0xaf, 0x87, 0xff}, ""},               // 108, #87af87
	{color.RGBA{0x87, 0xaf, 0xaf, 0xff}, ""},               // 109, #87afaf
	{color.RGBA{0x87, 0xaf, 0xd7, 0xff}, ""},               // 110, #87afd7
	{color.RGBA{0x87, 0xaf, 0xff, 0xff}, ""},               // 111, #87afff
	{color.RGBA{0x87, 0xd7, 0x00, 0xff}, ""},               // 112, #87d700
	{color.RGBA{0x87, 0xd7, 0x5f, 0xff}, ""},               // 113, #87d75f
	{color.RGBA{0x87, 0xd7, 0x87, 0xff}, ""},               // 114, #87d787
	{color.RGBA{0x87, 0xd7, 0xaf, 0xff}, ""},               // 115, #87d7af
	{color.RGBA{0x87, 0xd7, 0xd7, 0xff}, ""},               // 116, #87d7d7
	{color.RGBA{0x87, 0xd7, 0xff, 0xff}, ""},               // 117, #87d7ff
	{color.RGBA{0x87, 0xff, 0x00, 0xff}, ""},               // 118, #87ff00
	{color.RGBA{0x87, 0xff, 0x5f, 0xff}, ""},               // 119, #87ff5f
	{color.RGBA{0x87, 0xff, 0x87, 0xff}, ""},               // 120, #87ff87
	{color.RGBA{0x87, 0xff, 0xaf, 0xff}, ""},               // 121, #87ffaf
	{color.RGBA{0x87, 0xff, 0xd7, 0xff}, ""},               // 122, #87ffd7
	{color.RGBA{0x87, 0xff, 0xff, 0xff}, ""},               // 123, #87ffff
	{color.RGBA{0xaf, 0x00, 0x00, 0xff}, "red 3"},          // 124, #af0000
	{color.RGBA{0xaf, 0x00, 0x5f, 0xff}, ""},               // 125, #af005f
	{color.RGBA{0xaf, 0x00, 0x87, 0xff}, ""},               // 126, #af0087
	{color.RGBA{0xaf, 0x00, 0xaf, 0xff}, ""},               // 127, #af00af
	{color.RGBA{0xaf, 0x00, 0xd7, 0xff}, ""},               // 128, #af00d7
	{color.RGBA{0xaf, 0x00, 0xff, 0xff}, ""},               // 129, #af00ff
	{color.RGBA{0xaf, 0x5f, 0x00, 0xff}, ""},               // 130, #af5f00
	{color.RGBA{0xaf, 0x5f, 0x5f, 0xff}, ""},               // 131, #af5f5f
	{color.RGBA{0xaf, 0x5f, 0x87, 0xff}, ""},               // 132, #af5f87
	{color.RGBA{0xaf, 0x5f, 0xaf, 0xff}, ""},               // 133, #af5faf
	{color.RGBA{0xaf, 0x5f, 0xd7, 0xff}, ""},               // 134, #af5fd7
	{color.RGBA{0xaf, 0x5f, 0xff, 0xff}, ""},               // 135, #af5fff
	{color.RGBA{0xaf, 0x87, 0x00, 0xff}, ""},               // 136, #af8700
	{color.RGBA{0xaf, 0x87, 0x5f, 0xff}, ""},               // 137, #af875f
	{color.RGBA{0xaf, 0x87, 0x87, 0xff}, ""},               // 138, #af8787
	{color.RGBA{0xaf, 0x87, 0xaf, 0xff}, ""},               // 139, #af87af
	{color.RGBA{0xaf, 0x87, 0xd7, 0xff}, ""},               // 140, #af87d7
	{color.RGBA{0xaf, 0x87, 0xff, 0xff}, ""},               // 141, #af87ff
	{color.RGBA{0xaf, 0xaf, 0x00, 0xff}, ""},               // 142, #afaf00
	{color.RGBA{0xaf, 0xaf, 0x5f, 0xff}, ""},               // 143, #afaf5f
	{color.RGBA{0xaf, 0xaf, 0x87, 0xff}, ""},               // 144, #afaf87
	{color.RGBA{0xaf, 0xaf, 0xaf, 0xff}, ""},               // 145, #afafaf
	{color.RGBA{0xaf, 0xaf, 0xd7, 0xff}, ""},               // 146, #afafd7
	{color.RGBA{0xaf, 0xaf, 0xff, 0xff}, ""},               // 147, #afafff
	{color.RGBA{0xaf, 0xd7, 0x00, 0xff}, ""},               // 148, #afd700
	{color.RGBA{0xaf, 0xd7, 0x5f, 0xff}, ""},               // 149, #afd75f
	{color.RGBA{0xaf, 0xd7, 0x87, 0xff}, ""},               // 150, #afd787
	{color.RGBA{0xaf, 0xd7, 0xaf, 0xff}, ""},               // 151, #afd7af
	{color.RGBA{0xaf, 0xd7, 0xd7, 0xff}, ""},               // 152, #afd7d7
	{color.RGBA{0xaf, 0xd7, 0xff, 0xff}, ""},               // 153, #afd7ff
	{color.RGBA{0xaf, 0xff, 0x00, 0xff}, ""},               // 154, #afff00
	{color.RGBA{0xaf, 0xff, 0x5f, 0xff}, ""},               // 155, #afff5f
	{color.RGBA{0xaf, 0xff, 0x87, 0xff}, ""},               // 156, #afff87
	{color.RGBA{0xaf, 0xff, 0xaf, 0xff}, ""},               // 157, #afffaf
	{color.RGBA{0xaf, 0xff, 0xd7, 0xff}, ""},               // 158, #afffd7
	{color.RGBA{0xaf, 0xff, 0xff, 0xff}, ""},               // 159, #afffff
	{color.RGBA{0xd7, 0x00, 0x00, 0xff}, "red 2"},          // 160, #d70000
	{color.RGBA{0xd7, 0x00, 0x5f, 0xff}, ""},               // 161, #d7005f
	{color.RGBA{0xd7, 0x00, 0x87, 0xff}, ""},               // 162, #d70087
	{color.RGBA{0xd7, 0x00, 0xaf, 0xff}, ""},               // 163, #d700af
	{color.RGBA{0xd7, 0x00, 0xd7, 0xff}, ""},               // 164, #d700d7
	{color.RGBA{0xd7, 0x00, 0xff, 0xff}, ""},               // 165, #d700ff
	{color.RGBA{0xd7, 0x5f, 0x00, 0xff}, ""},               // 166, #d75f00
	{color.RGBA{0xd7, 0x5f, 0x5f, 0xff}, ""},               // 167, #d75f5f
	{color.RGBA{0xd7, 0x5f, 0x87, 0xff}, ""},               // 168, #d75f87
	{color.RGBA{0xd7, 0x5f, 0xaf, 0xff}, ""},               // 169, #d75faf
	{color.RGBA{0xd7, 0x5f, 0xd7, 0xff}, ""},               // 170, #d75fd7
	{color.RGBA{0xd7, 0x5f, 0xff, 0xff}, ""},               // 171, #d75fff
	{color.RGBA{0xd7, 0x87, 0x00, 0xff}, ""},               // 172, #d78700
	{color.RGBA{0xd7, 0x87, 0x5f, 0xff}, ""},               // 173, #d7875f
	{color.RGBA{0xd7, 0x87, 0x87, 0xff}, ""},               // 174, #d78787
	{color.RGBA{0xd7, 0x87, 0xaf, 0xff}, ""},               // 175, #d787af
	{color.RGBA{0xd7, 0x87, 0xd7, 0xff}, ""},               // 176, #d787d7
	{color.RGBA{0xd7, 0x87, 0xff, 0xff}, ""},               // 177, #d787ff
	{color.RGBA{0xd7, 0xaf, 0x00, 0xff}, ""},               // 178, #d7af00
	{color.RGBA{0xd7, 0xaf, 0x5f, 0xff}, ""},               // 179, #d7af5f
	{color.RGBA{0xd7, 0xaf, 0x87, 0xff}, ""},               // 180, #d7af87
	{color.RGBA{0xd7, 0xaf, 0xaf, 0xff}, ""},               // 181, #d7afaf
	{color.RGBA{0xd7, 0xaf, 0xd7, 0xff}, ""},               // 182, #d7afd7
	{color.RGBA{0xd7, 0xaf, 0xff, 0xff}, ""},               // 183, #d7afff
	{color.RGBA{0xd7, 0xd7, 0x00, 0xff}, ""},               // 184, #d7d700
	{color.RGBA{0xd7, 0xd7, 0x5f, 0xff}, ""},               // 185, #d7d75f
	{color.RGBA{0xd7, 0xd7, 0x87, 0xff}, ""},               // 186, #d7d787
	{color.RGBA{0xd7, 0xd7, 0xaf, 0xff}, ""},               // 187, #d7d7af
	{color.RGBA{0xd7, 0xd7, 0xd7, 0xff}, ""},               // 188, #d7d7d7
	{color.RGBA{0xd7, 0xd7, 0xff, 0xff}, ""},               // 189, #d7d7ff
	{color.RGBA{0xd7, 0xff, 0x00, 0xff}, ""},               // 190, #d7ff00
	{color.RGBA{0xd7, 0xff, 0x5f, 0xff}, ""},               // 191, #d7ff5f
	{color.RGBA{0xd7, 0xff, 0x87, 0xff}, ""},               // 192, #d7ff87
	{color.RGBA{0xd7, 0xff, 0xaf, 0xff}, ""},               // 193, #d7ffaf
	{color.RGBA{0xd7, 0xff, 0xd7, 0xff}, ""},               // 194, #d7ffd7
	{color.RGBA{0xd7, 0xff, 0xff, 0xff}, ""},               // 195, #d7ffff
	{color.RGBA{0xff, 0x00, 0x00, 0xff}, "red 1"},          // 196, #ff0000
	{color.RGBA{0xff, 0x00, 0x5f, 0xff}, ""},               // 197, #ff005f
	{color.RGBA{0xff, 0x00, 0x87, 0xff}, ""},               // 198, #ff0087
	{color.RGBA{0xff, 0x00, 0xaf, 0xff}, ""},               // 199, #ff00af
	{color.RGBA{0xff, 0x00, 0xd7, 0xff}, ""},               // 200, #ff00d7
	{color.RGBA{0xff, 0x00, 0xff, 0xff}, ""},               // 201, #ff00ff
	{color.RGBA{0xff, 0x5f, 0x00, 0xff}, ""},               // 202, #ff5f00
	{color.RGBA{0xff, 0x5f, 0x5f, 0xff}, ""},               // 203, #ff5f5f
	{color.RGBA{0xff, 0x5f, 0x87, 0xff}, ""},               // 204, #ff5f87
	{color.RGBA{0xff, 0x5f, 0xaf, 0xff}, ""},               // 205, #ff5faf
	{color.RGBA{0xff, 0x5f, 0xd7, 0xff}, ""},               // 206, #ff5fd7
	{color.RGBA{0xff, 0x5f, 0xff, 0xff}, ""},               // 207, #ff5fff
	{color.RGBA{0xff, 0x87, 0x00, 0xff}, ""},               // 208, #ff8700
	{color.RGBA{0xff, 0x87, 0x5f, 0xff}, ""},               // 209, #ff875f
	{color.RGBA{0xff, 0x87, 0x87, 0xff}, ""},               // 210, #ff8787
	{color.RGBA{0xff, 0x87, 0xaf, 0xff}, ""},               // 211, #ff87af
	{color.RGBA{0xff, 0x87, 0xd7, 0xff}, ""},               // 212, #ff87d7
	{color.RGBA{0xff, 0x87, 0xff, 0xff}, ""},               // 213, #ff87ff
	{color.RGBA{0xff, 0xaf, 0x00, 0xff}, ""},               // 214, #ffaf00
	{color.RGBA{0xff, 0xaf, 0x5f, 0xff}, ""},               // 215, #ffaf5f
	{color.RGBA{0xff, 0xaf, 0x87, 0xff}, ""},               // 216, #ffaf87
	{color.RGBA{0xff, 0xaf, 0xaf, 0xff}, ""},               // 217, #ffafaf
	{color.RGBA{0xff, 0xaf, 0xd7, 0xff}, ""},               // 218, #ffafd7
	{color.RGBA{0xff, 0xaf, 0xff, 0xff}, ""},               // 219, #ffafff
	{color.RGBA{0xff, 0xd7, 0x00, 0xff}, ""},               // 220, #ffd700
	{color.RGBA{0xff, 0xd7, 0x5f, 0xff}, ""},               // 221, #ffd75f
	{color.RGBA{0xff, 0xd7, 0x87, 0xff}, ""},               // 222, #ffd787
	{color.RGBA{0xff, 0xd7, 0xaf, 0xff}, ""},               // 223, #ffd7af
	{color.RGBA{0xff, 0xd7, 0xd7, 0xff}, ""},               // 224, #ffd7d7
	{color.RGBA{0xff, 0xd7, 0xff, 0xff}, ""},               // 225, #ffd7ff
	{color.RGBA{0xff, 0xff, 0x00, 0xff}, ""},               // 226, #ffff00
	{color.RGBA{0xff, 0xff, 0x5f, 0xff}, ""},               // 227, #ffff5f
	{color.RGBA{0xff, 0xff, 0x87, 0xff}, ""},               // 228, #ffff87
	{color.RGBA{0xff, 0xff, 0xaf, 0xff}, ""},               // 229, #ffffaf
	{color.RGBA{0xff, 0xff, 0xd7, 0xff}, ""},               // 230, #ffffd7
	{color.RGBA{0xff, 0xff, 0xff, 0xff}, ""},               // 231, #ffffff
	{color.RGBA{0x08, 0x08, 0x08, 0xff}, "gray 23"},        // 232, #080808
	{color.RGBA{0x12, 0x12, 0x12, 0xff}, "gray 22"},        // 233, #121212
	{color.RGBA{0x1c, 0x1c, 0x1c, 0xff}, "gray 21"},        // 234, #1c1c1c
	{color.RGBA{0x26, 0x26, 0x26, 0xff}, "gray 20"},        // 235, #262626
	{color.RGBA{0x30, 0x30, 0x30, 0xff}, "gray 19"},        // 236, #303030
	{color.RGBA{0x3a, 0x3a, 0x3a, 0xff}, "gray 18"},        // 237, #3a3a3a
	{color.RGBA{0x44, 0x44, 0x44, 0xff}, "gray 17"},        // 238, #444444
	{color.RGBA{0x4e, 0x4e, 0x4e, 0xff}, "gray 16"},        // 239, #4e4e4e
	{color.RGBA{0x58, 0x58, 0x58, 0xff}, "gray 15"},        // 240, #585858
	{color.RGBA{0x62, 0x62, 0x62, 0xff}, "gray 14"},        // 241, #626262
	{color.RGBA{0x6c, 0x6c, 0x6c, 0xff}, "gray 13"},        // 242, #6c6c6c
	{color.RGBA{0x76, 0x76, 0x76, 0xff}, "gray 12"},        // 243, #767676
	{color.RGBA{0x80, 0x80, 0x80, 0xff}, "gray 11"},        // 244, #808080
	{color.RGBA{0x8a, 0x8a, 0x8a, 0xff}, "gray 10"},        // 245, #8a8a8a
	{color.RGBA{0x94, 0x94, 0x94, 0xff}, "gray 9"},         // 246, #949494
	{color.RGBA{0x9e, 0x9e, 0x9e, 0xff}, "gray 8"},         // 247, #9e9e9e
	{color.RGBA{0xa8, 0xa8, 0xa8, 0xff}, "gray 7"},         // 248, #a8a8a8
	{color.RGBA{0xb2, 0xb2, 0xb2, 0xff}, "gray 6"},         // 249, #b2b2b2
	{color.RGBA{0xbc, 0xbc, 0xbc, 0xff}, "gray 5"},         // 250, #bcbcbc
	{color.RGBA{0xc6, 0xc6, 0xc6, 0xff}, "gray 4"},         // 251, #c6c6c6
	{color.RGBA{0xd0, 0xd0, 0xd0, 0xff}, "gray 3"},         // 252, #d0d0d0
	{color.RGBA{0xda, 0xda, 0xda, 0xff}, "gray 2"},         // 253, #dadada
	{color.RGBA{0xe4, 0xe4, 0xe4, 0xff}, "gray 1"},         // 254, #e4e4e4
	{color.RGBA{0xee, 0xee, 0xee, 0xff}, "white"},          // 255, #eeeeee
}
