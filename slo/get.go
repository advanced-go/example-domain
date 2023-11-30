package slo

import (
	"context"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"strings"
)

const (
	getEntryLoc2        = PkgPath + ":getEntry"
	getEntryFromPathLoc = PkgPath + ":getEntryFromPath"
	getEntryHandlerLoc  = PkgPath + ":getEntryHandler"
)

type getEntryConstraints interface {
	[]Entry | []byte
}

func getEntryHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, uri *url.URL) (t []Entry, status runtime.Status) {
	var e E

	if runtime.IsDebugEnvironment() {
		status2 := runtime.StatusFromContext(ctx)
		if status2 != nil {
			e.Handle(status2, runtime.RequestId(h), getEntryHandlerLoc)
			return t, status2
		}
		location := h.Get(ContentLocation)
		if strings.HasPrefix(location, "file://") {
			t, status = getEntryFromPath(location)
			if !status.OK() {
				e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
			}
			return t, status
		}
	}
	t = queryEntries(uri)
	if len(t) == 0 {
		return nil, runtime.NewStatus(http.StatusNotFound)
	}
	return t, runtime.StatusOK()
}

func getEntryFromPath(location string) (t []Entry, status runtime.Status) {
	buf, status2 := io2.ReadFileFromPath(location)
	if !status2.OK() {
		return t, status2.AddLocation(getEntryFromPathLoc)
	}
	status = json2.Unmarshal(buf, &t)
	if !status.OK() {
		return t, status.AddLocation(getEntryFromPathLoc)
	}
	return t, runtime.NewStatusOK()
}
