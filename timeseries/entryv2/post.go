package entryv2

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
	postRouteName = "post"
)

func postHandler[E core.ErrorHandler](ctx context.Context, h http.Header, method string, _ url.Values, body any) (t any, status *core.Status) {
	var e E

	switch strings.ToUpper(method) {
	case http.MethodPut:
		var entries []Entry
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
		status = addEntry(ctx, entries)
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

func createEntries(body any) (entries []Entry, status *core.Status) {
	if body == nil {
		return nil, core.NewStatus(core.StatusInvalidContent).AddLocation()
	}

	switch ptr := body.(type) {
	case []Entry:
		entries = ptr
	case []byte:
		entries, status = json.New[[]Entry](ptr, nil)
		if !status.OK() {
			return nil, status.AddLocation()
		}
	case *http.Request:
		entries, status = json.New[[]Entry](ptr.Body, nil)
		if !status.OK() {
			return nil, status.AddLocation()
		}
	case io.ReadCloser:
		entries, status = json.New[[]Entry](ptr, nil)
		if !status.OK() {
			return nil, status.AddLocation()
		}
	default:
		return nil, core.NewStatusError(core.StatusInvalidContent, core.NewInvalidBodyTypeError(body))
	}
	return entries, core.StatusOK()
}
