package termcolor

import (
	"github.com/ilius/is"
	"testing"
)

func TestRoundValue(t *testing.T) {
	type Case struct {
		value     uint8
		mode      RoundMode
		expected  uint8
		expectErr string
	}
	test := func(tc Case) {
		is := is.New(t).AddMsg("value=%v, mode=%v", tc.value, tc.mode)
		actual, err := roundValue(tc.value, tc.mode)
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
