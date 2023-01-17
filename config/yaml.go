package config

import (
	"errors"
	"os"
	"path"

	"github.com/mskelton/gobble/utils"
	"gopkg.in/yaml.v3"
)

// Reads the YAML file and unmarshals it into the given struct
func ReadYaml[T any](name string, out T) error {
	buf, err := os.ReadFile(name)

	// If the file doesn't exist, don't error as we just use the defaults
	if errors.Is(err, os.ErrNotExist) {
		return nil
	} else if err != nil {
		return err
	}

	return yaml.Unmarshal(buf, &out)
}

// Writes the YAML file to the given filename
func WriteYaml(name string, data any) error {
	buf, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	return utils.WriteFile(name, buf)
}

// Returns the filename of the Gobble config file
func GetFilename(name string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, name), nil
}
