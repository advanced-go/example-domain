package entryv2

import (
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	PkgPath = "github/advanced-go/example-domain/timeseries/entryv2"
	//entryResource = "v2/entry"
)

// Get - get entries
func Get(h http.Header, values url.Values) (entries []Entry, status *runtime.Status) {
	h = runtime.AddRequestId(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, getLoc), getRouteName, "", -1, "", statusCode(&status))()
	return getHandler[runtime.Log](nil, h, values)
}

// PostConstraints - Post constraints
type PostConstraints interface {
	[]Entry | []byte | *http.Request | runtime.Nillable
}

// Post - exchange function for POST, PUT, DELETE...
func Post[T PostConstraints](h http.Header, method string, values url.Values, body T) (t any, status *runtime.Status) {
	h = runtime.AddRequestId(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, method, postLoc), postRouteName, "", -1, "", statusCode(&status))()
	return postHandler[runtime.Log](nil, h, method, values, body)
}
