package hcl

import (
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/restechnica/taskforce/internal/config"
	"log"
)

func Parse(fileName string) config.Root {
	var parser = hclparse.NewParser()
	var hclFile, hclDiagnostics = parser.ParseHCLFile(fileName)

	if hclDiagnostics.HasErrors() {
		log.Fatal(hclDiagnostics.Error())
	}

	var taskforceConfiguration config.Root
	hclDiagnostics = gohcl.DecodeBody(hclFile.Body, nil, &taskforceConfiguration)

	if hclDiagnostics.HasErrors() {
		log.Fatal(hclDiagnostics.Error())
	}

	return taskforceConfiguration
}
