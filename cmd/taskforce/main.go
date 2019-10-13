package main

import (
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/environment"
	"github.com/restechnica/taskforce/internal/execution"
	"github.com/restechnica/taskforce/internal/hcl"
	"log"
	"os"
	"path"
)

func main() {
	var err error
	var configuration config.Root
	var workingDirectory string

	if err = environment.Load("./.env"); err != nil {
		log.Println(err)
		log.Println("Failed to load .env file, proceeding without .env variables")
	}

	if workingDirectory, err = os.Getwd(); err != nil {
		log.Fatal(err)
	}

	var filePath = path.Join(workingDirectory, "taskforce.hcl")

	if configuration, err = hcl.LoadHCLFile(filePath); err != nil {
		log.Fatal(err)
	}

	var runner = execution.Runner{Configuration: configuration}

	if err = runner.RunTaskByName(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}
