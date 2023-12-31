package entryv2

import (
	"context"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/example-domain/timeseries/types"
	"net/http"
	"net/url"
)

const (
	getHandlerLoc = PkgPath + ":getHandler"
)

func getHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (t []types.EntryV2, status runtime.Status) {
	var e E

	t, status = queryEntries(ctx, values)
	if !status.OK() && !status.NotFound() {
		e.Handle(status, runtime.RequestId(h), getHandlerLoc)
	}
	if len(t) == 0 {
		return nil, runtime.NewStatus(http.StatusNotFound)
	}
	return
}
