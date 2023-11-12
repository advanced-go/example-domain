package timeseries

import (
	"context"
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/io2"
	"github.com/go-ai-agent/core/json2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
)

var (
	httpLoc     = PkgUri + "/httpHandler"
	httpWrapper = log2.WrapHttp(newHttpHandler[runtime.LogError]())
)

func newHttpHandler[E runtime.ErrorHandler]() runtime.HttpHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) *runtime.Status {
		return httpHandler[E](ctx, w, r)
	}
}

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

		entries, status := doHandler[E](newCtx, r, nil)
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
		_, status = doHandler[E](newCtx, r, buf)
		http2.WriteResponse[E](w, nil, status, nil)
		return status
	case http.MethodDelete:
		_, status := doHandler[E](newCtx, r, nil)
		http2.WriteResponse[E](w, nil, status.SetRequestId(r), nil)
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewStatus(http.StatusMethodNotAllowed)
}

//if buf == nil {
//	nc := runtime.NewStatus(runtime.StatusInvalidContent)
//	http2.WriteResponse[E](w, nil, nc, nil)
//	return nc
//}
//status = json2.Unmarshal(buf, &entries)
//if !status.OK() {
//	e.Handle(status, requestId, httpLoc)
//} else {
//  addEntry(entries)
//}
