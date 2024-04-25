package entryv2

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

const (
	getRouteName = "get"
)

func getHandler[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (t []Entry, status *core.Status) {
	var e E

	t, status = queryEntries(ctx, values)
	if !status.OK() && status.Code != http.StatusNotFound {
		e.Handle(status, core.RequestId(h))
	}
	if len(t) == 0 {
		return nil, core.NewStatus(http.StatusNotFound)
	}
	return
}
