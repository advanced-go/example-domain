package timeseries

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"strings"
)

const (
	httpLoc        = PkgPath + "/httpHandler"
	validateVarLoc = PkgPath + "/validateVariant"
)

func httpHandler[E runtime.ErrorHandler](ctx context.Context, w http.ResponseWriter, r *http.Request) runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewStatus(http.StatusBadRequest)
	}
	var e E

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
		buf, status := getEntryHandler[[]byte](ctx, r.Header, r.URL)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), httpLoc)
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		http2.WriteResponse[E](w, buf, status, []http2.Attr{{http2.ContentType, http2.ContentTypeJson}})
		return status
	case http.MethodPut:
		_, status := postEntryHandler(ctx, r, r.Body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), httpLoc)
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		http2.WriteResponse[E](w, nil, status, nil)
		return status
	case http.MethodDelete:
		_, status := postEntryHandler(ctx, r, nil)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), httpLoc)
		}
		http2.WriteResponse[E](w, nil, status.SetRequestId(r), nil)
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
	if variant != EntryV1Variant && variant != EntryV2Variant {
		s := variant
		if len(variant) == 0 {
			s = "<empty>"
		}
		err := errors.New(fmt.Sprintf("error invalid variant: [%v] for [%v]", s, PkgPath))
		return runtime.NewStatusError(runtime.StatusInvalidArgument, validateVarLoc, err).SetContent(err, false)
	}
	return runtime.NewStatusOK()
}

//if buf == nil {
//	nc := runtime.NewStatus(runtime.StatusInvalidContent)
//	http2.WriteResponse[E](w, nil, nc, nil)
//	return nc
//}
//status = json2.Unmarshal(buf, &entries)
//if !status.OK() {
//	e.Handle(status, requestId, httpLoc)
//} else {
//  addEntry(entries)
//}
