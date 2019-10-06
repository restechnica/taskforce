package arrays

import (
	"errors"
	"github.com/restechnica/taskforce/internal/config"
)

func Filter(commands []config.Command, predicate func(config.Command) bool) (config.Command, error) {
	for _, command := range commands {
		if predicate(command) {
			return command, nil
		}
	}

	return config.Command{}, errors.New("No matching command found")
}
