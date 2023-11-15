package timeseries

import (
	"errors"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

var (
	getEntryHandlerLoc = PkgUri + "/getEntryHandler"
	getLoc             = PkgUri + "/getEntry"
	fromAnyLoc         = PkgUri + "/entryFromAny"
)

func getEntryHandler[T GetEntryConstraints, E runtime.ErrorHandler](ctx any, uri string) (t T, status *runtime.Status) {
	var e E

	u, err := url.Parse(uri)
	if err != nil {
		status = runtime.NewStatusError(runtime.StatusInvalidContent, getEntryHandlerLoc, err)
		e.Handle(status, runtime.RequestId(ctx), "")
		return
	}
	if runtime.IsDebugEnvironment() {
		if fn := http2.GetHandlerProxy(ctx); fn != nil {
			a, status1 := fn(ctx, uri, "")
			if !status1.OK() {
				e.Handle(status, runtime.RequestId(ctx), "")
				return t, status1
			}
			return entryFromAny[T](a)
		}
	}
	t, status = getEntry[T](ctx, u, "")
	if !status.OK() {
		e.Handle(status, runtime.RequestId(ctx), getEntryHandlerLoc)
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

func getEntry[T getEntryConstraints](ctx any, u *url.URL, variant string) (T, *runtime.Status) {
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
				return nil, status.AddLocation(getLoc)
			}
			*ptr = buf
		} else {
			entries := queryEntriesV2(u)
			if len(entries) == 0 {
				return nil, runtime.NewStatus(http.StatusNotFound)
			}
			buf, status := json2.Marshal(entries)
			if !status.OK() {
				return nil, status.AddLocation(getLoc)
			}
			*ptr = buf
		}
	default:
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	}
	return t, runtime.NewStatusOK()
}
