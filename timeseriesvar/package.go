package timeseriesvar

import (
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"time"
)

const (
	PkgPath        = "github.com/advanced-go/example-domain/timeseriesvar"
	Pattern        = "/" + PkgPath + "/"
	EntryV1Variant = "github.com/advanced-go/example-domain/timeseriesvar:EntryV1"
	EntryV2Variant = "github.com/advanced-go/example-domain/timeseriesvar:EntryV2"

	postEntryLoc = PkgPath + "/PostEntry"
	getEntryLoc  = PkgPath + "/GetEntry"
)

// GetEntryConstraints - Get constraints
type GetEntryConstraints interface {
	[]EntryV1 | []EntryV2 | []byte
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
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, uri), -1, "", access.NewStatusCodeClosure(&status))()
	t, status = getEntryHandler[T](nil, h, u)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), getEntryLoc)
	}
	return
}

// PostEntryConstraints - Post constraints
type PostEntryConstraints interface {
	[]EntryV1 | []byte | runtime.Nillable
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
	http2.AddRequestId(r)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, method, uri), -1, "", access.NewStatusCodeClosure(&status))()
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
		defer access.LogDeferred(access.InternalTraffic, r, -1, "", access.NewStatusCodeClosure(&status))()
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
