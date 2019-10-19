package variables

import (
	"fmt"
	"github.com/restechnica/taskforce/internal/scripts"
	"github.com/tidwall/gjson"
)

type Resolver struct {
	VariableCache JSONCache
}

func NewResolver() (resolver Resolver) {
	return Resolver{
		VariableCache: NewJSONCache(),
	}
}

func (resolver Resolver) ResolveVariableFromScript(executable string, filePath string, key string) (value string, err error) {
	var output string
	var hasOutput, hasVariable bool

	if value, hasVariable = resolver.VariableCache.GetVariable(filePath, key); !hasVariable {
		if output, hasOutput = resolver.VariableCache.GetJSON(filePath); !hasOutput {
			if output, err = scripts.RunScript(executable, filePath); err != nil {
				return
			}

			resolver.VariableCache.AddJSON(filePath, output)
		}

		var result = gjson.Get(output, key)
		value = result.String()

		if result.Exists() {
			resolver.VariableCache.AddVariable(filePath, key, value)
		} else {
			err = fmt.Errorf(`variable "%s" does not exist in script "%s"`, key, filePath)
		}
	}

	return
}
