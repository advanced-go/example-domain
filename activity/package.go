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
	PkgPath = "github.com/advanced-go/example-domain/activity"
	Pattern = "/" + PkgPath + "/"

	entryResource = "entry"
	postRouteName = "post-entry"
	postEntryLoc  = PkgPath + ":PostEntry"

	getRouteName = "get-entry"
	getEntryLoc  = PkgPath + ":GetEntry"
)

// GetEntry - get entries with headers and uri
func GetEntry(h http.Header, values url.Values) (entries []Entry, status runtime.Status) {
	h = http2.AddRequestIdHeader(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, getEntryLoc), getRouteName, -1, "", access.NewStatusCodeClosure(&status))()
	return getEntryHandler[runtime.Log](h, values, "")
}

// PostEntryConstraints - Post constraints
type PostEntryConstraints interface {
	[]Entry | []byte | runtime.Nillable
}

// PostEntry - exchange function
func PostEntry[T PostEntryConstraints](h http.Header, method string, values url.Values, body T) (t any, status runtime.Status) {
	h = http2.AddRequestIdHeader(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, method, postEntryLoc), postRouteName, -1, "", access.NewStatusCodeClosure(&status))()
	return postEntryHandler[runtime.Log](h, method, values, body, "")
}

// HttpHandler - Http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		http2.WriteResponse[runtime.Log](w, nil, runtime.NewStatus(runtime.StatusInvalidArgument), nil)
		return
	}
	_, rsc, ok := http2.UprootUrn(r.URL.Path)
	if !ok || len(rsc) == 0 {
		status := runtime.NewStatusWithContent(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid path, not a valid URN: %v", r.URL.Path)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
		return
	}
	http2.AddRequestId(r)
	switch strings.ToLower(rsc) {
	case entryResource:
		func() (status runtime.Status) {
			defer access.LogDeferred(access.InternalTraffic, r, "HttpHandler", -1, "", access.NewStatusCodeClosure(&status))()
			return httpEntryHandler[runtime.Log](w, r, "")
		}()
	default:
		status := runtime.NewStatusWithContent(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: %v", rsc)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
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
