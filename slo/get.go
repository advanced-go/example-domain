package slo

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
	getEntryLoc        = PkgUri + "/getEntry"
	fromAnyLoc         = PkgUri + "/entryFromAny"
	getEntryHandlerLoc = PkgUri + "/getEntryHandler"
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

func getEntryFromLocation[T GetEntryConstraints](a any) (t T, status runtime.Status) {
	if a == nil {
		return
	}
	switch ptr := any(&t).(type) {
	case *[]EntryV1:
		if e, ok := a.([]EntryV1); ok {
			*ptr = e
		} else {
			return t, runtime.NewStatusError(runtime.StatusInvalidContent, fromAnyLoc, errors.New("T and any types do not match"))
		}
	case *[]byte:
		if b, ok := a.([]byte); ok {
			*ptr = b
		} else {
			return t, runtime.NewStatusError(runtime.StatusInvalidContent, fromAnyLoc, errors.New("T and any types do not match"))
		}
	default:
		return t, runtime.NewStatusError(runtime.StatusInvalidContent, fromAnyLoc, errors.New("invalid type"))
	}
	return t, runtime.NewStatusOK()
}

// getEntryConstraints - Get constraints
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
