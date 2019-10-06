package config

import (
	"github.com/kballard/go-shellquote"
	"os/exec"
)

type Command struct {
	Name string `hcl:"name,label"`

	Expression string `hcl:"expression,attr"`
}

func (command Command) Run() (string, error) {
	var output []byte
	var err error
	var splitCommand []string

	if splitCommand, err = shellquote.Split(command.Expression); err != nil {
		return "", err
	}

	var executable = splitCommand[0]
	var arguments = splitCommand[1:]

	var execution = exec.Command(executable, arguments...)

	if output, err = execution.CombinedOutput(); err != nil {
		return string(output), err
	}

	return string(output), nil
}

func (command Command) HasName(name string) bool {
	return command.Name == name
}
