package models

import (
	"net/http"

	"github.com/google/uuid"
)

type APIRequestStats struct {
	StorageBase
	Host          string `json:"host"`
	Method        string `json:"method"`
	URL           string `json:"url"`
	ContentLength int64  `json:"content_length"`
}

func ParseAPIRequestStatus(req *http.Request) *APIRequestStats {
	return &APIRequestStats{
		StorageBase: StorageBase{
			ID: uuid.New(),
		},
		Host:          req.Host,
		Method:        req.Method,
		URL:           req.URL.String(),
		ContentLength: req.ContentLength,
	}
}
