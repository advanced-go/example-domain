package entryv1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

const (
	getRouteName = "get"
	getLoc       = PkgPath + ":Get"
)

func getHandler[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (t []Entry, status *core.Status) {
	var e E

	t, status = queryEntries(ctx, values)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return nil, status
	}
	if len(t) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	}
	return
}
