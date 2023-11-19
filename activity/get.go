package activity

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"strings"
)

const (
	getEntryHandlerLoc = PkgUri + "/getEntryHandler"
	getEntryLoc        = PkgUri + "/getEntry"
	fromAnyLoc         = PkgUri + "/entryFromAny"
)

func getEntryHandler[T GetEntryConstraints, E runtime.ErrorHandler](ctx context.Context, h http.Header, uri *url.URL) (t T, status runtime.Status) {
	var e E

	if runtime.IsDebugEnvironment() {
		status2 := runtime.StatusFromContext(ctx)
		if status2 != nil {
			return t, status2
		}
		location := h.Get(http2.ContentLocation)
		if strings.HasPrefix(location, "file://") {
			return getEntryFromLocation[T](location)
		}
	}
	t, status = getEntry[T](uri, h.Get(http2.ContentLocation))
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
	}
	return
}

func getEntryFromLocation[T GetEntryConstraints](location string) (t T, status runtime.Status) {
	buf, status2 := http2.ReadContentFromLocation(location)
	if !status2.OK() {
		return t, status2
	}
	v1 := strings.Index(location, "entryv1")
	if v1 == -1 {
		return t, runtime.NewStatus(runtime.StatusInvalidContent)
	}
	switch ptr := any(&t).(type) {
	case *[]EntryV1:
		status = json2.Unmarshal(buf, ptr)
		if !status.OK() {
			return t, status
		}
		return t, runtime.NewStatusOK()
	case *[]byte:
		*ptr = buf
		return t, runtime.NewStatusOK()
	default:
		return t, runtime.NewStatusError(runtime.StatusInvalidContent, fromAnyLoc, errors.New("invalid type"))
	}
}

type getEntryConstraints interface {
	[]EntryV1 | []byte
}

func getEntry[T getEntryConstraints](u *url.URL, variant string) (T, runtime.Status) {
	var t T

	switch ptr := any(&t).(type) {
	case *[]EntryV1:
		entries := queryEntries(u)
		if len(entries) == 0 {
			return nil, runtime.NewStatus(http.StatusNotFound)
		}
		*ptr = entries
		return t, runtime.NewStatusOK()
	case *[]byte:
		if variant == EntryV1Variant {
			entries := queryEntries(u)
			if len(entries) == 0 {
				return nil, runtime.NewStatus(http.StatusNotFound)
			}
			buf, status := json2.Marshal(entries)
			if !status.OK() {
				return nil, status.AddLocation(getEntryLoc)
			}
			*ptr = buf
			return t, runtime.NewStatusOK()
		}
		return nil, runtime.NewStatusError(runtime.StatusInvalidContent, getEntryLoc, errors.New(fmt.Sprintf("invalid variant")))
	default:
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	}
}
