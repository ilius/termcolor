package termcolor

type RoundMode uint8

const (
	RoundCloser RoundMode = iota
	RoundDown
	RoundUp
)

func (mode RoundMode) isValid() bool {
	switch mode {
	case RoundCloser:
		return true
	case RoundDown:
		return true
	case RoundUp:
		return true
	}
	return false
}
