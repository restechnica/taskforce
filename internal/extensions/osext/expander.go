package osext

import (
	"os"
	"strings"
)

const n, tilde = 1, "~"

var home = os.Getenv("HOME")

func ExpandTilde(target string) string {
	if strings.HasPrefix(target, tilde) {
		return strings.Replace(target, tilde, home, n)
	}

	return target
}
