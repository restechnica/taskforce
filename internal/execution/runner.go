package execution

import (
	"errors"
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/extensions/osext"
	"os"
	"os/exec"
	"strings"
)

type Runner struct {
	Root config.Root
}

func (runner Runner) RunCommand(command config.Command) (err error) {
	var arguments = strings.Fields(command.Expression)
	var process = exec.Command(arguments[0], arguments[1:]...)

	if command.HasDirectory() {
		process.Dir = osext.ExpandTilde(command.Directory)
	}

	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	return process.Run()
}

func (runner Runner) RunCommandByName(name string) (err error) {
	var command config.Command

	if command, err = runner.Root.GetCommandByName(name); err != nil {
		return
	}

	return runner.RunCommand(command)
}

func (runner Runner) RunInstruction(instruction config.Instruction) (err error) {
	switch instruction.Type {
	case "command":
		return runner.RunCommandByName(instruction.Reference)
	case "task":
		return runner.RunTaskByName(instruction.Reference)
	default:
		err = errors.New("no matching instruction type found")
		return
	}
}

func (runner Runner) RunScript(script config.Script) (err error) {
	for _, instruction := range script.Instructions {
		if err = runner.RunInstruction(instruction); err != nil {
			return
		}
	}
	return
}

func (runner Runner) RunTask(task config.Task) (err error) {
	for _, script := range task.Scripts {
		if err = runner.RunScript(script); err != nil {
			return
		}
	}
	return
}

func (runner Runner) RunTaskByName(name string) (err error) {
	var task config.Task

	if task, err = runner.Root.GetTaskByName(name); err != nil {
		return
	}
	return runner.RunTask(task)
}
