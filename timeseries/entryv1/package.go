package entryv1

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

const (
	PkgPath = "github/advanced-go/example-domain/timeseries/entryv1"
	//entryResource = "v1/entry"
)

// Get - get entries
func Get(h http.Header, values url.Values) (entries []Entry, status *core.Status) {
	h = core.AddRequestId(h)
	//defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, getLoc), getRouteName, "", -1, "", access.StatusCode(&status))()
	return getHandler[core.Log](nil, h, values)
}

// PostConstraints - Post constraints
type PostConstraints interface {
	[]Entry | []byte | *http.Request
}

// Post - exchange function
func Post[T PostConstraints](h http.Header, method string, values url.Values, body T) (t any, status *core.Status) {
	h = core.AddRequestId(h)
	//defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, method, postLoc), postRouteName, "", -1, "", access.StatusCode(&status))()
	return postHandler[core.Log](nil, h, method, values, body)
}
