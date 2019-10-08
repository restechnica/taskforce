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

func ParseHCLFile(filePath string) (root config.Root, err error) {
	var hclFile *hcl.File
	var hclDiagnostics hcl.Diagnostics

	var hclParser = hclparse.NewParser()

	if hclFile, hclDiagnostics = hclParser.ParseHCLFile(filePath); hclDiagnostics.HasErrors() {
		err = errors.New(hclDiagnostics.Error())
		return
	}

	var hclEvalContext = newHCLEvalContext()

	if hclDiagnostics = gohcl.DecodeBody(hclFile.Body, &hclEvalContext, &root); hclDiagnostics.HasErrors() {
		err = errors.New(hclDiagnostics.Error())
		return
	}

	return
}

func mapEnvironmentVariables() map[string]cty.Value {
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

func newHCLEvalContext() hcl.EvalContext {
	return hcl.EvalContext{
		Variables: map[string]cty.Value{
			"env": cty.ObjectVal(mapEnvironmentVariables()),
		},
	}
}
