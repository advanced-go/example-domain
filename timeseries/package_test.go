package timeseries

import (
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"net/http"
	"net/http/httptest"
)

func Example_HttpHandler() {
	access.EnableTestLogger()

	// Bad Request
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("", "https://localhost:8080/advanced-go/example-domain/timeseries/entry", nil)
	HttpHandler(rec, req)
	buf, _ := http2.ReadAll(rec.Result())
	fmt.Printf("test: HttpHandler() -> [status:%v] [body:%v]\n", rec.Code, string(buf))

	// Not Found - invalid resource
	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("", "https://localhost:8080/github.com/advanced-go/example-domain/timeseries:entry", nil)
	HttpHandler(rec, req)
	buf, _ = http2.ReadAll(rec.Result())
	fmt.Printf("test: HttpHandler() -> [status:%v] [body:%v]\n", rec.Code, string(buf))

	// Not Found - no content
	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("", "https://localhost:8080/github.com/advanced-go/example-domain/timeseries:v2/entry", nil)
	HttpHandler(rec, req)
	buf, _ = http2.ReadAll(rec.Result())
	fmt.Printf("test: HttpHandler() -> [status:%v] [body:%v]\n", rec.Code, string(buf))

	//Output:
	//test: HttpHandler() -> [status:400] [body:error invalid path, not a valid URN: /advanced-go/example-domain/timeseries/entry]
	//test: HttpHandler() -> [status:404] [body:error invalid URI, resource was not found: entry]
	//test: HttpHandler() -> [status:404] [body:]

}
