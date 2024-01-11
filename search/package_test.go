package search

import (
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/http/httptest"
	"reflect"
)

func Example_PkgUri() {
	pkgUri := reflect.TypeOf(any(pkg{})).PkgPath()
	fmt.Printf("test: PkgPath  = \"%v\"\n", pkgUri)

	//Output:
	//test: PkgPath  = "github.com/advanced-go/example-domain/search"

}

func Example_HttpHandler_4xx() {
	access.EnableTestLogger()

	// Bad Request
	uri := "http://localhost:8080/github.com/advanced-go/example-domain/service/entry"
	r, _ := http.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()
	HttpHandler(w, r)
	buf, status := runtime.NewBytes(w.Result().Body)
	if !status.OK() {
		fmt.Printf("test: NewBytes() -> [status:%v]\n", status)
	}
	fmt.Printf("test: HttpHandler() -> [status:%v] [content:%v]\n", w.Result().StatusCode, string(buf))

	// Resource Not Found
	uri = "http://localhost:8080/github.com/advanced-go/example-domain/service:invalid"
	r, _ = http.NewRequest("GET", uri, nil)
	w = httptest.NewRecorder()
	HttpHandler(w, r)
	buf, status = runtime.NewBytes(w.Result())
	if !status.OK() {
		fmt.Printf("test: NewBytes() -> [status:%v]\n", status)
	}
	fmt.Printf("test: HttpHandler() -> [status:%v] [content:%v]\n", w.Result().StatusCode, string(buf))

	// Content Not Found
	//uri = "http://localhost:8080/github.com/advanced-go/example-domain/activity:entry"
	//r, _ = http.NewRequest("GET", uri, nil)
	//w = httptest.NewRecorder()
	//HttpHandler(w, r)
	//fmt.Printf("test: HttpHandler() -> [status:%v]\n", w.Result().StatusCode)

	//Output:
	//test: HttpHandler() -> [status:400] [content:error invalid path, not a valid URN: /github.com/advanced-go/example-domain/search/entry]
	//test: HttpHandler() -> [status:404] [content:error invalid URI, resource was not found: invalid]

}
