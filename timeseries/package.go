package timeseries

import (
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"time"
)

const (
	PkgUri         = "github.com/advanced-go/example-domain/timeseries"
	PkgPath        = "/advanced-go/example-domain/timeseries"
	Pattern        = "/advanced-go/example-domain/timeseries/"
	EntryV1Variant = "github.com/advanced-go/example-domain/timeseries/EntryV1"
	EntryV2Variant = "github.com/advanced-go/example-domain/timeseries/EntryV2"

	postEntryLoc = PkgUri + "/PostEntry"
	getEntryLoc  = PkgUri + "/GetEntry"
)

// GetEntryConstraints - Get constraints
type GetEntryConstraints interface {
	[]EntryV1 | []EntryV2 | []byte
}

// GetEntry - generic get function with context and uri for resource selection and filtering
func GetEntry[T GetEntryConstraints](h http.Header, uri string) (t T, status runtime.Status) {
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
	defer access.LogDeferred(h, "GET", uri, access.NewStatusCodeClosure(&status))()
	return getEntryHandler[T, runtime.LogError](nil, h, u)
}

// PostEntryConstraints - Post constraints
type PostEntryConstraints interface {
	[]EntryV1 | []byte | runtime.Nillable
}

// PostEntry - exchange function
func PostEntry[T PostEntryConstraints](h http.Header, method, uri, variant string, body T) (any, runtime.Status) {
	var e runtime.LogError

	r, status := http2.NewRequest(h, method, uri, variant, nil)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), postEntryLoc)
		return nil, status
	}
	http2.AddRequestId(r)
	defer access.LogDeferred(h, method, uri, access.NewStatusCodeClosure(&status))()
	return postEntryHandler[runtime.LogError](nil, r, body)
}

// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	http2.AddRequestId(r)
	func() (status runtime.Status) {
		defer access.LogDeferred(r.Header, r.Method, r.URL.String(), access.NewStatusCodeClosure(&status))()
		return httpHandler[runtime.LogError](nil, w, r)
	}()
}

type EntryV1 struct {
	CreatedTS time.Time
	Traffic   string
	Start     time.Time
	Duration  int

	RequestId string

	// Request attributes
	Url         string // {scheme}://{host}/{path} No query
	Protocol    string
	Host        string
	Path        string
	Method      string
	StatusCode  int32
	StatusFlags string

	Timeout   int32
	RateLimit float64
	RateBurst int32
}

type EntryV2 struct {
	CreatedTS time.Time
	Traffic   string
	Start     time.Time
	Duration  int

	RequestId string

	// Request attributes
	Url            string // {scheme}://{host}/{path} No query
	Protocol       string
	Host           string
	Path           string
	Method         string
	StatusCode     int32
	ThresholdFlags string
	Threshold      int
}
