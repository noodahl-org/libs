package models

import (
	"net/url"
)

type FeedSources struct {
	Sources []string
	Tags    []string
}

type Source struct {
	StorageBase
	URL    string `json:"url"`
	Domain string `json:"domain"`
	Tags   Tags   `json:"tags" gorm:"type:text"`
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
