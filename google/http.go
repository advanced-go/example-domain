package google

import (
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

func httpHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) runtime.Status {
	result, status := Search(r)
	http2.WriteResponse[E](w, result, status, status.ContentHeader())
	return status
}
