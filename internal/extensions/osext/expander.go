package osext

import (
	"os"
	"strings"
)

const tilde = "~"

var home = os.Getenv("HOME")

func ExpandTilde(target string) string {
	if strings.HasPrefix(target, tilde) {
		return strings.Replace(target, tilde, home, 1)
	}

	return target
}
