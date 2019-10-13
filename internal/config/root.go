package config

import (
	"fmt"
)

type Root struct {
	Commands []Command `hcl:"command,block"`
	Tasks    []Task    `hcl:"task,block"`
}

func (root Root) GetCommandByName(name string) (command Command, err error) {
	for _, command = range root.Commands {
		if command.Name == name {
			return
		}
	}

	err = fmt.Errorf("no matching command '%s' found", name)
	return
}

func (root Root) GetTaskByName(name string) (task Task, err error) {
	for _, task = range root.Tasks {
		if task.Name == name {
			return
		}
	}

	err = fmt.Errorf("no matching task '%s' found", name)
	return
}
