package execution

import (
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/extensions/osext"
	"os"
	"os/exec"
	"strings"
)

func RunCommand(command config.Command) (err error) {
	var arguments = splitExpression(command.Expression)
	var process = exec.Command(arguments[0], arguments[1:]...)

	if command.HasDirectory() {
		process.Dir = osext.ExpandTilde(command.Directory)
	}

	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	return process.Run()
}

func splitExpression(command string) []string {
	return strings.Fields(command)
}
