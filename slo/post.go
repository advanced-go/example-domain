package slo

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	postRouteName = "post-entry"
)

func postEntryHandler[E core.ErrorHandler](ctx context.Context, h http.Header, method string, _ url.Values, body any) (t any, status *core.Status) {
	var e E

	switch strings.ToUpper(method) {
	case http.MethodPut:
		var entries []EntryV1
		entries, status = createEntries(body)
		if !status.OK() {
			e.Handle(status, core.RequestId(h))
			return nil, status
		}
		if len(entries) == 0 {
			status = core.NewStatusError(core.StatusInvalidContent, errors.New("error: no entries found"))
			e.Handle(status, core.RequestId(h))
			return nil, status
		}
		status = addEntries(ctx, entries)
		if !status.OK() {
			e.Handle(status, core.RequestId(h))
		}
		return nil, status
	case http.MethodDelete:
		status = deleteEntries(ctx)
		if !status.OK() {
			e.Handle(status, core.RequestId(h))
		}
		return nil, status
	default:
		return nil, core.NewStatus(http.StatusMethodNotAllowed)
	}
}

func createEntries(body any) (entries []EntryV1, status *core.Status) {
	if body == nil {
		return nil, core.NewStatus(core.StatusInvalidContent).AddLocation()
	}

	switch ptr := body.(type) {
	case []EntryV1:
		entries = ptr
	case []byte:
		entries, status = json.New[[]EntryV1](ptr, nil)
		if !status.OK() {
			return nil, status.AddLocation()
		}
	case *http.Request:
		entries, status = json.New[[]EntryV1](ptr.Body, nil)
		if !status.OK() {
			return nil, status.AddLocation()
		}
	case io.ReadCloser:
		entries, status = json.New[[]EntryV1](ptr, nil)
		if !status.OK() {
			return nil, status.AddLocation()
		}
	default:
		return nil, core.NewStatusError(core.StatusInvalidContent, core.NewInvalidBodyTypeError(body))
	}
	return entries, core.StatusOK()
}
