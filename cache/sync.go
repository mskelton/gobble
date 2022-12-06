package cache

import (
	"time"

	"github.com/mskelton/gobble/config"
	"github.com/mskelton/gobble/rss"
)

type CachedFeedItem struct {
	Id          string    `json:"id"`
	Uri         string    `json:"uri"`
	Title       string    `json:"title"`
	PublishedAt time.Time `json:"publishedAt"`
	Categories  []string  `json:"categories"`
}

type CachedFeed struct {
	Id    string           `json:"id"`
	Title string           `json:"title"`
	Items []CachedFeedItem `json:"items"`
}

type Cache struct {
	Feeds       []CachedFeed `json:"feeds"`
	LastUpdated time.Time    `json:"lastUpdated"`
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

	err := writeState(c)
	return c, err
}
