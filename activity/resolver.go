package activity

import (
	"github.com/advanced-go/core/runtime"
	"strings"
)

type resolverFunc func(string) string

var (
	defaultOrigin = "http://localhost:8080"
	resolverList  []resolverFunc
)

func addResolver(fn resolverFunc) {
	if !runtime.IsDebugEnvironment() || fn == nil {
		return
	}
	// do not need mutex, as this is only called from test
	resolverList = append(resolverList, fn)
}

// resolve - resolve a string to an url.
func resolve(s string) string {
	if !runtime.IsDebugEnvironment() {
		return defaultResolver(s)
	}
	if resolverList != nil {
		for _, r := range resolverList {
			url := r(s)
			if len(url) != 0 {
				return url
			}
		}
	}
	return defaultResolver(s)
}

func defaultResolver(uri string) string {
	// if an endpoint, then default to defaultOrigin
	if strings.HasPrefix(uri, "/") {
		return defaultOrigin + uri
	}
	// else pass through
	return uri
}
