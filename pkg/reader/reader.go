package reader

import "github.com/nestorov88/rss_reader/pkg/reader"

// RssReaderService is interface that wraps basic Parse method
// Parse is parsing RSS Feed urls and return array of reader.RssItem
// In case of error both the array and error is returned
type RssReaderService interface {
	Parse(urls []string) ([]reader.RssItem, error)
}
