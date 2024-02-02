package entryv1

import (
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

func statusCode(s *runtime.Status) access.StatusCodeFunc {
	return func() int {
		if s == nil || *s == nil {
			return http.StatusOK
		}
		return (*(s)).Code()
	}
}
