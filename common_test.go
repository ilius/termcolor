package termcolor

import (
	"github.com/ilius/is/v2"
)

func isFloatBetween(is *is.Is, actual float64, min float64, max float64) {
	is.AddMsg(
		"%v is not in range [%v, %v]",
		actual,
		min,
		max,
	).True(min <= actual && actual <= max)
}

func uint8Ptr(x uint8) *uint8 {
	x2 := x
	return &x2
}
