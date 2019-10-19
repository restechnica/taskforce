package hcl

import (
	"errors"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/restechnica/taskforce/internal/config"
	"github.com/restechnica/taskforce/internal/extensions/stringsext"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
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
		var key, value = stringsext.ParseKeyValuePair(pair, '=')
		var ctyStringValue = cty.StringVal(value)
		env[strings.ToLower(key)] = ctyStringValue
		env[strings.ToUpper(key)] = ctyStringValue
	}

	return env
}

func newHCLEvalContext() hcl.EvalContext {
	return hcl.EvalContext{
		Functions: map[string]function.Function{
			"script": ResolveVariableFromScript,
		},
		Variables: map[string]cty.Value{
			"env": cty.ObjectVal(mapEnvironmentVariables()),
		},
	}
}
