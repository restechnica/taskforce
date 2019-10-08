package runner

import (
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/extensions/osext"
	"github.com/restechnica/taskforce/internal/extensions/shellext"
	"os/exec"
)

func RunCommand(command config.Command) (output string, err error) {
	var executable string
	var arguments []string

	if executable, arguments, err = shellext.SplitShellCommand(command.Expression); err != nil {
		return
	}

	var execution = exec.Command(executable, arguments...)

	if command.HasDirectory() {
		command.Directory = osext.ExpandTilde(command.Directory)
	}

	var combinedOutput []byte

	if combinedOutput, err = execution.CombinedOutput(); err != nil {
		return
	}

	return string(combinedOutput), nil
}
