package service

import (
	"github.com/advanced-go/core/runtime"
	"net/http"
	"strings"
)

func searchHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) (status runtime.Status) {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewStatus(http.StatusBadRequest)
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		return runtime.StatusOK()
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return runtime.NewStatus(http.StatusMethodNotAllowed)
	}
}
