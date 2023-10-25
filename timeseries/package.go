package timeseries

import (
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
	"sync/atomic"
)

type pkg struct{}

var (
	PkgUrl    = runtime.ParsePkgUrl(reflect.TypeOf(any(pkg{})).PkgPath())
	PkgUri    = PkgUrl.Host + PkgUrl.Path
	EntryPath = PkgUrl.Path + "/entry"
	started   int64
)

// IsPkgStarted - returns status of startup
func IsPkgStarted() bool {
	return atomic.LoadInt64(&started) != 0
}

func DoHandler(req *http.Request) (*http.Response, error) {
	recorder := httpx.NewRecorder()
	status := entryHandler[runtime.DebugError](recorder, req)
	var err error
	if status.IsErrors() {
		err = status.Errors()[0]
	}
	return recorder.Result(), err
}

func EntryHandler(w http.ResponseWriter, r *http.Request) {
	entryHandler[runtime.LogError](w, r)
}

func entryHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewHttpStatusCode(http.StatusBadRequest)
	}
	requestId := runtime.GetOrCreateRequestId(r)
	if r.Header.Get(runtime.XRequestId) == "" {
		r.Header.Set(runtime.XRequestId, requestId)
	}
	// Need to create new request context??
	switch r.Method {
	case http.MethodGet:
		buf, status := runtime.MarshalType(queryEntries(r))
		status.SetRequestId(requestId)
		if !status.OK() || buf == nil {
			httpx.WriteMinResponse[E](w, status)
			return status
		}
		httpx.WriteResponse[E](w, buf, status, runtime.ContentType, runtime.ContentTypeJson)
		return status
	case http.MethodPut:
		var entries []entry

		buf, status := httpx.ReadAll(r.Body)
		if !status.OK() || buf == nil {
			httpx.WriteMinResponse[E](w, status.SetRequestId(requestId))
			return status
		}
		entries, status = runtime.UnmarshalType[[]entry](buf)
		if status.OK() {
			addEntry(entries)
		}
		httpx.WriteMinResponse[E](w, status.SetRequestId(requestId))
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewHttpStatusCode(http.StatusMethodNotAllowed)
}

func queryEntries(r *http.Request) []entry {
	name := ""
	if r.URL.Query() != nil {
		name = r.URL.Query().Get(ConrollerName)
	}
	if len(name) != 0 {
		return getEntriesByController(name)
	} else {
		return getEntries()
	}
	return nil
}
