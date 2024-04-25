package entryv2

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
)

func httpHandler[E core.ErrorHandler](w http.ResponseWriter, r *http.Request) *core.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return core.NewStatus(http.StatusBadRequest)
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		buf, status := getHandler[E](r.Context(), r.Header, r.URL.Query())
		if !status.OK() {
			httpx.WriteResponse[E](w, nil, status.HttpCode(), nil)
			return status
		}
		httpx.WriteResponse[E](w, []httpx.Attr{{httpx.ContentType, httpx.ContentTypeJson}}, status.HttpCode(), buf)
		return status
	case http.MethodPut:
		_, status := postHandler[E](r.Context(), r.Header, r.Method, r.URL.Query(), r.Body)
		httpx.WriteResponse[E](w, nil, status.HttpCode(), nil)
		return status
	case http.MethodDelete:
		_, status := postHandler[E](r.Context(), r.Header, r.Method, r.URL.Query(), nil)
		httpx.WriteResponse[E](w, nil, status.HttpCode(), nil)
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return core.NewStatus(http.StatusMethodNotAllowed)
}

//if buf == nil {
//	nc := core.NewStatus(core.StatusInvalidContent)
//	http2.WriteResponse[E](w, nil, nc, nil)
//	return nc
//}
//status = json2.Unmarshal(buf, &entries)
//if !status.OK() {
//	e.Handle(status, requestId, httpLoc)
//} else {
//  addEntry(entries)
//}
