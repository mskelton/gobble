package cache

import (
	"time"

	"github.com/mskelton/gobble/config"
)

func Read() (Cache, error) {
	cache := Cache{Feeds: []CachedFeed{}}

	name, err := getCacheFilename()
	if err != nil {
		return cache, err
	}

	err = config.ReadYaml(name, &cache)
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
