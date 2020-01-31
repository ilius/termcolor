package termcolor

import (
	"image/color"
	"testing"

	"github.com/ilius/is"
)

func uint8Ptr(x uint8) *uint8 {
	x2 := x
	return &x2
}

func TestClosestToRGB(t *testing.T) {
	type Case struct {
		input        color.RGBA
		mode         RoundMode
		grayMaxDelta *uint8
		expectedErr  string
		expectedCode uint8
	}
	test := func(tc Case) {
		is := is.New(t).AddMsg("input=%v, mode=%v", tc.input, tc.mode)
		actual, err := ClosestToRGB(&ClosestToRGBInput{
			Target:       tc.input,
			RoundMode:    tc.mode,
			GrayMaxDelta: tc.grayMaxDelta,
		})
		if tc.expectedErr != "" {
			is.ErrMsg(err, tc.expectedErr)
			is.Nil(actual)
			return
		}
		expectedCode := tc.expectedCode
		expectedRGBA := Colors[expectedCode].RGBA
		is.Equal(actual.Code, expectedCode)
		is.Equal(actual.RGBA, expectedRGBA)
	}
	test(Case{
		input:        color.RGBA{95, 175, 95, 255},
		mode:         RoundCloser,
		expectedCode: 71,
	})
	test(Case{
		input:        color.RGBA{93, 177, 97, 255},
		mode:         RoundCloser,
		expectedCode: 71,
	})
	test(Case{
		input:        color.RGBA{93, 177, 97, 255},
		mode:         RoundDown,
		expectedCode: 35, // {0, 175, 95, 255}
	})
	test(Case{
		input:        color.RGBA{93, 177, 97, 255},
		mode:         RoundUp,
		expectedCode: 78, // {95, 215, 135, 255}
	})
	test(Case{
		input:        color.RGBA{8, 8, 8, 255},
		mode:         RoundCloser,
		expectedCode: 232, // {8, 8, 8, 255}
	})
	test(Case{
		input:        color.RGBA{9, 9, 9, 255},
		mode:         RoundCloser,
		expectedCode: 232, // {8, 8, 8, 255}
	})
	test(Case{
		input:        color.RGBA{10, 10, 10, 255},
		mode:         RoundCloser,
		expectedCode: 232, // {8, 8, 8, 255}
	})
	test(Case{
		input:        color.RGBA{8, 9, 8, 255},
		mode:         RoundCloser,
		expectedCode: 232, // {8, 8, 8, 255}
	})
	test(Case{
		input:        color.RGBA{8, 8, 10, 255},
		mode:         RoundCloser,
		expectedCode: 232, // {8, 8, 8, 255}
	})
	test(Case{
		input:        color.RGBA{6, 8, 6, 255},
		mode:         RoundCloser,
		expectedCode: 232, // {8, 8, 8, 255}
	})
	test(Case{
		input:        color.RGBA{7, 6, 8, 255},
		mode:         RoundCloser,
		expectedCode: 232, // {8, 8, 8, 255}
	})
	test(Case{
		input:        color.RGBA{6, 8, 6, 255},
		mode:         RoundCloser,
		grayMaxDelta: uint8Ptr(2),
		expectedCode: 232, // {8, 8, 8, 255}
	})
	test(Case{
		input:        color.RGBA{6, 8, 6, 255},
		mode:         RoundCloser,
		grayMaxDelta: uint8Ptr(1),
		expectedCode: 0, // {0, 0, 0, 255}
	})
}

func TestDivideRound(t *testing.T) {
	type Case struct {
		a        uint8
		b        uint8
		mode     RoundMode
		expected uint8
	}
	test := func(tc Case) {
		is := is.New(t).AddMsg("a=%v, b=%v, mode=%v", tc.a, tc.b, tc.mode)
		actual := divideRound(tc.a, tc.b, tc.mode)
		is.Equal(actual, tc.expected)
	}
	test(Case{
		a:        0,
		b:        10,
		mode:     RoundDown,
		expected: 0,
	})
	test(Case{
		a:        0,
		b:        10,
		mode:     RoundUp,
		expected: 0,
	})
	test(Case{
		a:        30,
		b:        10,
		mode:     RoundDown,
		expected: 3,
	})
	test(Case{
		a:        30,
		b:        10,
		mode:     RoundUp,
		expected: 3,
	})
	test(Case{
		a:        31,
		b:        10,
		mode:     RoundDown,
		expected: 3,
	})
	test(Case{
		a:        31,
		b:        10,
		mode:     RoundUp,
		expected: 4,
	})
	test(Case{
		a:        31,
		b:        10,
		mode:     RoundCloser,
		expected: 3,
	})
	test(Case{
		a:        39,
		b:        10,
		mode:     RoundDown,
		expected: 3,
	})
	test(Case{
		a:        39,
		b:        10,
		mode:     RoundUp,
		expected: 4,
	})
	test(Case{
		a:        39,
		b:        10,
		mode:     RoundCloser,
		expected: 4,
	})
	test(Case{
		a:        38,
		b:        10,
		mode:     RoundCloser,
		expected: 4,
	})
	test(Case{
		a:        37,
		b:        10,
		mode:     RoundCloser,
		expected: 4,
	})
	test(Case{
		a:        36,
		b:        10,
		mode:     RoundCloser,
		expected: 4,
	})
	test(Case{
		a:        35,
		b:        10,
		mode:     RoundCloser,
		expected: 3,
	})
}

