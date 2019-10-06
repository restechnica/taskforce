package hcl

import "github.com/restechnica/taskforce/internal/config"

func LoadHCL(filePath string) (config.Root, error) {
	var configuration config.Root
	var err error

	if configuration, err = Parse(filePath); err != nil {
		return configuration, err
	}

	return configuration, nil
}
