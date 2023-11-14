package activity

import (
	"errors"
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/io2"
	"github.com/go-ai-agent/core/json2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	postWrapper = log2.WrapPost(newPostEntryHandler[runtime.LogError]())
	postLoc     = PkgUri + "/postEntryHandler"
	getEntryLoc = PkgUri + "/getEntry"
	putEntryLoc = PkgUri + "/putEntry"
	fromAnyLoc  = PkgUri + "/entryFromAny"
)

// newPostEntryHandler - templated function providing a PostHandler
func newPostEntryHandler[E runtime.ErrorHandler]() runtime.PostHandler {
	return func(ctx any, r *http.Request, body any) (any, *runtime.Status) {
		return postEntryHandler[E](ctx, r, body)
	}
}

func postEntryHandler[E runtime.ErrorHandler](ctx any, r *http.Request, body any) (any, *runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	var e E

	statusVar := validateVariant(r)
	if !statusVar.OK() {
		e.Handle(statusVar, runtime.RequestId(r), httpLoc)
		return nil, statusVar
	}
	if runtime.IsDebugEnvironment() {
		if fn := http2.PostHandlerProxy(ctx); fn != nil {
			return fn(ctx, r, body)
		}
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodPut:
		status := putEntry(nil, r.Header.Get(http2.ContentLocation), body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(ctx), postLoc)
			return nil, status
		}
	case http.MethodDelete:
		status := deleteEntry(ctx, r.Header.Get(http2.ContentLocation))
		if !status.OK() {
			e.Handle(status, runtime.RequestId(ctx), postLoc)
		}
		deleteEntries()
		return nil, runtime.NewStatusOK()
	default:
	}
	return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
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
	[]EntryV1 | []byte
}

func getEntry[T getEntryConstraints](ctx any, u *url.URL, variant string) (T, *runtime.Status) {
	var t T

	switch ptr := any(&t).(type) {
	case *[]EntryV1:
		entries := queryEntries(u)
		if len(entries) == 0 {
			return nil, runtime.NewStatus(http.StatusNotFound)
		}
		*ptr = entries
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
		}
	default:
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	}
	return t, runtime.NewStatusOK()
}

func putEntry(ctx any, variant string, body any) *runtime.Status {
	if body == nil {
		runtime.NewStatus(runtime.StatusInvalidContent)
	}
	switch variant {
	case EntryV1Variant:
		var entries []EntryV1

		switch ptr := body.(type) {
		case []EntryV1:
			entries = ptr
		case []byte:
			status := json2.Unmarshal(ptr, &entries)
			if !status.OK() {
				return status.AddLocation(putEntryLoc)
			}
		case io.ReadCloser:
			buf, status := io2.ReadAll(ptr)
			if !status.OK() {
				return status.AddLocation(putEntryLoc)
			}
			status = json2.Unmarshal(buf, &entries)
			if !status.OK() {
				return status.AddLocation(putEntryLoc)
			}
		default:
			return runtime.NewStatusError(runtime.StatusInvalidContent, putEntryLoc, runtime.NewInvalidBodyTypeError(body))
		}
		if len(entries) == 0 {
			return runtime.NewStatus(runtime.StatusInvalidContent)
		}
		addEntry(entries)
	default:
		return runtime.NewStatus(runtime.StatusInvalidContent)
	}
	return runtime.NewStatusOK()
}

func deleteEntry(ctx any, variant string) *runtime.Status {
	switch variant {
	case EntryV1Variant:
		deleteEntries()
	default:
		return runtime.NewStatus(runtime.StatusInvalidContent)
	}
	return runtime.NewStatusOK()
}
