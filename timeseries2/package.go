package timeseries2

import (
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
	PkgPath = "github.com/advanced-go/example-domain/timeseries2"
	Pattern = "/" + PkgPath + "/"

	entryResource   = "entry"
	postEntryLoc    = PkgPath + "/PostEntry"
	getEntryLoc     = PkgPath + "/GetEntry"
	ContentLocation = "Content-Location"
)

// GetEntry - get entries
func GetEntry(h http.Header, uri string) (entries []Entry, status runtime.Status) {
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
	entries, status = getEntryHandler[[]Entry](nil, h, u)
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
	http2.AddRequestId(r)
	defer access.LogDeferred(h, method, uri, access.NewStatusCodeClosure(&status))()
	t, status = postEntryHandler(nil, r, body)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), postEntryLoc)
	}
	return
}

// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	_, rsc, ok := http2.UprootUrn(r.URL.Path)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http2.AddRequestId(r)
	switch strings.ToLower(rsc) {
	case entryResource:
		u := r.URL.String()
		if len(r.URL.Host) == 0 {
			u = "http://" + r.Host + r.URL.Path
		}
		func() (status runtime.Status) {
			defer access.LogDeferred(r.Header, r.Method, u, access.NewStatusCodeClosure(&status))()
			return httpEntryHandler[runtime.LogError](nil, w, r)
		}()
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

type Entry struct {
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
