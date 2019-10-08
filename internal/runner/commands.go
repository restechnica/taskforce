package runner

import (
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/extensions/osext"
	"github.com/restechnica/taskforce/internal/extensions/shellext"
	"os"
	"os/exec"
)

func RunCommand(command config.Command) (err error) {
	var executable string
	var arguments []string

	if executable, arguments, err = shellext.SplitShellCommand(command.Expression); err != nil {
		return
	}

	var process = exec.Command(executable, arguments...)

	if command.HasDirectory() {
		process.Dir = osext.ExpandTilde(command.Directory)
	}

	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	return process.Run()
}
