package shellext

import (
	"github.com/kballard/go-shellquote"
)

func SplitShellCommand(command string) (executable string, arguments []string, err error) {
	var split []string

	if split, err = shellquote.Split(command); err != nil {
		return
	}

	executable, arguments = split[0], split[1:]
	return
}
