package cache

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"time"

	"github.com/mskelton/gobble/config"
	"github.com/mskelton/gobble/utils"
)

func Read() (Cache, error) {
	cache := Cache{Feeds: []CachedFeed{}}

	name, err := getStateFilename()
	if err != nil {
		return cache, err
	}

	data, err := os.ReadFile(name)

	// If the file doesn't exist, don't error as we just use the defaults
	if errors.Is(err, os.ErrNotExist) {
		return cache, nil
	} else if err != nil {
		return cache, err
	}

	err = json.Unmarshal(data, &cache)
	return cache, err
}

func ReadS(cfg config.Config) (Cache, error) {
	c, err := Read()
	if err != nil {
		return c, err
	}

	// If the cache is fresh, return it without syncing
	if c.LastUpdated.Add(time.Hour * 4).After(time.Now()) {
		return c, nil
	}

	return Sync(cfg)
}

func getStateFilename() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".local", "state", "gobble", "state.json"), nil
}

func writeState(cache Cache) error {
	name, err := getStateFilename()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cache)
	if err != nil {
		return err
	}

	return utils.WriteFile(name, data)
}
