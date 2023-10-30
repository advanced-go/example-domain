package timeseries

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

// InConstraints - defining constraints for the TypeHandler
type InConstraints interface {
	[]EntryV1 | *struct{}
}

func TypeHandler[T InConstraints](r *http.Request, t T) (any, *runtime.Status) {
	return typeHandler[runtime.LogError, T](r, t)
}

func typeHandler[E runtime.ErrorHandler, T InConstraints](r *http.Request, content T) (any, *runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
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
			return nil, runtime.NewStatus(http.StatusNotFound)
		}
		return entries, runtime.NewStatusOK()
	case http.MethodPut:
		var entries []EntryV1

		switch ptr := any(content).(type) {
		case []EntryV1:
			entries = ptr
		default:
		}
		if len(entries) == 0 {
			return nil, runtime.NewStatus(runtime.StatusInvalidContent)
		}
		addEntry(entries)
		return nil, runtime.NewStatusOK()
	case http.MethodDelete:
		deleteEntries()
		return nil, runtime.NewStatusOK()
	default:
	}
	return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpHandler[runtime.LogError](w, r)
}

func httpHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewStatus(http.StatusBadRequest)
	}
	requestId := runtime.GetOrCreateRequestId(r)
	if r.Header.Get(runtime.XRequestId) == "" {
		r.Header.Set(runtime.XRequestId, requestId)
	}
	// Need to create as new request as upstream calls may not be http, and rely on the context for a request id
	rc := r.Clone(runtime.ContextWithRequestId(r.Context(), requestId))
	switch rc.Method {
	case http.MethodGet:
		entries, status := typeHandler[E, *struct{}](rc, nil)
		if !status.OK() {
			httpx.WriteMinResponse[E](w, status, nil)
			return status
		}
		buf, status1 := runtime.MarshalType(entries)
		if !status1.OK() {
			var e E
			e.HandleStatus(status1, requestId, loc)
			httpx.WriteMinResponse[E](w, status, nil)
			return status
		}
		httpx.WriteResponse[E](w, buf, status1, []httpx.Attr{{httpx.ContentType, httpx.ContentTypeJson}})
		return status
	case http.MethodPut:
		var entries []EntryV1
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
		entries, status = runtime.UnmarshalType[[]EntryV1](buf)
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
		httpx.WriteMinResponse[E](w, status.SetRequestId(requestId), nil)
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewStatus(http.StatusMethodNotAllowed)
}
