package entryv1

import (
	"context"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	getHandlerLoc = PkgPath + ":getHandler"
	getRouteName  = "get"
)

func getHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (t []entry, status runtime.Status) {
	var e E
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, getHandlerLoc), getRouteName, "", -1, "", &status)()

	t, status = queryEntries(ctx, values)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), getHandlerLoc)
		return nil, status
	}
	if len(t) == 0 {
		status = runtime.NewStatus(http.StatusNotFound)
	}
	return
}
