package activity

import (
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"strings"
)

type pkg struct{}

const (
	ContentLocation = "Content-Location"
	PkgPath         = "github.com/advanced-go/example-domain/activity"
	Pattern         = "/" + PkgPath + "/"

	entryResource = "entry"
	postEntryLoc  = PkgPath + "/PostEntry"
	getEntryLoc   = PkgPath + "/GetEntry"
)

// GetEntry - get entries with headers and uri
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
func PostEntry[T PostEntryConstraints](h http.Header, method, uri string, body T) (t any, status runtime.Status) {
	var r *http.Request
	var e runtime.LogError

	r, status = http2.NewRequest(h, method, uri, "", nil)
	if !status.OK() {
		var e runtime.LogError
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

// HttpHandler - Http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	_, rsc, ok := http2.UprootUrn(r.URL.Path)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//fmt.Printf("url -> [scheme:%v] [host:%v] [path:%v] [req-host:%v] [req-uri:%v]\n", r.URL.Scheme, r.URL.Host, r.URL.Path, r.Host, r.RequestURI)
	//fmt.Printf("rsc -> %v\n", rsc)
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
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

type Entry struct {
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
