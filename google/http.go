package google

import (
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
)

func httpHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	result, status := Do(nil, r, nil)
	http2.WriteResponse[E](w, result, status, status.Header())
	return status
}
