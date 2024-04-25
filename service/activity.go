package service

import (
	"github.com/advanced-go/example-domain/activity"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
)

func activityHandler[E core.ErrorHandler](w http.ResponseWriter, r *http.Request) (status *core.Status) {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return core.NewStatus(http.StatusBadRequest)
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		var entries []activity.EntryV1
		entries, status = activity.GetEntry(r.Header, r.URL.Query())
		httpx.WriteResponse[E](w, []httpx.Attr{{httpx.ContentType, httpx.ContentTypeJson}}, status.HttpCode(), entries)
		return status
	default:
		_, status = activity.PostEntry[*http.Request](r.Header, r.Method, r.URL.Query(), r)
		httpx.WriteResponse[E](w, nil, status.HttpCode(), nil)
		return status
	}
}
