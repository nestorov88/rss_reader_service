package service

import (
	rdrpkg "github.com/nestorov88/rss_reader/pkg/reader"
	"rss_reader_service/pkg/reader"
)

// rssReaderService implements Parsing for reader.RssReaderService
type rssReaderService struct {
}

// NewRssReaderService return instance of rssReaderService
func NewRssReaderService() reader.RssReaderService {
	return &rssReaderService{}
}

// Parse is parsing rss feed urls and return array of extracted RssItems
func (r rssReaderService) Parse(urls []string) ([]rdrpkg.RssItem, error) {
	return rdrpkg.Parse(urls)
}
