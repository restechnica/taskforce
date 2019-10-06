package main

import (
	"fmt"
	"github.com/restechnica/taskforce/internal/arrays"
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/runner"
	"log"
	"os"
)

func main() {
	var err error
	var command config.Command
	var configuration config.Root
	var output, workingDirectory string

	if workingDirectory, err = os.Getwd(); err != nil {
		log.Fatal(err)
	}

	if configuration, err = config.Load(workingDirectory); err != nil {
		log.Fatal(err)
	}

	if command, err = arrays.Filter(configuration.Commands, func(command config.Command) bool {
		return command.HasName("build")
	}); err != nil {
		log.Fatal(err)
	}

	if output, err = runner.RunCommand(command); err != nil {
		log.Print(output)
		log.Fatal(err)
	}

	fmt.Printf(output)
	fmt.Printf("%+v\n", configuration)
}
