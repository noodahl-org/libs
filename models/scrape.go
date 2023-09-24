package models

type ScrapeType int

const (
	SCRAPE_INELIGIBLE ScrapeType = iota
	SCRAPE_RSS_FEED
	SCRAPE_URL
)

type Scrape struct {
	StorageBase
	URL          string
	HTTPResponse int
	Type         ScrapeType
}
