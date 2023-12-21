package google

import (
	"github.com/advanced-go/core/uri"
	"net/url"
)

const (
	searchTag  = "search"
	searchPath = "/search"
)

var (
	r uri.Resolver
)

func init() {
	r = uri.NewResolver("https://www.google.com", defaultFunc)
}

func defaultFunc(id string) string {
	switch id {
	case searchTag:
		return searchPath
	}
	return id
}

func resolve(id string, values url.Values) string {
	return r.Resolve(id, values)
}

func setOverride(t any, host string) {
	r.SetOverride(t, host)
}
