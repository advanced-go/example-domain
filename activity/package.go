package activity

import (
	"encoding/json"
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
		buf, status := MarshalEntry[E](GetEntries())
		httpx.WriteResponse[E](w, buf, status)
		return status
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
	return runtime.NewStatusOK()
}

func MarshalEntry[E runtime.ErrorHandler](entry []Entry) ([]byte, *runtime.Status) {
	buf, err := json.Marshal(entry)
	if err != nil {
		var e E
		return nil, e.Handle(nil, "marshal", err)
	}
	return buf, runtime.NewStatusOK()
}

func UnmarshalEntry[E runtime.ErrorHandler](buf []byte) ([]Entry, *runtime.Status) {
	var entry []Entry

	err := json.Unmarshal(buf, &entry)
	if err != nil {
		var e E
		return nil, e.Handle(nil, "unmarshal", err)
	}
	return entry, runtime.NewStatusOK()
}
