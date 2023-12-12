package entryv1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"strings"
)

const (
	PkgPath         = "github.com/advanced-go/example-domain/timeseries/entryv1"
	getLoc          = PkgPath + ":Get"
	postLoc         = PkgPath + ":Post"
	ContentLocation = "Content-Location"
)

// Get - get entries
func Get(h http.Header, uri string) (entries []Entry, status runtime.Status) {
	return get[runtime.Log](h, uri)
}

func get[E runtime.ErrorHandler](h http.Header, uri string) (entries []Entry, status runtime.Status) {
	u, err := url.Parse(uri)
	if err != nil {
		status = runtime.NewStatusError(runtime.StatusInvalidContent, getLoc, err)
		return
	}
	h = http2.AddRequestIdHeader(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, getLoc), "", -1, "", access.NewStatusCodeClosure(&status))()
	return getHandler[E](h, u)
}

// PostConstraints - Post constraints
type PostConstraints interface {
	[]Entry | []byte | runtime.Nillable
}

// Post - exchange function
func Post[T PostConstraints](h http.Header, method, uri string, body T) (t any, status runtime.Status) {
	return post[runtime.Log, T](h, method, uri, body)
}

func post[E runtime.ErrorHandler, T PostConstraints](h http.Header, method, uri string, body T) (t any, status runtime.Status) {
	var r *http.Request

	r, status = http2.NewRequest(h, method, uri, nil)
	if !status.OK() {
		return nil, status
	}
	http2.AddRequestId(r)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, method, postLoc), "", -1, "", access.NewStatusCodeClosure(&status))()
	return postHandler[runtime.Log](r, body)
}

// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	//if r != nil && len(r.Header.Get(ContentLocation)) > 0 {
	//	status := runtime.NewStatusError(http.StatusBadRequest, httpHandlerLoc, errors.New("error content location not supported"))
	//	http2.WriteResponse[runtime.Log](w, nil, status, nil)
	//	return
	//}
	_, rsc, ok := http2.UprootUrn(r.URL.Path)
	if !ok || len(rsc) == 0 {
		status := runtime.NewStatusWithContent(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid path, not a valid URN: %v", r.URL.Path)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
		return
	}
	switch strings.ToLower(rsc) {
	case "entry":
		httpHandler[runtime.Log](w, r)
	default:
		status := runtime.NewStatusWithContent(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: %v", rsc)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
	}
}

func httpHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) runtime.Status {
	http2.AddRequestId(r)
	return func() (status runtime.Status) {
		defer access.LogDeferred(access.InternalTraffic, r, "", -1, "", access.NewStatusCodeClosure(&status))()
		return httpHandler2[E](w, r)
	}()
}
