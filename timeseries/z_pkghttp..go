package timeseries

import (
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/json"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
)

var (
	Pattern = pkgPath + "/"
	httpLoc = pkgPath + "/httpHandler"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpHandler[runtime.LogError](w, r)
}

func httpHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewStatus(http.StatusBadRequest)
	}
	r = httpx.UpdateHeadersAndContext(r)
	switch r.Method {
	case http.MethodGet:
		var buf []byte

		entries, status := TypeHandler[runtime.Nillable](r, nil)
		if !status.OK() {
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		buf, status = json.Marshal(entries)
		if !status.OK() {
			var e E
			e.Handle(status, runtime.RequestId(r), httpLoc)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		httpx.WriteResponse[E](w, buf, status, []httpx.Attr{{httpx.ContentType, httpx.ContentTypeJson}})
		return status
	case http.MethodPut:
		var e E

		buf, status := httpx.ReadAll(r.Body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), httpLoc)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		_, status = TypeHandler[[]byte](r, buf)
		httpx.WriteResponse[E](w, nil, status, nil)
		return status
	case http.MethodDelete:
		_, status := TypeHandler[runtime.Nillable](r, nil)
		httpx.WriteResponse[E](w, nil, status.SetRequestId(r), nil)
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewStatus(http.StatusMethodNotAllowed)
}

//if buf == nil {
//	nc := runtime.NewStatus(runtime.StatusInvalidContent)
//	httpx.WriteResponse[E](w, nil, nc, nil)
//	return nc
//}
//status = json.Unmarshal(buf, &entries)
//if !status.OK() {
//	e.Handle(status, requestId, httpLoc)
//} else {
//  addEntry(entries)
//}
