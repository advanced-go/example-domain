package timeseries

import (
	"github.com/go-ai-agent/core/exchange"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
)

func DoHandler(req *http.Request) (*http.Response, error) {
	recorder := exchange.NewRecorder()
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
		buf, status := exchange.ReadAll[E](r.Body)
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
