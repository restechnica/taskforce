package execution

import (
	"errors"
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/extensions/osext"
	"github.com/restechnica/taskforce/internal/shell"
	"os"
	"os/exec"
)

type Runner struct {
	Configuration config.Root
}

func (runner Runner) RunCommand(command config.Command) (err error) {
	var arguments []string

	if arguments, err = shell.Parse(command.Text); err != nil {
		return
	}

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

	if command, err = runner.Configuration.GetCommandByName(name); err != nil {
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

func (runner Runner) RunTask(task config.Task) (err error) {
	for _, instruction := range task.Instructions {
		if err = runner.RunInstruction(instruction); err != nil {
			return
		}
	}
	return
}

func (runner Runner) RunTaskByName(name string) (err error) {
	var task config.Task

	if task, err = runner.Configuration.GetTaskByName(name); err != nil {
		return
	}
	return runner.RunTask(task)
}
