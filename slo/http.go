package slo

import (
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/io2"
	"github.com/go-ai-agent/core/json2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
)

func httpHandler[E runtime.ErrorHandler](ctx any, w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewStatus(http.StatusBadRequest)
	}
	var newCtx any
	if ctx != nil {
		newCtx = ctx
	} else {
		newCtx = r
	}
	http2.AddRequestId(r)
	switch r.Method {
	case http.MethodGet:
		var buf []byte

		entries, status := Do(newCtx, r.Method, r.URL.String(), r.Header.Get(http2.ContentLocation), nil)
		if !status.OK() {
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		buf, status = json2.Marshal(entries)
		if !status.OK() {
			var e E
			e.Handle(status, runtime.RequestId(r), httpLoc)
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		http2.WriteResponse[E](w, buf, status, []http2.Attr{{http2.ContentType, http2.ContentTypeJson}})
		return status
	case http.MethodPut:
		var e E

		buf, status := io2.ReadAll(r.Body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), httpLoc)
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		_, status = Do(newCtx, r.Method, r.URL.String(), r.Header.Get(http2.ContentLocation), buf)
		http2.WriteResponse[E](w, nil, status, nil)
		return status
	case http.MethodDelete:
		_, status := Do(newCtx, r.Method, r.URL.String(), "", nil)
		http2.WriteResponse[E](w, nil, status, nil)
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewStatus(http.StatusMethodNotAllowed)
}
