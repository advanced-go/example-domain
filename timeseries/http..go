package timeseries

import (
	"context"
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"strings"
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
	var e E

	statusVar := validateVariant(r)
	if !statusVar.OK() {
		e.Handle(statusVar, runtime.RequestId(r), httpLoc)
		http2.WriteResponse[E](w, nil, statusVar, nil)
		return statusVar
	}
	if runtime.IsDebugEnvironment() {
		if fn := http2.HttpHandlerProxy(ctx); fn != nil {
			return fn(ctx, w, r)
		}
	}
	var newCtx any
	if ctx != nil {
		newCtx = ctx
	} else {
		newCtx = r
	}
	http2.AddRequestId(r)
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		buf, status := getEntry[[]byte](ctx, r.URL, r.Header.Get(http2.ContentLocation))
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), httpLoc)
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		http2.WriteResponse[E](w, buf, status, []http2.Attr{{http2.ContentType, http2.ContentTypeJson}})
		return status
	case http.MethodPut:
		status := putEntry(newCtx, r.Header.Get(http2.ContentLocation), r.Body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), httpLoc)
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		http2.WriteResponse[E](w, nil, status, nil)
		return status
	case http.MethodDelete:
		status := deleteEntry(newCtx, r.Header.Get(http2.ContentLocation))
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), httpLoc)
		}
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
