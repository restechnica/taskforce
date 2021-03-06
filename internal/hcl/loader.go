package hcl

import "github.com/restechnica/taskforce/internal/config"

func LoadHCLFile(filePath string) (config.Root, error) {
	var configuration config.Root
	var err error

	if configuration, err = ParseHCLFile(filePath); err != nil {
		return configuration, err
	}

	return configuration, nil
}
