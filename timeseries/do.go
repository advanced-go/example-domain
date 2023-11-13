package timeseries

import (
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/io2"
	"github.com/go-ai-agent/core/json2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"io"
	"net/http"
	"net/url"
)

type pkg struct{}

var (
	postWrapper = log2.WrapPost(newPostHandler[runtime.LogError]())
	doLoc       = PkgUri + "/doHandler"
	putLoc      = PkgUri + "/put"
	getLoc      = PkgUri + "/get"
)

func newPostHandler[E runtime.ErrorHandler]() log2.PostHandler {
	return func(ctx any, r *http.Request, body any) (any, *runtime.Status) {
		return postHandler[E](ctx, r, body)
	}
}

func postHandler[E runtime.ErrorHandler](ctx any, r *http.Request, body any) (any, *runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	}
	if runtime.IsDebugEnvironment() {
		if fn := http2.PostHandlerProxy(ctx); fn != nil {
			return fn(ctx, r, body)
		}
	}
	var e E
	variant2 := r.Header.Get(http2.ContentLocation)
	if variant2 != EntryV1Variant || variant2 != EntryV2Variant {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	}
	switch r.Method {
	/*
		case http.MethodGet:
			switch variant {
			case EntryV2Variant:
				entries, status := getEntry[[]EntryV2](ctx, variant, r.URL)
				if !status.OK() {
					e.Handle(status, runtime.RequestId(r), doLoc)
					return nil, status
				}
				//if len(entries) == 0 {
				//	return nil, runtime.NewStatus(http.StatusNotFound)
				//}
				return entries, runtime.NewStatusOK()
			case EntryV1Variant:
				entries, status := getEntry[[]EntryV1](ctx, variant, r.URL)
				if !status.OK() {
					e.Handle(status, runtime.RequestId(r), doLoc)
					return nil, status
				}
				//entries := queryEntriesV1(r.URL)
				//if len(entries) == 0 {
				//	return nil, runtime.NewStatus(http.StatusNotFound)
				//}
				return entries, runtime.NewStatusOK()
			default:
			}
			return nil, runtime.NewStatus(runtime.StatusInvalidContent)

	*/
	case http.MethodPut:
		status := putEntry(nil, variant2, body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(ctx), doLoc)
			return nil, status
		}
		return nil, runtime.NewStatusOK()
	case http.MethodDelete:
		status := deleteEntry(ctx, variant2)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(ctx), doLoc)
		}
		return nil, status
	default:
	}
	return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
}

func getEntry[T GetConstraints](ctx, uri any, variant string) (T, *runtime.Status) {
	var t T
	var u *url.URL

	switch ptr := uri.(type) {
	case *url.URL:
		u = ptr
	case string:
		u2, err := url.Parse(ptr)
		if err != nil {
			u = u2
		}
	default:
		return t, runtime.NewStatus(runtime.StatusInvalidContent)
	}
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
		// How to allow for query arg determination??
		// Need to determine the variant for []byte type as that is coming from an Http request.
		variant = verifyVariant(ctx, variant)
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
				return status.AddLocation(putLoc)
			}
		case io.ReadCloser:
			buf, status := io2.ReadAll(ptr)
			if !status.OK() {
				return status.AddLocation(putLoc)
			}
			status = json2.Unmarshal(buf, &entries)
			if !status.OK() {
				return status.AddLocation(putLoc)
			}
		default:
			return runtime.NewStatusError(runtime.StatusInvalidContent, putLoc, runtime.NewInvalidBodyTypeError(body))
		}
		if len(entries) == 0 {
			return runtime.NewStatus(runtime.StatusInvalidContent)
		}
		addEntryV1(entries)
	case EntryV2Variant:
		var entries []EntryV2

		switch ptr := body.(type) {
		case []EntryV2:
			entries = ptr
		case []byte:
			status := json2.Unmarshal(ptr, &entries)
			if !status.OK() {
				return status.AddLocation(putLoc)
			}
		case io.ReadCloser:
			buf, status := io2.ReadAll(ptr)
			if !status.OK() {
				return status.AddLocation(putLoc)
			}
			status = json2.Unmarshal(buf, &entries)
			if !status.OK() {
				return status.AddLocation(putLoc)
			}
		default:
			return runtime.NewStatusError(runtime.StatusInvalidContent, putLoc, runtime.NewInvalidBodyTypeError(body))
		}
		if len(entries) == 0 {
			return runtime.NewStatus(runtime.StatusInvalidContent)
		}
		addEntryV2(entries)
	default:
		return runtime.NewStatus(runtime.StatusInvalidContent)
	}
	return runtime.NewStatusOK()
}

func deleteEntry(ctx any, variant string) *runtime.Status {
	switch variant {
	case EntryV1Variant:
		deleteEntriesV1()
	case EntryV2Variant:
		deleteEntriesV2()
	default:
		return runtime.NewStatus(runtime.StatusInvalidContent)
	}
	return runtime.NewStatusOK()
}

func verifyVariant(ctx any, variant string) string {
	if r, ok := ctx.(*http.Request); ok {
		v := r.URL.Query().Get("v")
		if len(v) > 0 {
			if v == "1" {
				return EntryV1Variant
			}
			if v == "2" {
				return EntryV2Variant
			}
		}
	}
	if len(variant) == 0 {
		return EntryV1Variant
	}
	return variant
}
