package google

import (
	uri2 "github.com/advanced-go/core/uri"
)

const (
	searchTag  = "search"
	searchPath = "/search"
)

var (
	resolver = uri2.NewResolver("https://www.google.com", defaultFunc)
)

func defaultFunc(id string) string {
	switch id {
	case searchTag:
		return searchPath
	}
	return id
}
