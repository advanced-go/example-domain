package activity

import (
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
)

type pkg struct{}

var (
	EntryEndpoint = pkgPath + "/entry"

	PkgUri  = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath = runtime.PathFromUri(PkgUri)
	loc     = pkgPath + "/entryHandler"
)

// IsPkgStarted - returns status of startup
func IsPkgStarted() bool {
	return true
}

func DoHandler(req *http.Request) (*http.Response, error) {
	recorder := httpx.NewRecorder()
	status := entryHandler[runtime.BypassError](recorder, req)
	var err error
	if status.IsErrors() {
		err = status.FirstError()
	}
	return recorder.Result(), err
}

func EntryHandler(w http.ResponseWriter, r *http.Request) {
	entryHandler[runtime.LogError](w, r)
}
func entryHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewHttpStatus(http.StatusBadRequest)
	}
	requestId := runtime.GetOrCreateRequestId(r)
	if r.Header.Get(runtime.XRequestId) == "" {
		r.Header.Set(runtime.XRequestId, requestId)
	}
	// Need to create as new request as upstream calls may not be http, and rely on the context for a request id
	rc := r.Clone(runtime.ContextWithRequestId(r.Context(), requestId))
	switch rc.Method {
	case http.MethodGet:
		entries := queryEntries(rc.URL)
		if len(entries) == 0 {
			status := runtime.NewStatus(runtime.StatusNotFound)
			httpx.WriteMinResponse[E](w, status, nil)
			return status
		}
		buf, status := runtime.MarshalType(entries)
		if !status.OK() {
			var e E
			e.HandleStatus(status, requestId, loc)
			httpx.WriteMinResponse[E](w, status, nil)
			return status
		}
		httpx.WriteResponse[E](w, buf, status, []httpx.Attr{{httpx.ContentType, httpx.ContentTypeJson}})
		return status
	case http.MethodPut:
		var entries []entry
		var e E

		buf, status := httpx.ReadAll(rc.Body)
		if !status.OK() {
			e.HandleStatus(status, requestId, loc)
			httpx.WriteMinResponse[E](w, status, nil)
			return status
		}
		if buf == nil {
			nc := runtime.NewStatus(runtime.StatusInvalidContent)
			httpx.WriteMinResponse[E](w, nc, nil)
			return nc
		}
		entries, status = runtime.UnmarshalType[[]entry](buf)
		if !status.OK() {
			e.HandleStatus(status, requestId, loc)
		} else {
			addEntry(entries)
		}
		httpx.WriteMinResponse[E](w, status, nil)
		return status
	case http.MethodDelete:
		deleteEntries()
		status := runtime.NewStatusOK()
		httpx.WriteMinResponse[E](w, status, nil)
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewHttpStatus(http.StatusMethodNotAllowed)
}
