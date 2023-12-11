package entryv1

import (
	"context"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"io"
	"net/http"
	"strings"
)

const (
	postLoc2 = PkgPath + ":postHandler"
	putLoc   = PkgPath + ":put"
)

func postHandler[E runtime.ErrorHandler](r *http.Request, body any) (any, runtime.Status) {
	var e E

	if r == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	}
	ctx := runtime.NewFileUrlContext(nil, r.URL.String())
	switch strings.ToUpper(r.Method) {
	case http.MethodPut:
		status := put(ctx, body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), postLoc2)
		}
		return nil, status
	case http.MethodDelete:
		status := deleteEntries(ctx)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), postLoc2)
		}
		return nil, status
	default:
		return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
	}
}

func put(ctx context.Context, body any) runtime.Status {
	if body == nil {
		runtime.NewStatus(runtime.StatusInvalidContent).AddLocation(putLoc)
	}
	var entries []Entry

	switch ptr := body.(type) {
	case []Entry:
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
	return addEntry(ctx, entries)
}
