package cache

import (
	"path"

	"github.com/mskelton/gobble/config"
)

// Writes the cache to the state file
func writeCache(cache Cache) error {
	name, err := getCacheFilename()
	if err != nil {
		return err
	}

	return config.WriteYaml(name, cache)
}

// Get's the filename of the cache file
func getCacheFilename() (string, error) {
	return config.GetFilename(path.Join(".local", "state", "gobble", "cache.yml"))
}
