package timeseriesvar

import (
	"context"
	"errors"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
	"strings"
)

const (
	getEntryHandlerLoc  = PkgPath + "/getEntryHandler"
	getEntryLoc2        = PkgPath + "/getEntry"
	getEntryFromPathLoc = PkgPath + "/getEntryFromPath"
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
	//if !status.OK() {
	//	e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
	//}
	return t, status.AddLocation(getEntryHandlerLoc)
}

func getEntryFromPath[T GetEntryConstraints](location string) (t T, status runtime.Status) {
	buf, status2 := io2.ReadFileFromPath(location)
	if !status2.OK() {
		return t, status2.AddLocation(getEntryFromPathLoc)
	}
	v1 := strings.Index(location, "entry-v1") != -1
	v2 := strings.Index(location, "entry-v2") != -1
	if !v1 && !v2 {
		return t, runtime.NewStatus(runtime.StatusInvalidContent).AddLocation(getEntryFromPathLoc)
	}
	switch ptr := any(&t).(type) {
	case *[]EntryV1:
		status = json2.Unmarshal(buf, ptr)
		if !status.OK() {
			return t, status.AddLocation(getEntryFromPathLoc)
		}
		return t, runtime.NewStatusOK()
	case *[]EntryV2:
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
		entries := queryEntriesV1(u)
		if len(entries) == 0 {
			return nil, runtime.NewStatus(http.StatusNotFound)
		}
		*ptr = entries
		return t, runtime.NewStatusOK()
	case *[]EntryV2:
		entries := queryEntriesV2(u)
		if len(entries) == 0 {
			return nil, runtime.NewStatus(http.StatusNotFound)
		}
		*ptr = entries
		return t, runtime.NewStatusOK()
	case *[]byte:
		variant = verifyVariant(u, variant)
		if variant == EntryV1Variant {
			entries := queryEntriesV1(u)
			if len(entries) == 0 {
				return nil, runtime.NewStatus(http.StatusNotFound)
			}
			buf, status := json2.Marshal(entries)
			if !status.OK() {
				return nil, status.AddLocation(getEntryLoc2)
			}
			*ptr = buf
			return t, runtime.NewStatusOK()
		} else {
			entries := queryEntriesV2(u)
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
	default:
		return nil, runtime.NewStatus(runtime.StatusInvalidContent).AddLocation(getEntryLoc2)
	}
}
