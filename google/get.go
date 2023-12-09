package google

import (
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

func getHandler[E runtime.ErrorHandler](r *http.Request) (any, runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	requestId := "invalid-change"
	http2.AddRequestId(r)
	switch r.Method {
	case http.MethodGet:
		var e E

		req, err := http.NewRequest(http.MethodGet, resolve(searchUri(r.URL, googleEndpoint)), nil)
		if err != nil {
			return nil, e.Handle(runtime.NewStatusError(http.StatusInternalServerError, searchLocation, err), requestId, "")
		}
		// http2.Do() will always return a non nil *http.Response
		resp, status := http2.Do(req)
		if !status.OK() {
			return nil, e.Handle(status, requestId, searchLocation)
		}
		var buf []byte
		buf, status = io2.ReadAll(resp.Body)
		if !status.OK() {
			return nil, e.Handle(status, requestId, searchLocation)
		}
		status = runtime.NewStatusOK()
		status.ContentHeader().Set(http2.ContentType, resp.Header.Get(http2.ContentType))
		status.ContentHeader().Set(http2.ContentLength, fmt.Sprintf("%v", len(buf)))
		return buf, status
	}
	return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
}
