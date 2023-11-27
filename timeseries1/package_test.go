package timeseries1

import (
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"net/http"
	"reflect"
	"time"
)

var entries = []Entry{{
	Traffic:   "ingress",
	Start:     time.Now().UTC(),
	Duration:  750,
	RequestId: "98765",
	Url:       "https://www.somestupiddomain.com/help",
}}

func Example_PkgUri() {
	pkgUri2 := reflect.TypeOf(any(pkg{})).PkgPath()

	fmt.Printf("test: PkgPath = \"%v\"\n", pkgUri2)

	//Output:
	//test: PkgPath = "github.com/advanced-go/example-domain/timeseries1"

}

func Example_HttpHandler() {
	access.EnableTestLogHandler()

	// Bad Request
	uri := "http://localhost:8080/github.com/advanced-go/example-domain/timeseries1/entry"
	r, _ := http.NewRequest("GET", uri, nil)
	w := http2.NewRecorder()
	HttpHandler(w, r)
	buf, status := io2.ReadAll(w.Result().Body)
	if !status.OK() {
		fmt.Printf("test: ReadAll() -> [status:%v]\n", status)
	}
	fmt.Printf("test: HttpHandler() -> [status:%v] [content:%v]\n", w.Result().StatusCode, string(buf))

	// Resource Not Found
	uri = "http://localhost:8080/github.com/advanced-go/example-domain/timeseries1:invalid"
	r, _ = http.NewRequest("GET", uri, nil)
	w = http2.NewRecorder()
	HttpHandler(w, r)
	buf, status = io2.ReadAll(w.Result().Body)
	if !status.OK() {
		fmt.Printf("test: ReadAll() -> [status:%v]\n", status)
	}
	fmt.Printf("test: HttpHandler() -> [status:%v] [content:%v]\n", w.Result().StatusCode, string(buf))

	// Content Not Found
	uri = "http://localhost:8080/github.com/advanced-go/example-domain/timeseries1:entry"
	r, _ = http.NewRequest("GET", uri, nil)
	w = http2.NewRecorder()
	HttpHandler(w, r)
	fmt.Printf("test: HttpHandler() -> [status:%v]\n", w.Result().StatusCode)

	//Output:
	//test: HttpHandler() -> [status:400] [content:error invalid path, not a valid URN: /github.com/advanced-go/example-domain/timeseries1/entry]
	//test: HttpHandler() -> [status:404] [content:error invalid URI, resource was not found: invalid]
	//test: HttpHandler() -> [status:404]

}
