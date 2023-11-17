package activity

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"strings"
)

type httpEntryHandlerFn func(w http.ResponseWriter, r *http.Request) runtime.Status

const (
	httpLoc        = PkgUri + "/httpHandler"
	validateVarLoc = PkgUri + "/validateVariant"
)

func httpHandler[E runtime.ErrorHandler](proxy httpEntryHandlerFn, w http.ResponseWriter, r *http.Request) runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewStatus(http.StatusBadRequest)
	}
	var e E

	if proxy != nil {
		return proxy(w, r)
	}
	/*
		statusVar := validateVariant(r)
		if !statusVar.OK() {
			e.Handle(statusVar, runtime.RequestId(r), httpLoc)
			http2.WriteResponse[E](w, nil, statusVar, nil)
			return statusVar
		}

	*/
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		buf, status := getEntry[[]byte](r.URL, r.Header.Get(http2.ContentLocation))
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), httpLoc)
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		http2.WriteResponse[E](w, buf, status, []http2.Attr{{http2.ContentType, http2.ContentTypeJson}})
		return status
	case http.MethodPut:
		status := putEntry(r.Header.Get(http2.ContentLocation), r.Body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), httpLoc)
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		http2.WriteResponse[E](w, nil, status, nil)
		return status
	case http.MethodDelete:
		status := deleteEntry(r.Header.Get(http2.ContentLocation))
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), httpLoc)
		}
		http2.WriteResponse[E](w, nil, status, nil)
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewStatus(http.StatusMethodNotAllowed)
}

func validateVariant(r *http.Request) runtime.Status {
	if r == nil {
		return runtime.NewStatus(http.StatusBadRequest)
	}
	variant := r.Header.Get(http2.ContentLocation)
	if variant != EntryV1Variant {
		s := variant
		if len(variant) == 0 {
			s = "<empty>"
		}
		err := errors.New(fmt.Sprintf("error invalid variant: [%v] for [%v]", s, PkgUri))
		return runtime.NewStatusError(runtime.StatusInvalidArgument, validateVarLoc, err).SetContent(err, false)
	}
	return runtime.NewStatusOK()
}
