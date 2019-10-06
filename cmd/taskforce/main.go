package main

import (
	"fmt"
	"github.com/restechnica/taskforce/internal/arrays"
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/hcl"
	"log"
	"os/exec"
)

func main() {
	var configuration = hcl.Parse("./assets/example2.hcl")

	var command, _ = arrays.Filter(configuration.Commands, func(command config.Command) bool {
		return command.Name == "build"
	})

	var cmd = exec.Command(command.Executable, command.Arguments...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Printf("%s", out)
	fmt.Printf("%+v\n", configuration)
}
