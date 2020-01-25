package termcolor

import (
	"fmt"
)

func RGBToHtml(r int, g int, b int) string {
	return fmt.Sprintf("#%.2x%.2x%.2x", r, g, b)
}
