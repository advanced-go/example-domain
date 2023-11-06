package activity

import (
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/json"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
)

type pkg struct{}

var (
	Pattern = pkgPath + "/"

	PkgUri  = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath = runtime.PathFromUri(PkgUri)
	loc     = pkgPath + "/entryHandler"
)

// InConstraints - defining constraints for the TypeHandler
type InConstraints interface {
	[]EntryV1 | runtime.Nil
}

func TypeHandler[T InConstraints](r *http.Request, body T) (any, *runtime.Status) {
	return typeHandler[runtime.LogError, T](r, body)
}

func typeHandler[E runtime.ErrorHandler, T InConstraints](r *http.Request, body T) (any, *runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	requestId := runtime.GetOrCreateRequestId(r)
	if r.Header.Get(runtime.XRequestId) == "" {
		r.Header.Set(runtime.XRequestId, requestId)
	}
	// Need to create as new request as upstream calls may not be http, and rely on the context for a request id
	rc := r.Clone(runtime.NewRequestIdContext(r.Context(), requestId))
	switch rc.Method {
	case http.MethodGet:
		entries := queryEntries(rc.URL)
		if len(entries) == 0 {
			return nil, runtime.NewStatus(http.StatusNotFound)
		}
		// Set content-type
		return entries, runtime.NewStatusOK()
	case http.MethodPut:
		var entries []EntryV1

		switch ptr := any(body).(type) {
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
	rc := r.Clone(runtime.NewRequestIdContext(r.Context(), requestId))
	switch rc.Method {
	case http.MethodGet:
		var buf []byte

		entries, status := typeHandler[E, runtime.Nil](rc, nil)
		if !status.OK() {
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}

		buf, status = json.Marshal(entries)
		if !status.OK() {
			var e E
			e.Handle(status, requestId, loc)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		httpx.WriteResponse[E](w, buf, status, []httpx.Attr{{httpx.ContentType, httpx.ContentTypeJson}})
		return status
	case http.MethodPut:
		var entries []EntryV1
		var e E

		buf, status := httpx.ReadAll(rc.Body)
		if !status.OK() {
			e.Handle(status, requestId, loc)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		if buf == nil {
			nc := runtime.NewStatus(runtime.StatusInvalidContent)
			httpx.WriteResponse[E](w, nil, nc, nil)
			return nc
		}
		status = json.Unmarshal(buf, &entries)
		if !status.OK() {
			e.Handle(status, requestId, loc)
		} else {
			addEntry(entries)
		}
		httpx.WriteResponse[E](w, nil, status, nil)
		return status
	case http.MethodDelete:
		deleteEntries()
		status := runtime.NewStatusOK()
		httpx.WriteResponse[E](w, nil, status, nil)
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewStatus(http.StatusMethodNotAllowed)
}
