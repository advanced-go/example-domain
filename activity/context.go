package activity

import (
	"context"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"strings"
)

type statusT struct{}
type contentT struct{}

var (
	statusKey  = statusT{}
	contentKey = contentT{}
)

// NewStatusContext - creates a new Context with a Status
func NewStatusContext(ctx context.Context, status runtime.Status) context.Context {
	if ctx == nil {
		ctx = context.Background()
	} else {
		i := ctx.Value(statusKey)
		if i != nil {
			return ctx
		}
	}
	return context.WithValue(ctx, statusKey, status)
}

// StatusFromContext - return a Status from a context2
func StatusFromContext(ctx any) runtime.Status {
	if ctx == nil {
		return nil
	}
	if ctx2, ok := ctx.(context.Context); ok {
		i := ctx2.Value(statusKey)
		if status, ok2 := i.(runtime.Status); ok2 {
			return status
		}
	}
	return nil
}

// NewContentLocationContext - creates a new Context with a content location
func NewContentLocationContext(ctx context.Context, h http.Header) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if h == nil {
		return ctx
	}
	v := h.Get(ContentLocation)
	if len(v) == 0 {
		return ctx
	}
	return context.WithValue(ctx, contentKey, v)
}

// ContentLocationFromContext - return a content location from a context
func ContentLocationFromContext(ctx context.Context) (string, bool) {
	if ctx == nil {
		return "", false
	}
	i := ctx.Value(contentKey)
	if i == nil {
		return "", false
	}
	if location, ok := i.(string); ok {
		if strings.HasPrefix(location, "file://") {
			return location, true
		}
	}
	return "", false
}
