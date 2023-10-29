package google

import (
	"fmt"
	"net/url"
)

func searchUri(u *url.URL, uri string) string {
	if u == nil {
		return uri
	}
	if u.Query().Get(googleQueryArgName) != "" {
		return uri + fmt.Sprintf("?%v=%v", googleQueryArgName, u.Query().Get(googleQueryArgName))
	}
	return uri
}
