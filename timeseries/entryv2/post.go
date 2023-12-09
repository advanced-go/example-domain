package entryv2

import (
	"context"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/example-domain/timeseries/context2"
	"io"
	"net/http"
	"strings"
)

const (
	postLoc2 = PkgPath + ":postHandler"
	putLoc   = PkgPath + ":put"
)

func postHandler[E runtime.ErrorHandler](ctx context.Context, r *http.Request, body any) (any, runtime.Status) {
	var e E

	if r == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	}
	if runtime.IsDebugEnvironment() {
		status2 := context2.StatusFromContext(ctx)
		if status2 != nil {
			e.Handle(status2, runtime.RequestId(r), postLoc2)
			return nil, status2
		}
		location := r.Header.Get(http2.ContentLocation)
		if strings.HasPrefix(location, "file://") {
			// Need to deserialize return any
			return nil, runtime.NewStatusOK()
		}
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodPut:
		status := putEntry(body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), postLoc2)
		}
		return nil, status
	case http.MethodDelete:
		status := deleteEntry()
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), postLoc2)
		}
		return nil, status
	default:
		return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
	}
}

func putEntry(body any) runtime.Status {
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
	addEntry(entries)
	return runtime.StatusOK()
}

func deleteEntry() runtime.Status {
	deleteEntries()
	return runtime.StatusOK()
}
