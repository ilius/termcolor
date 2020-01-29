//go:generate go run gen-byname.go

package termcolor

var ColorLookup = map[string]uint8{
	"black":            0,
	"css:black":        0,
	"#000000":          0,
	"dark red 1":       1,
	"#aa0000":          1,
	"mixed green 1":    2,
	"#00aa00":          2,
	"mixed orange 1":   3,
	"#aa5500":          3,
	"mixed blue 1":     4,
	"#0000aa":          4,
	"purple 1":         5,
	"#aa00aa":          5,
	"cyan 1":           6,
	"#00aaaa":          6,
	"light gray":       7,
	"#b9b9b9":          7,
	"dark gray":        8,
	"#555555":          8,
	"light red":        9,
	"#ff5555":          9,
	"light green":      10,
	"#55ff55":          10,
	"yellow":           11,
	"#ffff55":          11,
	"light blue":       12,
	"#5555ff":          12,
	"light purple":     13,
	"#ff55ff":          13,
	"light cyan":       14,
	"#55ffff":          14,
	"white":            15,
	"css:white":        15,
	"#ffffff":          15,
	"blue 4":           17,
	"#00005f":          17,
	"blue 3":           18,
	"css:darkblue":     18,
	"#000087":          18,
	"blue 2":           19,
	"#0000af":          19,
	"blue 1":           20,
	"css:mediumblue":   20,
	"#0000d7":          20,
	"blue":             21,
	"css:blue":         21,
	"#0000ff":          21,
	"green 4":          22,
	"css:darkgreen":    22,
	"#005f00":          22,
	"blue stone":       23,
	"#005f5f":          23,
	"orient":           24,
	"#005f87":          24,
	"#005faf":          25,
	"#005fd7":          26,
	"#005fff":          27,
	"green 3":          28,
	"css:green":        28,
	"#008700":          28,
	"#00875f":          29,
	"css:darkcyan":     30,
	"#008787":          30,
	"#0087af":          31,
	"#0087d7":          32,
	"#0087ff":          33,
	"green 2":          34,
	"#00af00":          34,
	"#00af5f":          35,
	"#00af87":          36,
	"#00afaf":          37,
	"#00afd7":          38,
	"#00afff":          39,
	"green 1":          40,
	"#00d700":          40,
	"#00d75f":          41,
	"#00d787":          42,
	"#00d7af":          43,
	"#00d7d7":          44,
	"#00d7ff":          45,
	"green":            46,
	"css:lime":         46,
	"#00ff00":          46,
	"#00ff5f":          47,
	"css:springgreen":  48,
	"#00ff87":          48,
	"#00ffaf":          49,
	"#00ffd7":          50,
	"cyan":             51,
	"css:cyan":         51,
	"css:aqua":         51,
	"#00ffff":          51,
	"red 5":            52,
	"#5f0000":          52,
	"#5f005f":          53,
	"#5f0087":          54,
	"#5f00af":          55,
	"#5f00d7":          56,
	"#5f00ff":          57,
	"#5f5f00":          58,
	"#5f5f5f":          59,
	"#5f5f87":          60,
	"#5f5faf":          61,
	"#5f5fd7":          62,
	"#5f5fff":          63,
	"#5f8700":          64,
	"glade green":      65,
	"#5f875f":          65,
	"#5f8787":          66,
	"#5f87af":          67,
	"#5f87d7":          68,
	"#5f87ff":          69,
	"#5faf00":          70,
	"#5faf5f":          71,
	"#5faf87":          72,
	"#5fafaf":          73,
	"tradewind":        74,
	"#5fafd7":          74,
	"#5fafff":          75,
	"#5fd700":          76,
	"#5fd75f":          77,
	"#5fd787":          78,
	"#5fd7af":          79,
	"#5fd7d7":          80,
	"#5fd7ff":          81,
	"#5fff00":          82,
	"#5fff5f":          83,
	"#5fff87":          84,
	"#5fffaf":          85,
	"#5fffd7":          86,
	"#5fffff":          87,
	"red 4":            88,
	"css:darkred":      88,
	"#870000":          88,
	"#87005f":          89,
	"css:darkmagenta":  90,
	"#870087":          90,
	"#8700af":          91,
	"#8700d7":          92,
	"electric violet":  93,
	"#8700ff":          93,
	"#875f00":          94,
	"#875f5f":          95,
	"#875f87":          96,
	"#875faf":          97,
	"#875fd7":          98,
	"#875fff":          99,
	"css:olive":        100,
	"#878700":          100,
	"clay creek":       101,
	"#87875f":          101,
	"#878787":          102,
	"#8787af":          103,
	"#8787d7":          104,
	"#8787ff":          105,
	"#87af00":          106,
	"#87af5f":          107,
	"#87af87":          108,
	"#87afaf":          109,
	"#87afd7":          110,
	"#87afff":          111,
	"#87d700":          112,
	"#87d75f":          113,
	"#87d787":          114,
	"#87d7af":          115,
	"#87d7d7":          116,
	"css:lightskyblue": 117,
	"#87d7ff":          117,
	"css:chartreuse":   118,
	"#87ff00":          118,
	"#87ff5f":          119,
	"#87ff87":          120,
	"#87ffaf":          121,
	"css:aquamarine":   122,
	"#87ffd7":          122,
	"#87ffff":          123,
	"red 3":            124,
	"bright red":       124,
	"#af0000":          124,
	"#af005f":          125,
	"#af0087":          126,
	"#af00af":          127,
	"#af00d7":          128,
	"#af00ff":          129,
	"#af5f00":          130,
	"#af5f5f":          131,
	"#af5f87":          132,
	"#af5faf":          133,
	"#af5fd7":          134,
	"#af5fff":          135,
	"#af8700":          136,
	"#af875f":          137,
	"#af8787":          138,
	"#af87af":          139,
	"#af87d7":          140,
	"#af87ff":          141,
	"#afaf00":          142,
	"#afaf5f":          143,
	"#afaf87":          144,
	"silver chalice":   145,
	"#afafaf":          145,
	"#afafd7":          146,
	"#afafff":          147,
	"#afd700":          148,
	"#afd75f":          149,
	"#afd787":          150,
	"#afd7af":          151,
	"#afd7d7":          152,
	"#afd7ff":          153,
	"#afff00":          154,
	"#afff5f":          155,
	"#afff87":          156,
	"#afffaf":          157,
	"#afffd7":          158,
	"#afffff":          159,
	"red 2":            160,
	"#d70000":          160,
	"#d7005f":          161,
	"#d70087":          162,
	"#d700af":          163,
	"#d700d7":          164,
	"#d700ff":          165,
	"#d75f00":          166,
	"#d75f5f":          167,
	"#d75f87":          168,
	"#d75faf":          169,
	"#d75fd7":          170,
	"#d75fff":          171,
	"#d78700":          172,
	"#d7875f":          173,
	"#d78787":          174,
	"#d787af":          175,
	"#d787d7":          176,
	"#d787ff":          177,
	"#d7af00":          178,
	"#d7af5f":          179,
	"css:tan":          180,
	"#d7af87":          180,
	"#d7afaf":          181,
	"#d7afd7":          182,
	"#d7afff":          183,
	"#d7d700":          184,
	"#d7d75f":          185,
	"#d7d787":          186,
	"#d7d7af":          187,
	"gray 3":           188,
	"#d7d7d7":          188,
	"#d7d7ff":          189,
	"#d7ff00":          190,
	"#d7ff5f":          191,
	"#d7ff87":          192,
	"#d7ffaf":          193,
	"snowy mint":       194,
	"#d7ffd7":          194,
	"css:lightcyan":    195,
	"#d7ffff":          195,
	"red":              196,
	"css:red":          196,
	"#ff0000":          196,
	"#ff005f":          197,
	"#ff0087":          198,
	"#ff00af":          199,
	"#ff00d7":          200,
	"css:fuchsia":      201,
	"css:magenta":      201,
	"#ff00ff":          201,
	"#ff5f00":          202,
	"#ff5f5f":          203,
	"#ff5f87":          204,
	"#ff5faf":          205,
	"#ff5fd7":          206,
	"#ff5fff":          207,
	"css:darkorange":   208,
	"#ff8700":          208,
	"#ff875f":          209,
	"geraldine":        210,
	"#ff8787":          210,
	"#ff87af":          211,
	"#ff87d7":          212,
	"#ff87ff":          213,
	"css:orange":       214,
	"#ffaf00":          214,
	"#ffaf5f":          215,
	"#ffaf87":          216,
	"cornflower lilac": 217,
	"#ffafaf":          217,
	"#ffafd7":          218,
	"#ffafff":          219,
	"css:gold":         220,
	"#ffd700":          220,
	"dandelion":        221,
	"#ffd75f":          221,
	"#ffd787":          222,
	"css:navajowhite":  223,
	"light apricot":    223,
	"#ffd7af":          223,
	"cosmos":           224,
	"#ffd7d7":          224,
	"#ffd7ff":          225,
	"css:yellow":       226,
	"#ffff00":          226,
	"#ffff5f":          227,
	"#ffff87":          228,
	"portafino":        229,
	"#ffffaf":          229,
	"css:cornsilk":     230,
	"cumulus":          230,
	"#ffffd7":          230,
	"gray 24":          232,
	"cod gray":         232,
	"#080808":          232,
	"gray 23":          233,
	"#121212":          233,
	"gray 22":          234,
	"#1c1c1c":          234,
	"gray 21":          235,
	"#262626":          235,
	"gray 20":          236,
	"mine shaft":       236,
	"#303030":          236,
	"gray 19":          237,
	"#3a3a3a":          237,
	"gray 18":          238,
	"#444444":          238,
	"gray 17":          239,
	"#4e4e4e":          239,
	"gray 16":          240,
	"#585858":          240,
	"gray 15":          241,
	"#626262":          241,
	"gray 14":          242,
	"css:dimgray":      242,
	"dove gray":        242,
	"#6c6c6c":          242,
	"gray 13":          243,
	"#767676":          243,
	"gray 12":          244,
	"css:gray":         244,
	"css:grey":         244,
	"#808080":          244,
	"gray 11":          245,
	"#8a8a8a":          245,
	"gray 10":          246,
	"#949494":          246,
	"gray 9":           247,
	"#9e9e9e":          247,
	"gray 8":           248,
	"css:darkgrey":     248,
	"#a8a8a8":          248,
	"gray 7":           249,
	"#b2b2b2":          249,
	"gray 6":           250,
	"css:silver":       250,
	"#bcbcbc":          250,
	"gray 5":           251,
	"#c6c6c6":          251,
	"gray 4":           252,
	"css:lightgray":    252,
	"#d0d0d0":          252,
	"gray 2":           253,
	"css:gainsboro":    253,
	"alto":             253,
	"#dadada":          253,
	"gray 1":           254,
	"mercury":          254,
	"#e4e4e4":          254,
	"gray 0":           255,
	"gallery":          255,
	"#eeeeee":          255,
}
