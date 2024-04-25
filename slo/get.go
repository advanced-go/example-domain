package slo

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

const (
	getRouteName = "get-entry"
	getEntryLoc  = PkgPath + ":GetEntry"
)

func getEntryHandler[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (t []EntryV1, status *core.Status) {
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
