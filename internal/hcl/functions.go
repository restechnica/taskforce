package hcl

import (
	"github.com/restechnica/taskforce/internal/extensions/osext"
	"github.com/restechnica/taskforce/internal/variables"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"path/filepath"
)

var resolver = variables.NewResolver()

//var ResolveVariableFromFile = function.New(&function.Spec{
//	Params: []function.Parameter{
//		{
//			Name: "filePath",
//			Type: cty.String,
//		}, {
//			Name: "variableName",
//			Type: cty.String,
//		},
//	},
//	Type: function.StaticReturnType(cty.String),
//	Impl: func(args []cty.Value, retType cty.Type) (value cty.Value, err error) {
//		var filePath = args[0].AsString()
//		var variableName = args[1].AsString()
//
//		if filePath, err = osext.ExpandTilde(filePath); err != nil {
//			return
//		}
//
//		var variableValue string
//
//		if variableValue, err = resolver.(filePath, variableName); err != nil {
//			return
//		}
//
//		value = cty.StringVal(variableValue)
//
//		return
//	},
//})

var ResolveVariableFromScript = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "executable",
			Type: cty.String,
		},
		{
			Name: "filePath",
			Type: cty.String,
		}, {
			Name: "variableName",
			Type: cty.String,
		},
	},
	Type: function.StaticReturnType(cty.String),
	Impl: func(args []cty.Value, retType cty.Type) (value cty.Value, err error) {
		var executable = args[0].AsString()
		var filePath = args[1].AsString()
		var variableName = args[2].AsString()

		if executable, err = osext.ExpandTilde(executable); err != nil {
			return
		}

		if filePath, err = osext.ExpandTilde(filePath); err != nil {
			return
		}

		if filePath, err = filepath.Abs(filePath); err != nil {
			return
		}

		var variableValue string

		if variableValue, err = resolver.ResolveVariableFromScript(executable, filePath, variableName); err != nil {
			return
		}

		value = cty.StringVal(variableValue)

		return
	},
})
