package activity

import (
	"context"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	getEntryHandlerLoc = PkgPath + ":getEntryHandler"
	getRouteName       = "get-entry"
	getEntryLoc        = PkgPath + ":GetEntry"
)

func getEntryHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (t []EntryV1, status runtime.Status) {
	var e E

	t, status = queryEntries(ctx, values)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
		return nil, status
	}
	if len(t) == 0 {
		return t, runtime.NewStatus(http.StatusNotFound)
	}
	return
}
