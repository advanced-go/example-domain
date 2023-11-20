package timeseries

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
	getEntryHandlerLoc = PkgUri + "/getEntryHandler"
	getEntryLoc2       = PkgUri + "/getEntry"
	fromAnyLoc         = PkgUri + "/entryFromAny"
)

func getEntryHandler[T GetEntryConstraints](ctx context.Context, h http.Header, uri *url.URL) (t T, status runtime.Status) {
	if runtime.IsDebugEnvironment() {
		status2 := runtime.StatusFromContext(ctx)
		if status2 != nil {
			return t, status2.AddLocation(getEntryHandlerLoc)
		}
		location := h.Get(ContentLocation)
		if strings.HasPrefix(location, "file://") {
			t, status = getEntryFromLocation[T](location)
			return t, status.AddLocation(getEntryHandlerLoc)
		}
	}
	t, status = getEntry[T](uri, h.Get(ContentLocation))
	//if !status.OK() {
	//	e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
	//}
	return t, status.AddLocation(getEntryHandlerLoc)
}

func getEntryFromLocation[T GetEntryConstraints](location string) (t T, status runtime.Status) {
	buf, status2 := io2.ReadFileFromPath(location)
	if !status2.OK() {
		return t, status2
	}
	v1 := strings.Index(location, "entry-v1") != -1
	v2 := strings.Index(location, "entry-v2") != -1
	if !v1 && !v2 {
		return t, runtime.NewStatus(runtime.StatusInvalidContent)
	}
	switch ptr := any(&t).(type) {
	case *[]EntryV1:
		status = json2.Unmarshal(buf, ptr)
		if !status.OK() {
			return t, status
		}
		return t, runtime.NewStatusOK()
	case *[]EntryV2:
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
	/*
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
		case *[]EntryV2:
			if e, ok := a.([]EntryV2); ok {
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

	*/
}

type getEntryConstraints interface {
	[]EntryV1 | []EntryV2 | []byte
}

func getEntry[T getEntryConstraints](u *url.URL, variant string) (T, runtime.Status) {
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
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	}
}
