package activity

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

const (
	getRouteName = "get-entry"
)

func getEntryHandler[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (t []EntryV1, status *core.Status) {
	var e E

	t, status = queryEntries(ctx, values)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return nil, status
	}
	if len(t) == 0 {
		return t, core.NewStatus(http.StatusNotFound)
	}
	return
}
