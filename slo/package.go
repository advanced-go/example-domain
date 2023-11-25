package slo

import (
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"time"
)

type pkg struct{}

const (
	ContentLocation = "Content-Location"
	PkgPath         = "github.com/advanced-go/example-domain/slo"
	Pattern         = "/" + PkgPath + "/"

	postEntryLoc = PkgPath + "/PostEntry"
	getEntryLoc  = PkgPath + "/GetEntry"
)

// GetEntryConstraints - Get constraints
type GetEntryConstraints interface {
	[]Entry | []byte
}

// GetEntry - generic get function with context and uri for resource selection and filtering
func GetEntry[T GetEntryConstraints](h http.Header, uri string) (t T, status runtime.Status) {
	var e runtime.LogError

	u, err := url.Parse(uri)
	if err != nil {
		status = runtime.NewStatusError(runtime.StatusInvalidContent, getEntryLoc, err)
		e.Handle(status, runtime.RequestId(h), "")
		return
	}
	if h == nil {
		h = make(http.Header)
	}
	http2.AddRequestIdHeader(h)
	defer access.LogDeferred(h, "GET", uri, access.NewStatusCodeClosure(&status))()
	t, status = getEntryHandler[T](nil, h, u)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), getEntryLoc)
	}
	return
}

// PostEntryConstraints - Post constraints
type PostEntryConstraints interface {
	[]Entry | []byte | runtime.Nillable
}

// PostEntry - exchange function
func PostEntry[T PostEntryConstraints](h http.Header, method, uri, variant string, body T) (t any, status runtime.Status) {
	var e runtime.LogError
	var r *http.Request

	r, status = http2.NewRequest(h, method, uri, variant, nil)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), postEntryLoc)
		return nil, status
	}
	http2.AddRequestIdHeader(h)
	defer access.LogDeferred(h, method, uri, access.NewStatusCodeClosure(&status))()
	t, status = postEntryHandler(nil, r, body)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), postEntryLoc)
	}
	return
}

// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	http2.AddRequestId(r)
	func() (status runtime.Status) {
		defer access.LogDeferred(r.Header, r.Method, r.URL.String(), access.NewStatusCodeClosure(&status))()
		return httpHandler[runtime.LogError](nil, w, r)
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
