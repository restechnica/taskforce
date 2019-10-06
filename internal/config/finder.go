package config

import (
	"errors"
	"os"
	"path"
)

const (
	hclFileName = "taskforce.hcl"
)

func Find(directory string) (string, error) {
	if configurationPath, err := findHclConfiguration(directory); err == nil {
		return configurationPath, nil
	}

	return "", errors.New("Failed to locate configuration file")
}

func findHclConfiguration(directory string) (string, error) {
	var configurationPath = path.Join(directory, hclFileName)

	if isFile(configurationPath) {
		return configurationPath, nil
	}

	return "", errors.New("Failed to locate HCL configuration file")

}

func isFile(file string) bool {
	var info, err = os.Stat(file)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
