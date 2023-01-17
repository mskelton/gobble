package config

import (
	"path"
)

type Feed struct {
	Id  string `yaml:"id"`
	Uri string `yaml:"uri"`
}

type Config struct {
	Feeds []Feed `yaml:"feeds"`
}

// Reads the Gobble config file
func ReadConfig() (Config, error) {
	cfg := Config{Feeds: []Feed{}}

	name, err := getConfigFilename()
	if err != nil {
		return cfg, err
	}

	err = ReadYaml(name, &cfg)
	return cfg, err
}

// Returns the filename of the Gobble config file
func getConfigFilename() (string, error) {
	return GetFilename(path.Join(".config", "gobble", "config.yml"))
}
