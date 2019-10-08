package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/extensions/slicext"
	"github.com/restechnica/taskforce/internal/hcl"
	"github.com/restechnica/taskforce/internal/runner"
	"log"
	"os"
	"path"
)

const hclFileName = "taskforce.hcl"
const dotenvFileName = ".env"

func main() {
	var err error
	var command config.Command
	var configuration config.Root
	var output, workingDirectory string

	if err = godotenv.Load(dotenvFileName); err != nil {
		log.Println(err)
		log.Println("Failed to load .env file, proceeding without .env variables")
	}

	if workingDirectory, err = os.Getwd(); err != nil {
		log.Fatal(err)
	}

	var filePath = path.Join(workingDirectory, hclFileName)

	if configuration, err = hcl.LoadHCLFile(filePath); err != nil {
		log.Fatal(err)
	}

	if command, err = slicext.Filter(configuration.Commands, func(command config.Command) bool {
		return command.HasName(os.Args[1])
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
