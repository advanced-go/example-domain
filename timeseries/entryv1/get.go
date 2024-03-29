package entryv1

import (
	"context"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	getRouteName = "get"
	getLoc       = PkgPath + ":Get"
)

func getHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (t []Entry, status *runtime.Status) {
	var e E

	t, status = queryEntries(ctx, values)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h))
		return nil, status
	}
	if len(t) == 0 {
		status = runtime.NewStatus(http.StatusNotFound)
	}
	return
}
