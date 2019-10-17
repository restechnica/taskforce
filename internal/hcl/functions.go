package hcl

import (
	"github.com/restechnica/taskforce/internal/interpolation"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

var ResolveVariableFromFile = function.New(&function.Spec{
	Params: []function.Parameter{
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
		var filePath = args[0].AsString()
		var variableName = args[1].AsString()

		var variableValue string

		if variableValue, err = interpolation.ResolveVariableFromFile(filePath, variableName); err != nil {
			return
		}

		value = cty.StringVal(variableValue)

		return
	},
})
