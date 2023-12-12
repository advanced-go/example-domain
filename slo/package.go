package slo

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type pkg struct{}

const (
	ContentLocation = "Content-Location"
	PkgPath         = "github.com/advanced-go/example-domain/slo"
	Pattern         = "/" + PkgPath + "/"

	entryResource = "entry"
	postEntryLoc  = PkgPath + ":PostEntry"
	getEntryLoc   = PkgPath + ":GetEntry"
)

// GetEntry - get entries
func GetEntry(h http.Header, uri string) (entries []Entry, status runtime.Status) {
	return getEntry[runtime.Log](h, uri)
}

func getEntry[E runtime.ErrorHandler](h http.Header, uri string) (entries []Entry, status runtime.Status) {
	u, err := url.Parse(uri)
	if err != nil {
		status = runtime.NewStatusError(runtime.StatusInvalidContent, getEntryLoc, err)
		return
	}
	h = http2.AddRequestIdHeader(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, getEntryLoc), "", -1, "", access.NewStatusCodeClosure(&status))()
	return getEntryHandler[runtime.Log](h, u)
}

// PostEntryConstraints - Post constraints
type PostEntryConstraints interface {
	[]Entry | []byte | runtime.Nillable
}

// PostEntry - exchange function
func PostEntry[T PostEntryConstraints](h http.Header, method, uri string, body T) (t any, status runtime.Status) {
	return postEntry[runtime.Log](h, method, uri, body)
}

func postEntry[E runtime.ErrorHandler, T PostEntryConstraints](h http.Header, method, uri string, body T) (t any, status runtime.Status) {
	var r *http.Request

	r, status = http2.NewRequest(h, method, uri, nil)
	if !status.OK() {
		return nil, status
	}
	http2.AddRequestIdHeader(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, method, postEntryLoc), "", -1, "", access.NewStatusCodeClosure(&status))()
	return postEntryHandler[runtime.Log](r, body)
}

// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	_, rsc, ok := http2.UprootUrn(r.URL.Path)
	if !ok || len(rsc) == 0 {
		status := runtime.NewStatusWithContent(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid path, not a valid URN: %v", r.URL.Path)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
		return
	}
	switch strings.ToLower(rsc) {
	case entryResource:
		httpHandler[runtime.Log](w, r)
	default:
		status := runtime.NewStatusWithContent(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: %v", rsc)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
	}
}

func httpHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) runtime.Status {
	if r == nil {
		return runtime.NewStatus(runtime.StatusInvalidArgument)
	}
	http2.AddRequestId(r)
	return func() (status runtime.Status) {
		defer access.LogDeferred(access.InternalTraffic, r, "", -1, "", access.NewStatusCodeClosure(&status))()
		return httpEntryHandler[E](w, r)
	}()
}

type Entry struct {
	CreatedTS time.Time
	Id        string
	// What does this apply to
	Controller string

	// Types of SLOs
	// percentage of traffic : 10% or 10
	// latency percentile: 99/500ms
	Threshold   string // Either percentage of traffic, or latency percentile
	StatusCodes string // For percentage
}
