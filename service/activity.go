package service

import (
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/example-domain/activity"
	"net/http"
	"strings"
)

func activityHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) (status *runtime.Status) {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewStatus(http.StatusBadRequest)
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		var entries []activity.EntryV1
		entries, status = activity.GetEntry(r.Header, r.URL.Query())
		http2.WriteResponse[E](w, entries, status, []http2.Attr{{http2.ContentType, http2.ContentTypeJson}})
		return status
	default:
		_, status = activity.PostEntry[*http.Request](r.Header, r.Method, r.URL.Query(), r)
		http2.WriteResponse[E](w, nil, status, nil)
		return status
	}
}
