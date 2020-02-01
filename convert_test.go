package termcolor

import (
	"image/color"
	"testing"

	"github.com/ilius/is"
)

func TestClosestToRGB(t *testing.T) {
	type Case struct {
		input        color.RGBA
		mode         RoundMode
		grayMaxDelta *uint8
		expectedErr  string
		expectedCode uint8
	}
	test := func(tc Case) {
		is := is.New(t).Lax().AddMsg(
			"input=%v, mode=%v",
			tc.input, tc.mode,
		)
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
		is = is.AddMsg(
			"expect=%v, actual=%v",
			tc.expectedCode, actual.Code,
		)
		if is.Equal(actual.RGBA, Colors[tc.expectedCode].RGBA) {
			is.Equal(actual.Code, tc.expectedCode)
		}
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
		expectedCode: 35, // 0, 175, 9
	})
	test(Case{
		input:        color.RGBA{93, 177, 97, 255},
		mode:         RoundUp,
		expectedCode: 78, // 95, 215, 135
	})
	test(Case{
		input:        color.RGBA{8, 8, 8, 255},
		mode:         RoundCloser,
		expectedCode: 232, // 8, 8, 8
	})
	test(Case{
		input:        color.RGBA{9, 9, 9, 255},
		mode:         RoundCloser,
		expectedCode: 232, // 8, 8, 8
	})
	test(Case{
		input:        color.RGBA{10, 10, 10, 255},
		mode:         RoundCloser,
		expectedCode: 232, // 8, 8, 8
	})
	test(Case{
		input:        color.RGBA{8, 9, 8, 255},
		mode:         RoundCloser,
		expectedCode: 232, // 8, 8, 8
	})
	test(Case{
		input:        color.RGBA{8, 8, 10, 255},
		mode:         RoundCloser,
		expectedCode: 232, // 8, 8, 8
	})
	test(Case{
		input:        color.RGBA{6, 8, 6, 255},
		mode:         RoundCloser,
		expectedCode: 232, // 8, 8, 8
	})
	test(Case{
		input:        color.RGBA{7, 6, 8, 255},
		mode:         RoundCloser,
		expectedCode: 232, // 8, 8, 8
	})
	test(Case{
		input:        color.RGBA{6, 8, 6, 255},
		mode:         RoundCloser,
		grayMaxDelta: uint8Ptr(2),
		expectedCode: 232, // 8, 8, 8
	})
	test(Case{
		input:        color.RGBA{6, 8, 6, 255},
		mode:         RoundCloser,
		grayMaxDelta: uint8Ptr(1),
		expectedCode: 0, // 0, 0, 0
	})

	test(Case{
		input:        color.RGBA{170, 0, 0, 255},
		mode:         RoundCloser,
		expectedCode: 1, // 170, 0, 0
	})
	test(Case{
		input:        color.RGBA{165, 3, 5, 255},
		mode:         RoundCloser,
		expectedCode: 1, // 170, 0, 0
	})
	test(Case{
		input:        color.RGBA{0, 170, 0, 255},
		mode:         RoundCloser,
		expectedCode: 2, // 0, 170, 0
	})
	test(Case{
		input:        color.RGBA{10, 165, 11, 255},
		mode:         RoundCloser,
		expectedCode: 2, // 0, 170, 0
	})
	test(Case{
		input:        color.RGBA{170, 85, 0, 255},
		mode:         RoundCloser,
		expectedCode: 3, // 170, 85, 0
	})
	test(Case{
		input:        color.RGBA{175, 80, 4, 255},
		mode:         RoundCloser,
		expectedCode: 3, // 170, 85, 0
	})
	test(Case{
		input:        color.RGBA{0, 0, 170, 255},
		mode:         RoundCloser,
		expectedCode: 4, // 0, 0, 170
	})
	test(Case{
		input:        color.RGBA{4, 3, 161, 255},
		mode:         RoundCloser,
		expectedCode: 4, // 0, 0, 170
	})
	test(Case{
		input:        color.RGBA{170, 0, 170, 255},
		mode:         RoundCloser,
		expectedCode: 5, // 170, 0, 170
	})
	test(Case{
		input:        color.RGBA{165, 7, 174, 255},
		mode:         RoundCloser,
		expectedCode: 5, // 170, 0, 170
	})
	test(Case{
		input:        color.RGBA{170, 170, 0, 255},
		mode:         RoundCloser,
		expectedCode: 142, // 175, 175, 0
	})
	test(Case{
		input:        color.RGBA{0, 170, 170, 255},
		mode:         RoundCloser,
		expectedCode: 6, // 0, 170, 170
	})
	test(Case{
		input:        color.RGBA{9, 177, 162, 255},
		mode:         RoundCloser,
		expectedCode: 6, // 0, 170, 170
	})
	test(Case{
		input:        color.RGBA{185, 185, 185, 255},
		mode:         RoundCloser,
		expectedCode: 7, // 185, 185, 185
	})
	test(Case{
		input:        color.RGBA{180, 190, 183, 255},
		mode:         RoundCloser,
		expectedCode: 7, // 185, 185, 185
	})
	test(Case{
		input:        color.RGBA{85, 85, 85, 255},
		mode:         RoundCloser,
		expectedCode: 8, // 85, 85, 85
	})
	test(Case{
		input:        color.RGBA{83, 80, 90, 255},
		mode:         RoundCloser,
		expectedCode: 8, // 85, 85, 85
	})
	test(Case{
		input:        color.RGBA{83, 80, 85, 255},
		mode:         RoundUp,
		expectedCode: 8, // 85, 85, 85
	})
	test(Case{
		input:        color.RGBA{85, 87, 87, 255},
		mode:         RoundDown,
		expectedCode: 8, // 85, 85, 85
	})
	test(Case{
		input:        color.RGBA{255, 85, 85, 255},
		mode:         RoundCloser,
		expectedCode: 9, // 255, 85, 85
	})
	test(Case{
		input:        color.RGBA{250, 79, 92, 255},
		mode:         RoundCloser,
		expectedCode: 9, // 255, 85, 85
	})
	test(Case{
		input:        color.RGBA{85, 255, 85, 255},
		mode:         RoundCloser,
		expectedCode: 10, // 85, 255, 85
	})
	test(Case{ // on the edge
		input:        color.RGBA{91, 249, 88, 255},
		mode:         RoundCloser,
		expectedCode: 10, // 85, 255, 85
	})
	test(Case{ // on the edge, equal distance, pallete is preferred
		input:        color.RGBA{91, 249, 89, 255},
		mode:         RoundCloser,
		expectedCode: 10, // 95 255 95
	})
	test(Case{ // on the edge
		input:        color.RGBA{92, 249, 89, 255},
		mode:         RoundCloser,
		expectedCode: 83, // 95 255 95
	})
	test(Case{
		input:        color.RGBA{85, 255, 85, 255},
		mode:         RoundUp,
		expectedCode: 10, // 85, 255, 85
	})
	test(Case{
		input:        color.RGBA{255, 255, 85, 255},
		mode:         RoundCloser,
		expectedCode: 11, // 255, 255, 85
	})
	test(Case{
		input:        color.RGBA{250, 251, 87, 255},
		mode:         RoundCloser,
		expectedCode: 11, // 255, 255, 85
	})
	test(Case{
		input:        color.RGBA{85, 85, 255, 255},
		mode:         RoundCloser,
		expectedCode: 12, // 85, 85, 255
	})
	test(Case{
		input:        color.RGBA{83, 88, 249, 255},
		mode:         RoundCloser,
		expectedCode: 12, // 85, 85, 255
	})
	test(Case{
		input:        color.RGBA{255, 85, 255, 255},
		mode:         RoundCloser,
		expectedCode: 13, // 255, 85, 255
	})
	test(Case{
		input:        color.RGBA{251, 83, 252, 255},
		mode:         RoundCloser,
		expectedCode: 13, // 255, 85, 255
	})
	test(Case{
		input:        color.RGBA{87, 255, 253, 255},
		mode:         RoundCloser,
		expectedCode: 14, // 85, 255, 255
	})
	test(Case{
		input:        color.RGBA{255, 255, 255, 255},
		mode:         RoundCloser,
		expectedCode: 15, // 255, 255, 255
	})
	test(Case{
		input:        color.RGBA{255, 255, 254, 255},
		mode:         RoundCloser,
		expectedCode: 15, // 255, 255, 255
	})
	test(Case{
		input:        color.RGBA{255, 254, 254, 255},
		mode:         RoundCloser,
		expectedCode: 15, // 255, 255, 255
	})
	test(Case{
		input:        color.RGBA{254, 253, 253, 255},
		mode:         RoundCloser,
		expectedCode: 15, // 255, 255, 255
	})
	test(Case{ // on the edge
		input:        color.RGBA{246, 247, 246, 255},
		mode:         RoundCloser,
		expectedCode: 255, // 238, 238, 238
	})
	test(Case{ // on the edge
		input:        color.RGBA{246, 246, 246, 255},
		mode:         RoundCloser,
		expectedCode: 255, // 238, 238, 238
	})
	test(Case{
		input:        color.RGBA{246, 246, 246, 255},
		mode:         RoundDown,
		expectedCode: 255, // 238, 238, 238
	})
	test(Case{
		input:        color.RGBA{246, 246, 246, 255},
		mode:         RoundUp,
		expectedCode: 15, // 255, 255, 255
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
