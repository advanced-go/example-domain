package slo

import (
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	getEntryFromPathLoc = PkgPath + ":getEntryFromPath"
	getEntryHandlerLoc  = PkgPath + ":getEntryHandler"
)

func getEntryHandler[E runtime.ErrorHandler](h http.Header, uri *url.URL) (t []Entry, status runtime.Status) {
	var e E
	ctx := runtime.NewFileUrlContext(nil, uri.String())

	t, status = queryEntries(ctx, uri)
	if !status.OK() && !status.NotFound() {
		e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
	}
	return
}
