package main

import (
	"fmt"
	"github.com/restechnica/taskforce/internal/arrays"
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/hcl"
	"log"
	"os"
)

func main() {
	var err error
	var command config.Command
	var configuration config.Root
	var configurationPath, output, workingDirectory string

	if workingDirectory, err = os.Getwd(); err != nil {
		log.Fatal(err)
	}

	if configurationPath, err = config.Find(workingDirectory); err != nil {
		log.Fatal(err)
	}

	if configuration, err = hcl.Parse(configurationPath); err != nil {
		log.Fatal(err)
	}

	if command, err = arrays.Filter(configuration.Commands, func(command config.Command) bool {
		return command.HasName("build")
	}); err != nil {
		log.Fatal(err)
	}

	if output, err = command.Run(); err != nil {
		log.Print(output)
		log.Fatal(err)
	}

	fmt.Printf(output)
	fmt.Printf("%+v\n", configuration)
}
