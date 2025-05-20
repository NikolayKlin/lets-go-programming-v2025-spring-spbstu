package configuration

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(path string) (*Configuration, error) {
	data, runErr := os.ReadFile(path)

	if runErr != nil {
		return nil, runErr
	}

	var cfg0 Configuration
	runErr = yaml.Unmarshal(data, &cfg0)

	if runErr != nil {
		return nil, runErr
	}

	return &cfg0, nil
}
