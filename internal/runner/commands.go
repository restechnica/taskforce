package runner

import (
	"bufio"
	"fmt"
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/extensions/osext"
	"github.com/restechnica/taskforce/internal/extensions/shellext"
	"io"
	"os/exec"
)

func RunCommand(command config.Command) (err error) {
	var executable string
	var arguments []string

	if executable, arguments, err = shellext.SplitShellCommand(command.Expression); err != nil {
		return
	}

	var execution = exec.Command(executable, arguments...)

	if command.HasDirectory() {
		execution.Dir = osext.ExpandTilde(command.Directory)
	}

	var stdout, stderr io.ReadCloser

	if stdout, err = execution.StdoutPipe(); err != nil {
		return
	}

	if stderr, err = execution.StderrPipe(); err != nil {
		return
	}

	var stdoutScanner, stderrScanner = bufio.NewScanner(stdout), bufio.NewScanner(stderr)

	go printScanner(stdoutScanner)
	go printScanner(stderrScanner)

	if err = execution.Start(); err != nil {
		return
	}

	if err = execution.Wait(); err != nil {
		return
	}

	return
}

func printScanner(scanner *bufio.Scanner) {
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
