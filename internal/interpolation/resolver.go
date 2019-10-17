package interpolation

import (
	"encoding/json"
	"fmt"
	"github.com/restechnica/taskforce/internal/extensions/osext"
	"os/exec"
	"path/filepath"
)

func ResolveVariableFromFile(filePath string, name string) (value string, err error) {
	filePath = osext.ExpandTilde(filePath)

	if filePath, err = filepath.Abs(filePath); err != nil {
		return
	}

	var extension = filepath.Ext(filePath)
	var process *exec.Cmd

	switch extension {
	case ".py":
		process = exec.Command("python", filePath)
	default:
		err = fmt.Errorf("file extension '%s' not supported", extension)
		return
	}

	var combinedOutput []byte

	if combinedOutput, err = process.CombinedOutput(); err != nil {
		return
	}

	var variables map[string]interface{}

	if err = json.Unmarshal(combinedOutput, &variables); err != nil {
		return
	}

	value = variables[name].(string)

	return
}
