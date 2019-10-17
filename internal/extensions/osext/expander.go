package osext

import (
	"github.com/mitchellh/go-homedir"
)

func ExpandTilde(target string) (result string, err error) {
	if result, err = homedir.Expand(target); err != nil {
		return
	}

	return
}
