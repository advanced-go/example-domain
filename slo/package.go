package slo

import (
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

type pkg struct{}

const (
	PkgPath = "github/advanced-go/example-domain/slo"
	//entryResource = "entry"
	//Pattern = "/" + PkgPath + "/"

)

// GetEntry - get entries
func GetEntry(h http.Header, values url.Values) (entries []EntryV1, status *runtime.Status) {
	h = runtime.AddRequestId(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, getEntryLoc), getRouteName, "", -1, "", statusCode(&status))()
	return getEntryHandler[runtime.Log](nil, h, values)
}

// PostEntryConstraints - Post constraints
type PostEntryConstraints interface {
	[]EntryV1 | []byte | *http.Request | runtime.Nillable
}

// PostEntry - exchange function
func PostEntry[T PostEntryConstraints](h http.Header, method string, values url.Values, body T) (t any, status *runtime.Status) {
	h = runtime.AddRequestId(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, method, postEntryLoc), postRouteName, "", -1, "", statusCode(&status))()
	return postEntryHandler[runtime.Log](nil, h, method, values, body)
}
