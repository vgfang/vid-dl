package services

import (
	"net/url"
	"strings"
)

// this file needs to be updated for more site support
// return the string representation of the URLtype
func GetUrlType(rawurl string) string {
	parsedUrl, err := url.ParseRequestURI(rawurl)

	if err != nil || parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		return "invalid"
	}

	switch {
	case strings.Contains(rawurl, "youtube.com") || strings.Contains(rawurl, "youtu.be"):
		return "youtube"
	case strings.Contains(rawurl, "tiktok.com"):
		return "tiktok"
	case strings.Contains(rawurl, "twitter.com"):
		return "twitter"
	case strings.Contains(rawurl, "instagram.com"):
		return "instagram"
	default:
		return "unknown"
	}
}
