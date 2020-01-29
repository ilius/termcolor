package termcolor

import (
	"errors"
	"fmt"
	"image/color"
)

var errInvalidHexColor = errors.New("invalid hex color")

func RGBToHexColor(r uint8, g uint8, b uint8) string {
	return fmt.Sprintf("#%.2x%.2x%.2x", r, g, b)
}

func RGBAToHexColor(c color.RGBA) string {
	return fmt.Sprintf("#%.2x%.2x%.2x", c.R, c.G, c.B)
}

func ParseHexColor(s string) (c *color.RGBA, err error) {
	if s[0] != '#' {
		return nil, errInvalidHexColor
	}
	c = &color.RGBA{}
	c.A = 0xff
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidHexColor
		return 0
	}
	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidHexColor
	}
	if err != nil {
		return nil, err
	}
	return
}
