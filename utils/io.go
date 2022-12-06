package utils

import (
	"errors"
	"os"
	"path"
)

func WriteFile(name string, data []byte) error {
	dir := path.Dir(name)
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		return err
	}

	return os.WriteFile(name, data, 0644)
}
