package cache

import (
	"time"

	"github.com/mskelton/gobble/config"
	"github.com/mskelton/gobble/rss"
)

type CachedFeedItem struct {
	Id          string    `yaml:"id"`
	Uri         string    `yaml:"uri"`
	Title       string    `yaml:"title"`
	PublishedAt time.Time `yaml:"publishedAt"`
	Categories  []string  `yaml:"categories"`
}

type CachedFeed struct {
	Id    string           `yaml:"id"`
	Title string           `yaml:"title"`
	Items []CachedFeedItem `yaml:"items"`
}

type RecentItem struct {
	Id    int    `yaml:"id"`
	Uri   string `yaml:"uri"`
	Title string `yaml:"title"`
}

type Cache struct {
	LastUpdated time.Time    `yaml:"lastUpdated"`
	Feeds       []CachedFeed `yaml:"feeds"`
	RecentItems []RecentItem `yaml:"recent-items"`
}

func Sync(cfg config.Config) (Cache, error) {
	c := Cache{
		LastUpdated: time.Now(),
		Feeds:       []CachedFeed{},
	}

	for _, source := range cfg.Feeds {
		rssFeed, err := rss.Read(source.Uri)
		if err != nil {
			return c, err
		}

		// Build the cached feed from the downloaded RSS feed
		cachedFeed := CachedFeed{
			Id:    source.Id,
			Title: rssFeed.Channel.Title,
			Items: []CachedFeedItem{},
		}

		// Populate the RSS items into the cache
		for _, rssItem := range rssFeed.Channel.Items {
			publishedAt, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", rssItem.PublishDate)
			if err != nil {
				return c, err
			}

			cachedFeed.Items = append(cachedFeed.Items, CachedFeedItem{
				Id:          rssItem.Id,
				Uri:         rssItem.Uri,
				Title:       rssItem.Title,
				PublishedAt: publishedAt,
				Categories:  rssItem.Categories,
			})
		}

		c.Feeds = append(c.Feeds, cachedFeed)
	}

	err := writeCache(c)
	return c, err
}
