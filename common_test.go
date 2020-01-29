package termcolor

import (
	"github.com/ilius/is"
)

func isFloatBetween(is *is.Is, actual float64, min float64, max float64) {
	is.AddMsg(
		"%v is not in range [%v, %v]",
		actual,
		min,
		max,
	).True(min <= actual && actual <= max)
}
