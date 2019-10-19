package scripts

import (
	"os/exec"
)

func RunScript(executable string, path string) (output string, err error) {
	var process = exec.Command(executable, path)

	var combinedOutput []byte

	if combinedOutput, err = process.CombinedOutput(); err != nil {
		return
	}

	output = string(combinedOutput)

	return
}
