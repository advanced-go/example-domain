package activity

import (
	"context"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	getEntryHandlerLoc = PkgPath + ":getEntryHandler"
	getRouteName       = "get-entry"
)

func getEntryHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (t []Entry, status runtime.Status) {
	var e E
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, getEntryHandlerLoc), getRouteName, "", -1, "", &status)()

	t, status = queryEntries(ctx, values)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
		return t, status
	}
	if len(t) == 0 {
		return t, runtime.NewStatus(http.StatusNotFound)
	}
	return
}
