package entryv1

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

func getHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (t []types.EntryV1, status runtime.Status) {
	var e E

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
