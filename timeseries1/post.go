package timeseries1

import (
	"context"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"io"
	"net/http"
	"strings"
)

const (
	postLoc = PkgPath + ":postEntryHandler"
	putLoc  = PkgPath + ":putEntry"
)

func postEntryHandler(ctx context.Context, r *http.Request, body any) (any, runtime.Status) {
	if r == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent)
	}

	if runtime.IsDebugEnvironment() {
		status2 := runtime.StatusFromContext(ctx)
		if status2 != nil {
			return nil, status2.AddLocation(postLoc)
		}
		location := r.Header.Get(http2.ContentLocation)
		if strings.HasPrefix(location, "file://") {
			// Need to deserialize return any
			return nil, runtime.NewStatusOK()
		}
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodPut:
		return nil, putEntry(body).AddLocation(postLoc)
	case http.MethodDelete:
		return nil, deleteEntry().AddLocation(postLoc)
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
	return runtime.NewStatusOK()
}

func deleteEntry() runtime.Status {
	deleteEntries()
	return runtime.NewStatusOK()
}