func TestRoundValue(t *testing.T) {
	type Case struct {
		value     uint8
		mode      RoundMode
		expected  uint8
		expectErr string
	}
	test := func(tc Case) {
		is := is.New(t).AddMsg("value=%v, mode=%v", tc.value, tc.mode)
		actualIndex, err := roundValue(tc.value, tc.mode)
		actual := values[actualIndex]
		is.Equal(actual, tc.expected)
		if tc.expectErr != "" {
			is.ErrMsg(err, tc.expectErr)
		} else {
			is.NotErr(err)
		}
	}
	test(Case{
		value:     0,
		mode:      RoundMode(10),
		expected:  0,
		expectErr: "roundValue: invalid mode",
	})
	test(Case{
		value:    0,
		mode:     RoundCloser,
		expected: 0,
	})
	test(Case{
		value:    47,
		mode:     RoundCloser,
		expected: 0,
	})
	test(Case{
		value:    48,
		mode:     RoundCloser,
		expected: 95,
	})
	test(Case{
		value:    95,
		mode:     RoundCloser,
		expected: 95,
	})
	test(Case{
		value:    115,
		mode:     RoundCloser,
		expected: 95,
	})
	test(Case{
		value:    116,
		mode:     RoundCloser,
		expected: 135,
	})
	test(Case{
		value:    135,
		mode:     RoundCloser,
		expected: 135,
	})
	test(Case{
		value:    155,
		mode:     RoundCloser,
		expected: 135,
	})
	test(Case{
		value:    156,
		mode:     RoundCloser,
		expected: 175,
	})
	test(Case{
		value:    175,
		mode:     RoundCloser,
		expected: 175,
	})
	test(Case{
		value:    195,
		mode:     RoundCloser,
		expected: 175,
	})
	test(Case{
		value:    196,
		mode:     RoundCloser,
		expected: 215,
	})
	test(Case{
		value:    215,
		mode:     RoundCloser,
		expected: 215,
	})
	test(Case{
		value:    235,
		mode:     RoundCloser,
		expected: 215,
	})
	test(Case{
		value:    236,
		mode:     RoundCloser,
		expected: 255,
	})
	test(Case{
		value:    255,
		mode:     RoundCloser,
		expected: 255,
	})

	test(Case{
		value:    0,
		mode:     RoundDown,
		expected: 0,
	})
	test(Case{
		value:    1,
		mode:     RoundDown,
		expected: 0,
	})
	test(Case{
		value:    94,
		mode:     RoundDown,
		expected: 0,
	})
	test(Case{
		value:    95,
		mode:     RoundDown,
		expected: 95,
	})
	test(Case{
		value:    96,
		mode:     RoundDown,
		expected: 95,
	})
	test(Case{
		value:    134,
		mode:     RoundDown,
		expected: 95,
	})
	test(Case{
		value:    135,
		mode:     RoundDown,
		expected: 135,
	})
	test(Case{
		value:    136,
		mode:     RoundDown,
		expected: 135,
	})
	test(Case{
		value:    174,
		mode:     RoundDown,
		expected: 135,
	})
	test(Case{
		value:    175,
		mode:     RoundDown,
		expected: 175,
	})
	test(Case{
		value:    176,
		mode:     RoundDown,
		expected: 175,
	})
	test(Case{
		value:    214,
		mode:     RoundDown,
		expected: 175,
	})
	test(Case{
		value:    215,
		mode:     RoundDown,
		expected: 215,
	})
	test(Case{
		value:    216,
		mode:     RoundDown,
		expected: 215,
	})
	test(Case{
		value:    254,
		mode:     RoundDown,
		expected: 215,
	})
	test(Case{
		value:    255,
		mode:     RoundDown,
		expected: 255,
	})

	test(Case{
		value:    0,
		mode:     RoundUp,
		expected: 0,
	})
	test(Case{
		value:    1,
		mode:     RoundUp,
		expected: 95,
	})
	test(Case{
		value:    94,
		mode:     RoundUp,
		expected: 95,
	})
	test(Case{
		value:    95,
		mode:     RoundUp,
		expected: 95,
	})
	test(Case{
		value:    96,
		mode:     RoundUp,
		expected: 135,
	})
	test(Case{
		value:    134,
		mode:     RoundUp,
		expected: 135,
	})
	test(Case{
		value:    135,
		mode:     RoundUp,
		expected: 135,
	})
	test(Case{
		value:    136,
		mode:     RoundUp,
		expected: 175,
	})
	test(Case{
		value:    174,
		mode:     RoundUp,
		expected: 175,
	})
	test(Case{
		value:    175,
		mode:     RoundUp,
		expected: 175,
	})
	test(Case{
		value:    176,
		mode:     RoundUp,
		expected: 215,
	})
	test(Case{
		value:    214,
		mode:     RoundUp,
		expected: 215,
	})
	test(Case{
		value:    215,
		mode:     RoundUp,
		expected: 215,
	})
	test(Case{
		value:    216,
		mode:     RoundUp,
		expected: 255,
	})
	test(Case{
		value:    254,
		mode:     RoundUp,
		expected: 255,
	})
	test(Case{
		value:    255,
		mode:     RoundUp,
		expected: 255,
	})
}
