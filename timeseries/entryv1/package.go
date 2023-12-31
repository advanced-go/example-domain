package entryv1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/core/uri"
	"github.com/advanced-go/example-domain/timeseries/types"
	"net/http"
	"net/url"
	"strings"
)

const (
	PkgPath = "github.com/advanced-go/example-domain/timeseries/entryv1"

	entryResource        = "v1/entry"
	httpHandlerRouteName = "http-handler"
	getRouteName         = "get"
	getLoc               = PkgPath + ":Get"

	postRouteName = "post"
	postLoc       = PkgPath + ":Post"
)

// Get - get entries
func Get(h http.Header, values url.Values) (entries []types.EntryV1, status runtime.Status) {
	h = runtime.AddRequestId(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, getLoc), getRouteName, "", -1, "", &status)()
	return getHandler[runtime.Log](nil, h, values)
}

// PostConstraints - Post constraints
type PostConstraints interface {
	[]types.EntryV1 | []byte | runtime.Nillable
}

// Post - exchange function
func Post[T PostConstraints](h http.Header, method string, values url.Values, body T) (t any, status runtime.Status) {
	h = runtime.AddRequestId(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, method, postLoc), postRouteName, "", -1, "", &status)()
	return postHandler[runtime.Log](nil, h, method, values, body)
}

// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		http2.WriteResponse[runtime.Log](w, nil, runtime.NewStatus(runtime.StatusInvalidArgument), nil)
		return
	}
	_, rsc, ok := uri.UprootUrn(r.URL.Path)
	if !ok || len(rsc) == 0 {
		status := runtime.NewStatusWithContent(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid path, not a valid URN: %v", r.URL.Path)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
		return
	}
	runtime.AddRequestId(r)
	switch strings.ToLower(rsc) {
	case entryResource:
		func() (status runtime.Status) {
			defer access.LogDeferred(access.InternalTraffic, r, httpHandlerRouteName, "", -1, "", &status)()
			return httpHandler[runtime.Log](w, r)
		}()
	default:
		status := runtime.NewStatusWithContent(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: %v", rsc)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
	}
}
