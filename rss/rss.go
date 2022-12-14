package rss

type RssItem struct {
	Id          string   `xml:"post-id"`
	Uri         string   `xml:"link"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Author      string   `xml:"dc:creator"`
	PublishDate string   `xml:"pubDate"`
	Categories  []string `xml:"category"`
}

type RssChannel struct {
	Title         string    `xml:"title"`
	Description   string    `xml:"description"`
	Link          string    `xml:"link"`
	LastBuildDate string    `xml:"lastBuildDate"`
	Items         []RssItem `xml:"item"`
}

type RssFeed struct {
	Channel RssChannel `xml:"channel"`
}

func Read(uri string) (RssFeed, error) {
	feed := RssFeed{}
	err := GetXml(uri, &feed)

	return feed, err
}
