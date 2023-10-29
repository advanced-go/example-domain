package google

import (
	"net/url"
)

func searchEndpoint(u *url.URL) string {
	if u == nil {
		return searchPath
	}
	if u.Query().Get(queryArgName) != "" {
		return "https://www.google.com/search?q=" + u.Query().Get(queryArgName)
	}
	return searchPath
}
