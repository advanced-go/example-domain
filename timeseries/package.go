package timeseries

import (
	"errors"
	"fmt"
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/log2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
)

var (
	Pattern = pkgPath + "/"

	PkgUri  = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath = runtime.PathFromUri(PkgUri)

	postEntryLoc2  = PkgUri + "/PostEntry"
	validateVarLoc = PkgUri + "/validateVariant"

	EntryV1Variant = PkgUri + "/" + reflect.TypeOf(EntryV1{}).Name()
	EntryV2Variant = PkgUri + "/" + reflect.TypeOf(EntryV2{}).Name()
)

// GetEntryConstraints - Get constraints
type GetEntryConstraints interface {
	[]EntryV1 | []EntryV2
}

// GetEntry - generic get function with context and uri for resource selection and filtering
func GetEntry[T GetEntryConstraints](ctx any, uri string) (t T, status *runtime.Status) {
	defer log2.Log(ctx, "GET", uri, log2.NewStatusCodeClosure(&status))()
	return getEntryHandler[T, runtime.LogError](ctx, uri)
}

// PostEntryConstraints - Post constraints
type PostEntryConstraints interface {
	[]EntryV1 | []byte | runtime.Nillable
}

// PostEntry - exchange function
func PostEntry[T PostEntryConstraints](ctx any, method, uri, variant string, body T) (any, *runtime.Status) {
	var e runtime.LogError

	req, status := http2.NewRequest(ctx, method, uri, variant, nil)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(ctx), postEntryLoc2)
		return nil, status
	}
	return postWrapper(ctx, req, body)
}

// HttpHandler - http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpWrapper(nil, w, r)
}

func validateVariant(r *http.Request) *runtime.Status {
	if r == nil {
		return runtime.NewStatus(http.StatusBadRequest)
	}
	variant := r.Header.Get(http2.ContentLocation)
	if variant != EntryV1Variant && variant != EntryV2Variant {
		s := variant
		if len(variant) == 0 {
			s = "<empty>"
		}
		err := errors.New(fmt.Sprintf("error invalid variant: [%v] for [%v]", s, PkgUri))
		return runtime.NewStatusError(runtime.StatusInvalidArgument, validateVarLoc, err).SetContent(err, false)
	}
	return runtime.NewStatusOK()
}
