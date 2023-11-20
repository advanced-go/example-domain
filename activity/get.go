package activity

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"strings"
)

const (
	getEntryHandlerLoc  = PkgUri + "/getEntryHandler"
	getEntryLoc2        = PkgUri + "/getEntry"
	getEntryFromPathLoc = PkgUri + "/getEntryFromPath"
)

func getEntryHandler[T GetEntryConstraints](ctx context.Context, h http.Header, uri *url.URL) (t T, status runtime.Status) {
	if runtime.IsDebugEnvironment() {
		status2 := runtime.StatusFromContext(ctx)
		if status2 != nil {
			return t, status2.AddLocation(getEntryHandlerLoc)
		}
		location := h.Get(ContentLocation)
		if strings.HasPrefix(location, "file://") {
			t, status = getEntryFromPath[T](location)
			return t, status.AddLocation(getEntryHandlerLoc)
		}
	}
	t, status = getEntry[T](uri, h.Get(ContentLocation))
	return t, status.AddLocation(getEntryHandlerLoc)
}

func getEntryFromPath[T GetEntryConstraints](location string) (t T, status runtime.Status) {
	buf, status2 := io2.ReadFileFromPath(location)
	if !status2.OK() {
		return t, status2.AddLocation(getEntryFromPathLoc)
	}
	v1 := strings.Index(location, "entry-v1")
	if v1 == -1 {
		return t, runtime.NewStatus(runtime.StatusInvalidContent)
	}
	switch ptr := any(&t).(type) {
	case *[]EntryV1:
		status = json2.Unmarshal(buf, ptr)
		if !status.OK() {
			return t, status.AddLocation(getEntryFromPathLoc)
		}
		return t, runtime.NewStatusOK()
	case *[]byte:
		*ptr = buf
		return t, runtime.NewStatusOK()
	default:
		return t, runtime.NewStatusError(runtime.StatusInvalidContent, getEntryFromPathLoc, errors.New("invalid type"))
	}
}

func getEntry[T GetEntryConstraints](u *url.URL, variant string) (T, runtime.Status) {
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
				return nil, status.AddLocation(getEntryLoc2)
			}
			*ptr = buf
			return t, runtime.NewStatusOK()
		}
		return nil, runtime.NewStatusError(runtime.StatusInvalidContent, getEntryLoc2, errors.New(fmt.Sprintf("invalid variant")))
	default:
		return nil, runtime.NewStatus(runtime.StatusInvalidContent).AddLocation(getEntryLoc2)
	}
}
