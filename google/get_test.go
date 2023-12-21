package google

import (
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

func _Example_Get() {
	req, _ := http.NewRequest("", "http://localhost:8080"+"/"+PkgPath+"/search?q=test", nil)
	resp, status := Get(req)
	if buf, ok := resp.([]byte); ok {
		if buf != nil {
		}
	}

	fmt.Printf("test: Get(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, status.ContentHeader().Get(http2.ContentType), status.ContentHeader().Get(http2.ContentLength))

	//Output:
	//test: Get(http://localhost:8080/github.com/advanced-go/example-domain/google/search?q=test) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-length:100835]

}

func Example_getHandler() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+"/search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	resp, status := getHandler[runtime.Output](req)
	if buf, ok := resp.([]byte); ok {
		if buf != nil {
		}
	}
	fmt.Printf("test: getHandler(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, status.ContentHeader().Get(http2.ContentType), status.ContentHeader().Get(http2.ContentLength))

	//Output:
	//test: getHandler(http://localhost:8080/github.com/advanced-go/example-domain/google/search?q=test) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-length:100835]

}

func getHandlerOverrideFail(id string) (string, string) {
	switch id {
	case searchTag:
		return "file://[cwd]/resource/query-result.txt", ""
	}
	return "", ""
}

func Example_getHandler_OverrideFail() {
	setOverride(getHandlerOverrideFail, "")
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+"/search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	resp, status := getHandler[runtime.Output](req)
	if buf, ok := resp.([]byte); ok {
		if buf != nil {
		}
	}
	fmt.Printf("test: getHandler(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, status.ContentHeader().Get(http2.ContentType), status.ContentHeader().Get(http2.ContentLength))

	//Output:
	//test: getHandler(http://localhost:8080/github.com/advanced-go/example-domain/google/search?q=test) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-length:100835]

}

func getHandlerOverrideSuccess(id string) (string, string) {
	switch id {
	case searchTag:
		return "file://[cwd]/googletest/resource/query-result.txt", ""
	}
	return "", ""
}

func Example_getHandler_OverrideSuccess() {
	setOverride(getHandlerOverrideSuccess, "")
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+"/"+PkgPath+"/search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: NewRequest() -> %v\n", err)
	}
	resp, status := getHandler[runtime.Output](req)
	if buf, ok := resp.([]byte); ok {
		if buf != nil {
		}
	}
	fmt.Printf("test: getHandler(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, status.ContentHeader().Get(http2.ContentType), status.ContentHeader().Get(http2.ContentLength))

	//Output:
	//test: getHandler(http://localhost:8080/github.com/advanced-go/example-domain/google/search?q=golang) -> [status:OK] [content-type:text/plain] [content-length:47]

}
