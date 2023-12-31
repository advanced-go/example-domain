package activity

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/core/uri"
	"github.com/advanced-go/example-domain/activity/types"
	"net/http"
	"net/url"
	"strings"
)

type pkg struct{}

const (
	PkgPath = "github.com/advanced-go/example-domain/activity"
	Pattern = "/" + PkgPath + "/"

	entryResource        = "entry"
	httpHandlerRouteName = "http-handler"
	postRouteName        = "post-entry"
	postEntryLoc         = PkgPath + ":PostEntry"

	getRouteName = "get-entry"
	getEntryLoc  = PkgPath + ":GetEntry"
)

// GetEntry - get entries with headers and values
func GetEntry(h http.Header, values url.Values) (entries []types.Entry, status runtime.Status) {
	h = runtime.AddRequestId(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, getEntryLoc), getRouteName, "", -1, "", &status)()
	return getEntryHandler[runtime.Log](nil, h, values)
}

// PostEntryConstraints - Post constraints
type PostEntryConstraints interface {
	[]types.Entry | []byte | runtime.Nillable
}

// PostEntry - exchange function
func PostEntry[T PostEntryConstraints](h http.Header, method string, values url.Values, body T) (t any, status runtime.Status) {
	h = runtime.AddRequestId(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, method, postEntryLoc), postRouteName, "", -1, "", &status)()
	return postEntryHandler[runtime.Log](nil, h, method, values, body)
}

// HttpHandler - Http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		http2.WriteResponse[runtime.Log](w, nil, runtime.NewStatus(runtime.StatusInvalidArgument), nil)
		return
	}
	_, rsc, ok := uri.UprootUrn(r.URL.Path)
	if !ok || len(rsc) == 0 {
		status := runtime.NewStatusWithContent(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid path, not a valid URN: %v", r.URL.Path)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
		return
	}
	runtime.AddRequestId(r)
	switch strings.ToLower(rsc) {
	case entryResource:
		func() (status runtime.Status) {
			defer access.LogDeferred(access.InternalTraffic, r, httpHandlerRouteName, "", -1, "", &status)()
			return httpEntryHandler[runtime.Log](w, r)
		}()
	default:
		status := runtime.NewStatusWithContent(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: %v", rsc)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
	}
}

/*
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

*/
