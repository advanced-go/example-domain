package activity

import (
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

type pkg struct{}

const (
	PkgUri         = "github.com/advanced-go/example-domain/activity"
	PkgPath        = "/advanced-go/example-domain/activity"
	Pattern        = "/advanced-go/example-domain/activity/"
	EntryV1Variant = "github.com/advanced-go/example-domain/activity/EntryV1"

	postEntryLoc = PkgUri + "/PostEntry"
)

// GetEntryConstraints - Get constraints
type GetEntryConstraints interface {
	[]EntryV1
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
	r, status := http2.NewRequest(h, method, uri, variant, nil)
	if !status.OK() {
		var e runtime.LogError
		e.Handle(status, runtime.RequestId(h), postEntryLoc)
		return nil, status
	}
	http2.AddRequestId(r)
	defer access.LogDeferred(h, method, uri, access.NewStatusCodeClosure(&status))()
	return postEntryHandler[runtime.LogError](nil, r, body)
}

// HttpHandler - Http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	http2.AddRequestId(r)
	func() (status runtime.Status) {
		defer access.LogDeferred(r.Header, r.Method, r.URL.String(), access.NewStatusCodeClosure(&status))()
		return httpHandler[runtime.LogError](nil, w, r)
	}()
}

type EntryV1 struct {
	//CreatedTS    time.Time
	ActivityID   string // Some form of UUID
	ActivityType string // trace|action
	Agent        string
	AgentUri     string // {host}:{agent}

	Assignment  string
	Controller  string
	Behavior    string
	Description string
}
