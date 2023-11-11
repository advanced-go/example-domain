package activity

import (
	"context"
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/io2"
	"github.com/go-ai-agent/core/json2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
)

var (
	httpLoc = PkgUri + "/httpHandler"
)

func httpHandler[E runtime.ErrorHandler](ctx context.Context, w http.ResponseWriter, r *http.Request) *runtime.Status {
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

/*
	if buf == nil {
		nc := runtime.NewStatus(runtime.StatusInvalidContent)
		http2.WriteResponse[E](w, nil, nc, nil)
		return nc
	}
	status = json2.Unmarshal(buf, &entries)
	if !status.OK() {
		e.Handle(status, requestId, loc)
	} else {
		addEntry(entries)
	}

*/

//requestId := runtime.GetOrCreateRequestId(r)
//if r.Header.Get(runtime.XRequestId) == "" {
//	r.Header.Set(runtime.XRequestId, requestId)
//}
// Handled in Http
// Need to create as new request as upstream calls may not be http, and rely on the context for a request id
//rc := r.Clone(runtime.NewRequestIdContext(r.Context(), requestId))

/*

//controller  = log2.NewController(newTypeHandler[runtime.LogError]())
// newTypeHandler - templated function providing a TypeHandlerFn
func newTypeHandler[E runtime.ErrorHandler]() runtime.TypeHandlerFn {
	return func(r *http.Request, body any) (any, *runtime.Status) {
		return doHandler[E](nil, r, body)
	}
}

*/
