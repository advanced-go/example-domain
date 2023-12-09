package google

import (
	"strings"
)

type resolverFunc func(string) string

var (
	defaultOrigin = "http://localhost:8080"
	list          []resolverFunc
)

func addResolver(fn resolverFunc) {
	if fn == nil {
		return
	}
	// do not need mutex, as this is only called from test
	list = append(list, fn)
}

// resolve - resolve a string to an url.
func resolve(s string) string {
	if list != nil {
		for _, r := range list {
			url := r(s)
			if len(url) != 0 {
				return url
			}
		}
	}
	return defaultResolver(s)
}

func defaultResolver(u string) string {
	// if an endpoint, then default to defaultOrigin
	if strings.HasPrefix(u, "/") {
		return defaultOrigin + u
	}
	// else pass through
	return u
}
