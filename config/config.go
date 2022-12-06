package config

import (
	"errors"
	"os"
	"path"

	"github.com/mskelton/gobble/utils"
	"gopkg.in/yaml.v3"
)

type Feed struct {
	Id  string `yaml:"id"`
	Uri string `yaml:"uri"`
}

type Config struct {
	Feeds []Feed `yaml:"feeds"`
}

func Write(data []byte) error {
	name, err := getConfigFilename()
	if err != nil {
		return err
	}

	return utils.WriteFile(name, data)
}

func Read() (Config, error) {
	cfg := Config{Feeds: []Feed{}}

	name, err := getConfigFilename()
	if err != nil {
		return cfg, err
	}

	data, err := os.ReadFile(name)

	// If the file doesn't exist, don't error as we just use the defaults
	if errors.Is(err, os.ErrNotExist) {
		return cfg, nil
	} else if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(data, &cfg)
	return cfg, err
}

func getConfigFilename() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".config", "gobble", "config.yml"), nil
}
