package config

import (
	"errors"
	"github.com/restechnica/taskforce/internal/hcl"
	"os"
	"path"
)

const (
	hclFileName = "taskforce.hcl"
)

func Load(directory string) (Root, error) {
	var configurationPath string
	var configuration Root
	var err error

	if configurationPath, err = locate(directory); err != nil {
		return configuration, err
	}

	if configuration, err = hcl.Parse(configurationPath); err != nil {
		return configuration, err
	}

	return configuration, nil
}

func locate(directory string) (string, error) {
	if configurationPath, err := locateHclConfiguration(directory); err == nil {
		return configurationPath, nil
	}

	return "", errors.New("Failed to locate configuration file")
}

func locateHclConfiguration(directory string) (string, error) {
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
