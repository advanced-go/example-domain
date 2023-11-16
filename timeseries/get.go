package timeseries

import (
	"errors"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

type getEntryHandlerFn func(h http.Header, uri *url.URL) (any, *runtime.Status)

const (
	getEntryHandlerLoc = PkgUri + "/getEntryHandler"
	getEntryLoc2       = PkgUri + "/getEntry"
	fromAnyLoc         = PkgUri + "/entryFromAny"
)

func getEntryHandler[T GetEntryConstraints, E runtime.ErrorHandler](proxy getEntryHandlerFn, h http.Header, uri *url.URL) (t T, status *runtime.Status) {
	var e E

	if proxy != nil {
		a, status1 := proxy(h, uri)
		if !status1.OK() {
			e.Handle(status, runtime.RequestId(h), "")
			return t, status1
		}
		return entryFromAny[T](a)
	}
	t, status = getEntry[T](uri, "")
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
	}
	return
}

func entryFromAny[T GetEntryConstraints](a any) (t T, status *runtime.Status) {
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
	return t, runtime.NewStatusOK()
}

type getEntryConstraints interface {
	[]EntryV1 | []EntryV2 | []byte
}

func getEntry[T getEntryConstraints](u *url.URL, variant string) (T, *runtime.Status) {
	var t T

	switch ptr := any(&t).(type) {
	case *[]EntryV1:
		entries := queryEntriesV1(u)
		if len(entries) == 0 {
			return nil, runtime.NewStatus(http.StatusNotFound)
		}
		*ptr = entries
	case *[]EntryV2:
		entries := queryEntriesV2(u)
		if len(entries) == 0 {
			return nil, runtime.NewStatus(http.StatusNotFound)
		}
		*ptr = entries
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
		}
	default:
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	}
	return t, runtime.NewStatusOK()
}
