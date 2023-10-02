package accesslog

import (
	"context"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"net/url"
)

func GetMetrics[E runtime.ErrorHandler](ctx context.Context, url *url.URL) ([]byte, *runtime.Status) {
	return nil, runtime.NewStatusOK()
}

func SetMetrics[E runtime.ErrorHandler](ctx context.Context, req *http.Request) *runtime.Status {
	return runtime.NewStatusOK()
}
