package Yaml

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func Read(configPath string, config *interface{}) error {

	pwd, _ := os.Getwd()
	realPath := filepath.Join(pwd, configPath)

	yfile, err := os.ReadFile(realPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yfile, &config)
	if err != nil {
		return err
	}

	return nil
}
