package activity

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
)

const (
	entryResource = "entry"
)

func HttpEntryHandler[E core.ErrorHandler](w http.ResponseWriter, r *http.Request) (status *core.Status) {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return core.NewStatus(http.StatusBadRequest)
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		var buf []EntryV1
		buf, status = getEntryHandler[E](r.Context(), r.Header, r.URL.Query())
		if !status.OK() {
			httpx.WriteResponse[E](w, nil, status.HttpCode(), nil)
			return status
		}
		httpx.WriteResponse[E](w, []httpx.Attr{{httpx.ContentType, httpx.ContentTypeJson}}, status.HttpCode(), buf)
		return status
	case http.MethodPut:
		_, status = postEntryHandler[E](r.Context(), r.Header, r.Method, r.URL.Query(), r.Body)
		httpx.WriteResponse[E](w, nil, status.HttpCode(), nil)
		return status
	case http.MethodDelete:
		_, status = postEntryHandler[E](r.Context(), r.Header, r.Method, r.URL.Query(), nil)
		httpx.WriteResponse[E](w, nil, status.HttpCode(), nil)
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return core.NewStatus(http.StatusMethodNotAllowed)
}
