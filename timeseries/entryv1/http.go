package entryv1

import (
	"github.com/advanced-go/core/http2"
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
		entries, status := getHandler[E](r.Context(), r.Header, r.URL.Query())
		if !status.OK() {
			httpx.WriteResponse[E](w, nil, status.HttpCode(), nil)
			return status
		}
		httpx.WriteResponse[E](w, []httpx.Attr{{http2.ContentType, http2.ContentTypeJson}}, status.HttpCode(), entries)
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
