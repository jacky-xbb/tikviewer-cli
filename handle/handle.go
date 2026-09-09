// Package handle normalizes TikTok handles and profile URLs.
package handle

import (
	"errors"
	"regexp"
	"strings"
)

// ErrInvalid is returned when the input cannot be read as a TikTok handle.
var ErrInvalid = errors.New("not a TikTok handle or profile URL")

var valid = regexp.MustCompile(`^[A-Za-z0-9._]{2,24}$`)

// Normalize turns "@Name", "tiktok.com/@name?lang=en" or a bare "name"
// into the canonical lowercase handle without the leading "@".
func Normalize(in string) (string, error) {
	s := strings.TrimSpace(in)
	if s == "" {
		return "", ErrInvalid
	}
	if i := strings.Index(s, "tiktok.com/"); i >= 0 {
		s = s[i+len("tiktok.com/"):]
	}
	if i := strings.IndexAny(s, "?#"); i >= 0 {
		s = s[:i]
	}
	s = strings.Trim(s, "/")
	if i := strings.Index(s, "/"); i >= 0 {
		s = s[:i]
	}
	s = strings.TrimPrefix(s, "@")
	if !valid.MatchString(s) {
		return "", ErrInvalid
	}
	return strings.ToLower(s), nil
}

// StoryURL builds the TikViewer story-viewer URL for a handle.
func StoryURL(h string) string {
	return "https://tikviewer.org/@" + h
}
