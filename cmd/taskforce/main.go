package main

import (
	"flag"
	"fmt"
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/environment"
	"github.com/restechnica/taskforce/internal/hcl"
	"github.com/restechnica/taskforce/internal/version"
	"log"
	"os"
	"path"
)

func main() {
	var err error
	var configuration config.Root
	var workingDirectory string

	var isVersion bool
	flag.BoolVar(&isVersion, "version", false, "prints the binary version")
	flag.BoolVar(&isVersion, "v", false, "prints the binary version")

	var isCommand bool
	flag.BoolVar(&isCommand, "command", false, "runs a command by the given name")
	flag.BoolVar(&isCommand, "c", false, "runs a command by the given name")

	flag.Parse()

	if isVersion {
		fmt.Println(version.Version)
		return
	}

	if err = environment.Load(".env"); err != nil {
		log.Println(err)
		log.Println("Failed to load .env file, proceeding without .env variables")
	}

	if err = environment.Load("taskforce.env"); err != nil {
		log.Println(err)
		log.Println("Failed to load taskforce.env file, proceeding without .env variables")
	}

	if workingDirectory, err = os.Getwd(); err != nil {
		log.Fatal(err)
	}

	var filePath = path.Join(workingDirectory, "taskforce.hcl")

	if configuration, err = hcl.LoadHCLFile(filePath); err != nil {
		log.Fatal(err)
	}

	var runner = config.Runner{Configuration: configuration}
	var name = flag.Arg(0)

	if isCommand {
		err = runner.RunCommandByName(name)
	} else {
		err = runner.RunTaskByName(name)
	}

	if err != nil {
		log.Fatal(err)
	}
}
