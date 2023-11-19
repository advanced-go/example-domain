package activity

import (
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

func findGetProxy(proxies []any) func(h http.Header, uri *url.URL) (any, runtime.Status) {
	for _, p := range proxies {
		if fn, ok := p.(func(h http.Header, uri *url.URL) (any, runtime.Status)); ok {
			return fn
		}
	}
	return nil
}
