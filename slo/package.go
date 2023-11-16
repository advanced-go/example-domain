package slo

import (
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/log2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"time"
)

type pkg struct{}

const (
	PkgUri         = "github.com/advanced-go/example-domain/slo"
	PkgPath        = "/advanced-go/example-domain/slo"
	Pattern        = "/advanced-go/example-domain/slo/"
	EntryV1Variant = "github.com/advanced-go/example-domain/slo/EntryV1"

	postEntryLoc = PkgUri + "/PostEntry"
)

// GetEntryConstraints - Get constraints
type GetEntryConstraints interface {
	[]EntryV1
}

// GetEntry - generic get function with context and uri for resource selection and filtering
func GetEntry[T GetEntryConstraints](h http.Header, uri string) (t T, status *runtime.Status) {
	u, err := url.Parse(uri)
	if err != nil {
		var e runtime.LogError
		status = runtime.NewStatusError(runtime.StatusInvalidContent, getEntryLoc, err)
		e.Handle(status, runtime.RequestId(h), "")
		return
	}
	if h == nil {
		h = make(http.Header)
	}
	http2.AddRequestIdHeader(h)
	defer log2.Log(h, "GET", uri, log2.NewStatusCodeClosure(&status))()
	return getEntryHandler[T, runtime.LogError](nil, h, u)
}

// PostEntryConstraints - Post constraints
type PostEntryConstraints interface {
	[]EntryV1 | []byte | runtime.Nillable
}

// PostEntry - exchange function
func PostEntry[T PostEntryConstraints](h http.Header, method, uri, variant string, body T) (any, *runtime.Status) {
	var e runtime.LogError

	r, status := http2.NewRequest(h, method, uri, variant, nil)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), postEntryLoc)
		return nil, status
	}
	http2.AddRequestIdHeader(h)
	defer log2.Log(h, method, uri, log2.NewStatusCodeClosure(&status))()
	return postEntryHandler[runtime.LogError](nil, r, body)
}

// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	http2.AddRequestId(r)
	func() (status *runtime.Status) {
		defer log2.Log(r.Header, r.Method, r.URL.String(), log2.NewStatusCodeClosure(&status))()
		return httpHandler[runtime.LogError](nil, w, r)
	}()
}

type EntryV1 struct {
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
