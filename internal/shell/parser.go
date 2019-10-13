package shell

import (
	"github.com/kballard/go-shellquote"
	"strings"
)

func Parse(expression string) (arguments []string, err error) {
	var quotedArguments = strings.Fields(expression)

	if arguments, err = shellquote.Split(strings.Join(quotedArguments, " ")); err != nil {
		return
	}

	return
}
