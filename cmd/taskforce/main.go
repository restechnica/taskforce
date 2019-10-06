package main

import (
	"fmt"
	"github.com/kballard/go-shellquote"
	"github.com/restechnica/taskforce/internal/arrays"
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/hcl"
	"log"
	"os"
	"os/exec"
)

func main() {
	var configurationPath, workingDirectory string
	var err error
	var configuration config.Root

	if workingDirectory, err = os.Getwd(); err != nil {
		log.Fatalf("%s", err)
		return
	}

	if configurationPath, err = config.Find(workingDirectory); err != nil {
		log.Fatalf("%s", err)
		return
	}

	if configuration, err = hcl.Parse(configurationPath); err != nil {
		log.Fatalf("%s", err)
		return
	}

	var command, _ = arrays.Filter(configuration.Commands, func(command config.Command) bool {
		return command.Name == "build"
	})

	var splitCommand, _ = shellquote.Split(command.Expression)
	var executable = splitCommand[0]
	var arguments = splitCommand[1:]

	var cmd = exec.Command(executable, arguments...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Printf("%s", out)
	fmt.Printf("%+v\n", configuration)
}
