package entryv2

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	PkgPath         = "github.com/advanced-go/example-domain/timeseries/entryv2"
	getLoc          = PkgPath + ":Get"
	postLoc         = PkgPath + ":Post"
	ContentLocation = "Content-Location"
)

// Get - get entries
func Get(h http.Header, uri string) (entries []Entry, status runtime.Status) {
	var e runtime.LogError

	u, err := url.Parse(uri)
	if err != nil {
		status = runtime.NewStatusError(runtime.StatusInvalidContent, getLoc, err)
		e.Handle(status, runtime.RequestId(h), "")
		return
	}
	if h == nil {
		h = make(http.Header)
	}
	http2.AddRequestIdHeader(h)
	defer access.LogDeferred(access.IngressTraffic, access.NewRequest(h, http.MethodGet, uri), -1, "", access.NewStatusCodeClosure(&status))()
	entries, status = getHandler(nil, h, u)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), getLoc)
	}
	return
}

// PostConstraints - Post constraints
type PostConstraints interface {
	[]Entry | []byte | runtime.Nillable
}

// Post - exchange function
func Post[T PostConstraints](h http.Header, method, uri string, body T) (t any, status runtime.Status) {
	var e runtime.LogError
	var r *http.Request

	r, status = http2.NewRequest(h, method, uri, "", nil)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), postLoc)
		return nil, status
	}
	http2.AddRequestId(r)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, method, uri), -1, "", access.NewStatusCodeClosure(&status))()
	t, status = postHandler(nil, r, body)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), postLoc)
	}
	return
}

// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	_, _, ok := http2.UprootUrn(r.URL.Path)
	if !ok {
		status := runtime.NewStatus(http.StatusBadRequest)
		status.SetContent(errors.New(fmt.Sprintf("error invalid path, not a valid URN: %v", r.URL.Path)), false)
		http2.WriteResponse[runtime.LogError](w, nil, status, nil)
		return
	}
	http2.AddRequestId(r)
	func() (status runtime.Status) {
		defer access.LogDeferred(access.InternalTraffic, r, -1, "", access.NewStatusCodeClosure(&status))()
		return httpHandler[runtime.LogError](nil, w, r)
	}()

}
