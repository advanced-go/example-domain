package entryv2

import (
	"context"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/example-domain/timeseries/context2"
	"net/http"
	"net/url"
	"strings"
)

const (
	getHandlerLoc  = PkgPath + ":getHandler"
	getFromPathLoc = PkgPath + ":getFromPath"
)

func getHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, uri *url.URL) (t []Entry, status runtime.Status) {
	var e E

	if runtime.IsDebugEnvironment() {
		status2 := context2.StatusFromContext(ctx)
		if status2 != nil {
			e.Handle(status2, runtime.RequestId(h), getHandlerLoc)
			return t, status2
		}
		location := h.Get(ContentLocation)
		if strings.HasPrefix(location, "file://") {
			t, status = getFromPath(location)
			e.Handle(status2, runtime.RequestId(h), getHandlerLoc)
			return t, status
		}
	}
	t = queryEntries(uri)
	if len(t) == 0 {
		return nil, runtime.NewStatus(http.StatusNotFound)
	}
	return t, runtime.StatusOK()
}

func getFromPath(location string) (t []Entry, status runtime.Status) {
	buf, status2 := io2.ReadFileFromPath(location)
	if !status2.OK() {
		return t, status2.AddLocation(getFromPathLoc)
	}
	status = json2.Unmarshal(buf, &t)
	if !status.OK() {
		return t, status.AddLocation(getFromPathLoc)
	}
	return t, runtime.StatusOK()
}
