package hcl

import (
	"errors"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/restechnica/taskforce/internal/config"
)

func Parse(filePath string) (config.Root, error) {
	var hclFile *hcl.File
	var hclDiagnostics hcl.Diagnostics
	var taskforceConfiguration config.Root

	var hclParser = hclparse.NewParser()

	if hclFile, hclDiagnostics = hclParser.ParseHCLFile(filePath); hclDiagnostics.HasErrors() {
		return taskforceConfiguration, errors.New(hclDiagnostics.Error())
	}

	if hclDiagnostics = gohcl.DecodeBody(hclFile.Body, nil, &taskforceConfiguration); hclDiagnostics.HasErrors() {
		return taskforceConfiguration, errors.New(hclDiagnostics.Error())
	}

	return taskforceConfiguration, nil
}
