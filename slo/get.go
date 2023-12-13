package slo

import (
	"context"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	getEntryHandlerLoc = PkgPath + ":getEntryHandler"
)

func getEntryHandler[E runtime.ErrorHandler](h http.Header, values url.Values, uri *url.URL) (t []Entry, status runtime.Status) {
	var e E
	var ctx context.Context

	if uri != nil {
		ctx = runtime.NewFileUrlContext(nil, uri.String())
	}
	t, status = queryEntries(ctx, values)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
		return nil, status
	}
	if len(t) == 0 {
		status = runtime.NewStatus(http.StatusNotFound)
	}
	return
}
