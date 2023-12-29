package google

import (
	"fmt"
	"github.com/advanced-go/core/exchange"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

func getHandler[E runtime.ErrorHandler](r *http.Request) (any, runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	requestId := "invalid-change"
	runtime.AddRequestId(r)
	switch r.Method {
	case http.MethodGet:
		var e E
		newUrl := resolve(searchTag, r.URL.Query())
		req, err := http.NewRequest(http.MethodGet, newUrl, nil)
		if err != nil {
			return nil, e.Handle(runtime.NewStatusError(http.StatusInternalServerError, searchLocation, err), requestId, "")
		}
		// exchange.Do() will always return a non nil *http.Response
		resp, status := exchange.Do(req)
		if !status.OK() {
			return nil, e.Handle(status, requestId, searchLocation)
		}
		var buf []byte
		buf, status = http2.ReadAll(resp)
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
