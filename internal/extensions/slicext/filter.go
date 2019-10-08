package slicext

import (
	"errors"
	"github.com/restechnica/taskforce/internal/config"
)

func Filter(commands []config.Command, predicate func(config.Command) bool) (command config.Command, err error) {
	for _, command = range commands {
		if predicate(command) {
			return
		}
	}

	err = errors.New("no matching command found")
	return
}
