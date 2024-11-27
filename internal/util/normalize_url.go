package util

import (
	"fmt"
	"net/url"
	"strings"
)

func NormalizeURL(urlStr string) (string, error) {
	parsed, err := url.Parse(urlStr)

	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	fullPath := parsed.Host + parsed.Path

	fullPath = strings.ToLower(fullPath)

	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil

}
