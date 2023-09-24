package models

type ScrapeType int

const (
	SCRAPE_INELIGIBLE ScrapeType = iota
	SCRAPE_RSS_FEDD
	SCRAPE_URL
)

type Scrape struct {
	StorageBase
	URL  string
	Type ScrapeType
}
