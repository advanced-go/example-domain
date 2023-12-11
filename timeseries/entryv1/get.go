package entryv1

import (
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	getHandlerLoc = PkgPath + ":getHandler"
)

func getHandler[E runtime.ErrorHandler](h http.Header, uri *url.URL) (t []Entry, status runtime.Status) {
	var e E
	ctx := runtime.NewFileUrlContext(nil, uri.String())

	t, status = queryEntries(ctx, uri)
	if !status.OK() && !status.NotFound() {
		e.Handle(status, runtime.RequestId(h), getHandlerLoc)
	}
	if len(t) == 0 {
		return nil, runtime.NewStatus(http.StatusNotFound)
	}
	return
}
