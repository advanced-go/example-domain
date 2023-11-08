package activity

import (
	"github.com/go-ai-agent/core/runtime"
	"net/http"
)

type DoHandlerFn func(ctx any, r *http.Request, body any) (any, *runtime.Status)

// newTypeHandler - templated function providing a TypeHandlerFn with a closure
func newTypeHandler2[E runtime.ErrorHandler]() DoHandlerFn {
	return func(ctx any, r *http.Request, body any) (any, *runtime.Status) {
		return doHandler2[E](ctx, r, body)
	}
}

// BodyConstraints2 - defining constraints for the TypeHandler body
type BodyConstraints2 interface {
	[]EntryV1 | []byte | runtime.Nillable
}

func Do2[T BodyConstraints2](ctx any, method, uri, variant string, body T) (any, *runtime.Status) {
	//return controller.Apply(httpx.UpdateHeadersAndContext(r), body)
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, runtime.NewStatusError(http.StatusBadRequest, "/Do", err)
	}
	req.Header.Set(runtime.ContentLocation, variant)
	return doHandler2[runtime.LogError](ctx, req, body)
}

func doHandler2[E runtime.ErrorHandler](ctx any, r *http.Request, body any) (any, *runtime.Status) {
	return nil, nil
}
