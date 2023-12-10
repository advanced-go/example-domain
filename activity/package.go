package activity

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

type pkg struct{}

const (
	ContentLocation = "Content-Location"
	PkgPath         = "github.com/advanced-go/example-domain/activity"
	Pattern         = "/" + PkgPath + "/"

	entryResource  = "entry"
	postEntryLoc   = PkgPath + ":PostEntry"
	getEntryLoc    = PkgPath + ":GetEntry"
	httpHandlerLoc = PkgPath + ":HttpHandler"
)

// GetEntry - get entries with headers and uri
func GetEntry(h http.Header, uri string) (entries []Entry, status runtime.Status) {
	if h != nil && len(h.Get(ContentLocation)) > 0 {
		return nil, runtime.NewStatusError(http.StatusBadRequest, getEntryLoc, errors.New("error content location not supported"))
	}
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
	return getEntryHandler[E](h, u)
}

// PostEntryConstraints - Post constraints
type PostEntryConstraints interface {
	[]Entry | []byte | runtime.Nillable
}

// PostEntry - exchange function
func PostEntry[T PostEntryConstraints](h http.Header, method, uri string, body T) (t any, status runtime.Status) {
	if h != nil && len(h.Get(ContentLocation)) > 0 {
		return nil, runtime.NewStatusError(http.StatusBadRequest, postEntryLoc, errors.New("error: content location not supported"))
	}
	return postEntry[runtime.Log, T](h, method, uri, body)
}

func postEntry[E runtime.ErrorHandler, T PostEntryConstraints](h http.Header, method, uri string, body T) (t any, status runtime.Status) {
	var r *http.Request

	r, status = http2.NewRequest(h, method, uri, nil)
	if !status.OK() {
		return nil, status
	}
	http2.AddRequestId(r)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, method, postEntryLoc), "", -1, "", access.NewStatusCodeClosure(&status))()
	return postEntryHandler[E](r, body)
}

// HttpHandler - Http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r != nil && len(r.Header.Get(ContentLocation)) > 0 {
		status := runtime.NewStatusError(http.StatusBadRequest, httpHandlerLoc, errors.New("error content location not supported"))
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
		return
	}
	httpHandler[runtime.Log](w, r)
}

func httpHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) {
	_, rsc, ok := http2.UprootUrn(r.URL.Path)
	if !ok || len(rsc) == 0 {
		status := runtime.NewStatusWithContent(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid path, not a valid URN: %v", r.URL.Path)), false)
		http2.WriteResponse[E](w, nil, status, nil)
		return
	}
	http2.AddRequestId(r)
	switch strings.ToLower(rsc) {
	case entryResource:
		func() (status runtime.Status) {
			defer access.LogDeferred(access.InternalTraffic, r, "", -1, "", access.NewStatusCodeClosure(&status))()
			return httpEntryHandler[E](w, r)
		}()
	default:
		status := runtime.NewStatusWithContent(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: %v", rsc)), false)
		http2.WriteResponse[E](w, nil, status, nil)
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
