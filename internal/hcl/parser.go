package hcl

import (
	"errors"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/restechnica/taskforce/internal/config"
	"github.com/zclconf/go-cty/cty"
	"os"
	"strings"
)

func Parse(filePath string) (config.Root, error) {
	var hclFile *hcl.File
	var hclDiagnostics hcl.Diagnostics
	var taskforceConfiguration config.Root

	var hclParser = hclparse.NewParser()

	if hclFile, hclDiagnostics = hclParser.ParseHCLFile(filePath); hclDiagnostics.HasErrors() {
		return taskforceConfiguration, errors.New(hclDiagnostics.Error())
	}

	var ctx = &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"env": cty.ObjectVal(getEnv()),
		},
	}

	if hclDiagnostics = gohcl.DecodeBody(hclFile.Body, ctx, &taskforceConfiguration); hclDiagnostics.HasErrors() {
		return taskforceConfiguration, errors.New(hclDiagnostics.Error())
	}

	return taskforceConfiguration, nil
}

func getEnv() map[string]cty.Value {
	var env = make(map[string]cty.Value)

	for _, pair := range os.Environ() {
		if index := strings.IndexByte(pair, '='); index >= 0 {
			var key = pair[:index]
			var value = pair[index+1:]

			env[key] = cty.StringVal(value)
		}
	}

	return env
}
