package termcolor

import (
	"image/color"
	"testing"

	"github.com/ilius/is/v2"
)

func TestClosestToRGB(t *testing.T) {
	type Case struct {
		grayMaxDelta *uint8
		expectedErr  string
		input        color.RGBA
		mode         RoundMode
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
		input:       color.RGBA{255, 0, 0, 255},
		mode:        RoundMode(10),
		expectedErr: "invalid RoundMode",
	})
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
		input:        color.RGBA{87, 87, 87, 255},
		mode:         RoundUp,
		expectedCode: 240, // 88, 88, 88
	})
	test(Case{
		input:        color.RGBA{179, 179, 179, 255},
		mode:         RoundDown,
		expectedCode: 249, // 178, 178, 178
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
	test(Case{
		input:        color.RGBA{120, 200, 50, 255},
		mode:         RoundCloser,
		expectedCode: 113, // 135, 215, 95
	})
	test(Case{
		input:        color.RGBA{120, 200, 50, 255},
		mode:         RoundUp,
		expectedCode: 113, // 135, 215, 95
	})
	test(Case{
		input:        color.RGBA{120, 200, 50, 255},
		mode:         RoundDown,
		expectedCode: 70, // 95, 175, 0
	})
}
