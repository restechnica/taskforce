package shellext

import (
	"fmt"
	"github.com/kballard/go-shellquote"
	"strings"
)

func SplitShellCommand(command string) (executable string, arguments []string, err error) {
	var fields = strings.Fields(command)

	var test []string

	if test, err = shellquote.Split(command); err != nil {
		return
	}

	fmt.Println(test)

	fmt.Println(fields)

	return fields[0], fields[1:], err
}
