package tracker

import "fmt"

const (
	Default    = "\x1b[39m"
	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
)

func Red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, Default)
}

func Green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, Default)
}
