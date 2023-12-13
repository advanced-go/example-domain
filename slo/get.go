package slo

import (
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

const (
	getEntryHandlerLoc = PkgPath + ":getEntryHandler"
)

func getEntryHandler[E runtime.ErrorHandler](h http.Header, values url.Values, variant string) (t []Entry, status runtime.Status) {
	var e E

	t, status = queryEntries(runtime.NewFileUrlContext(nil, variant), values)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
		return nil, status
	}
	if len(t) == 0 {
		status = runtime.NewStatus(http.StatusNotFound)
	}
	return
}
