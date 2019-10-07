package runner

import (
	"github.com/kballard/go-shellquote"
	"github.com/restechnica/taskforce/internal/config"
	"os"
	"os/exec"
	"strings"
)

func RunCommand(command config.Command) (string, error) {
	var output []byte
	var err error
	var splitCommand []string

	if splitCommand, err = shellquote.Split(command.Expression); err != nil {
		return "", err
	}

	var executable = splitCommand[0]
	var arguments = splitCommand[1:]

	var execution = exec.Command(executable, arguments...)
	execution.Env = os.Environ()

	if command.HasDirectory() {
		if strings.HasPrefix(command.Directory, "~") {
			var home = os.Getenv("HOME")
			command.Directory = strings.Replace(command.Directory, "~", home, 1)
		}
	}

	if output, err = execution.CombinedOutput(); err != nil {
		return string(output), err
	}

	return string(output), nil
}
