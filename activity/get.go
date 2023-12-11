package activity

import (
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	getEntryHandlerLoc = PkgPath + ":getEntryHandler"
)

func getEntryHandler[E runtime.ErrorHandler](h http.Header, uri *url.URL) (t []Entry, status runtime.Status) {
	var e E
	ctx := runtime.NewFileUrlContext(nil, uri.String())

	t, status = queryEntries(ctx, uri)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
		return t, status
	}
	if len(t) == 0 {
		return t, runtime.NewStatus(http.StatusNotFound)
	}
	return
}
