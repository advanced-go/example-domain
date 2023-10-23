package timeseries

import (
	"github.com/go-ai-agent/core/exchange"
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
	"sync/atomic"
)

type pkg struct{}

var (
	PkgUrl  = runtime.ParsePkgUrl(reflect.TypeOf(any(pkg{})).PkgPath())
	PkgUri  = PkgUrl.Host + PkgUrl.Path
	started int64
)

// IsPkgStarted - returns status of startup
func IsPkgStarted() bool {
	return atomic.LoadInt64(&started) != 0
}

func DoHandler(req *http.Request) (*http.Response, error) {
	recorder := httpx.NewRecorder()
	status := entryHandler[runtime.BypassError](recorder, req)
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
	switch r.Method {
	case "GET":
		var entries []Entry

		name := ""
		if r.URL.Query() != nil {
			name = r.URL.Query().Get(ConrollerName)
		}
		if len(name) != 0 {
			entries = GetEntriesByController(name)
		} else {
			entries = GetEntries()
		}
		buf, status := MarshalEntry[E](entries)
		exchange.WriteResponse[E](w, buf, status)
		return status
	case "PUT":
		buf, status := httpx.ReadAll[E](r.Body)
		if !status.OK() {
			exchange.WriteResponse[E](w, nil, status)
			return status
		}
		entries, status1 := UnmarshalEntry[E](buf)
		if status.OK() {
			AddEntry(entries)
		}
		exchange.WriteResponse[E](w, nil, status1)
		return status1
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
	return runtime.NewStatusOK()
}
