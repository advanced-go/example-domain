package entryv1

import (
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"strings"
)

func httpHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewStatus(http.StatusBadRequest)
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		entries, status := getHandler[E](r.Context(), r.Header, r.URL.Query())
		if !status.OK() {
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		http2.WriteResponse[E](w, entries, status, []http2.Attr{{http2.ContentType, http2.ContentTypeJson}})
		return status
	case http.MethodPut:
		_, status := postHandler[E](r.Context(), r.Header, r.Method, r.URL.Query(), r.Body)
		http2.WriteResponse[E](w, nil, status, nil)
		return status
	case http.MethodDelete:
		_, status := postHandler[E](r.Context(), r.Header, r.Method, r.URL.Query(), nil)
		http2.WriteResponse[E](w, nil, status, nil)
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewStatus(http.StatusMethodNotAllowed)
}
