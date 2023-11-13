package google

import (
	"fmt"
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/io2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
)

var (
	wrapper = log2.WrapDo(newDoHandler[runtime.LogError]())
)

func doHandler[E runtime.ErrorHandler](ctx any, r *http.Request, body any) (any, *runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	requestId := "invalid-change"
	http2.AddRequestId(r)
	switch r.Method {
	case http.MethodGet:
		var e E

		req, err := http.NewRequest(http.MethodGet, http2.Resolve(searchUri(r.URL, googleEndpoint)), nil)
		if err != nil {
			return nil, e.Handle(runtime.NewStatusError(http.StatusInternalServerError, searchLocation, err), requestId, "")
		}
		// exchange.Do() will always return a non nil *http.Response
		resp, status := http2.Do(req)
		if !status.OK() {
			return nil, e.Handle(status, requestId, searchLocation)
		}
		var buf []byte
		buf, status = io2.ReadAll(resp.Body)
		if !status.OK() {
			return nil, e.Handle(status, requestId, searchLocation)
		}
		status.Header().Set(http2.ContentType, resp.Header.Get(http2.ContentType))
		status.Header().Set(http2.ContentLength, fmt.Sprintf("%v", len(buf)))
		return buf, status
	}
	return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
}

// newDoHandler - templated function providing a DoHandler
func newDoHandler[E runtime.ErrorHandler]() log2.DoHandler {
	return func(ctx any, r *http.Request, body any) (any, *runtime.Status) {
		return doHandler[E](ctx, r, body)
	}
}
