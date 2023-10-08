package models

import (
	"net/url"
)

type FeedSources struct {
	Endpoints []string
	Tags      string
}

type DomainAggregation struct {
	Domain       string `json:"domain"`
	ArticleCount string `json:"article_count"`
}

type Source struct {
	StorageBase
	URL     string `json:"url"`
	Domain  string `json:"domain"`
	Tags    string `json:"tags" gorm:"type:text"`
	Enabled bool   `json:"enabled"`
}

func (s *Source) Valid() bool {
	u, err := url.Parse(s.URL)
	if err != nil {
		return false
	}
	if u.Host == "" || u.Scheme == "" {
		return false
	}
	return true
}
