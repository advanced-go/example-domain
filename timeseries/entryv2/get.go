package entryv2

import (
	"context"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	getHandlerLoc = PkgPath + ":getHandler"
	getRouteName  = "get"
	getLoc        = PkgPath + ":Get"
)

func getHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (t []Entry, status *runtime.Status) {
	var e E

	t, status = queryEntries(ctx, values)
	if !status.OK() && status.Code != http.StatusNotFound {
		e.Handle(status, runtime.RequestId(h), getHandlerLoc)
	}
	if len(t) == 0 {
		return nil, runtime.NewStatus(http.StatusNotFound)
	}
	return
}
