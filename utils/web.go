package utils

import (
	"io"
	"net/http"
	"net/url"
)

func ReadHTTPResponse(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return body, nil
}

func Domain(path string) string {
	u, err := url.Parse(path)
	if err != nil {
		return ""
	}
	return u.Host
}
