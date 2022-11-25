package config

import (
	"errors"
	"os"
	"path"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Feed struct {
	Uri   string `yaml:"uri"`
	Label string `yaml:"label"`
}

type Config struct {
	Feeds []Feed `yaml:"feeds"`
}

func Write(name string, data []byte) {
	dir := getConfigDir()

	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		cobra.CheckErr(err)
	}

	err := os.WriteFile(name, data, 0644)
	cobra.CheckErr(err)
}

func Read() Config {
	cfg := Config{Feeds: []Feed{}}
	name := path.Join(getConfigDir(), "config.yml")
	data, err := os.ReadFile(name)

	// If the file doesn't exist, don't error as we just use the defaults
	if errors.Is(err, os.ErrNotExist) {
		return cfg
	} else {
		cobra.CheckErr(err)
	}

	err = yaml.Unmarshal(data, &cfg)
	cobra.CheckErr(err)

	return cfg
}

func getConfigDir() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return path.Join(home, ".config", "gobble")
}
