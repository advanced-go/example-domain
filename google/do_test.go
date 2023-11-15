package google

import (
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime/runtimetest"
	"net/http"
)

func Example_Do() {
	req, _ := http.NewRequest("", "http://localhost:8080"+PkgPath+"?q=test", nil)
	resp, status := Do(nil, req, nil)
	if buf, ok := resp.([]byte); ok {
		if buf != nil {
		}
	}

	fmt.Printf("test: Do(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, status.Header().Get(http2.ContentType), status.Header().Get(http2.ContentLength))

	//Output:
	//test: Do(http://localhost:8080/advanced-go/example-domain/google/search?q=test) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-length:100835]

}

func Example_doHandler() {
	req, _ := http.NewRequest("", "http://localhost:8080"+PkgPath+"?q=test", nil)
	resp, status := doHandler[runtimetest.DebugError](nil, req, nil)
	if buf, ok := resp.([]byte); ok {
		if buf != nil {
		}
	}

	fmt.Printf("test: doHandler(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, status.Header().Get(http2.ContentType), status.Header().Get(http2.ContentLength))

	//Output:
	//test: doHandler(http://localhost:8080/advanced-go/example-domain/google/search?q=test) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-length:100835]

}

func Example_Another_Test() {
	fmt.Printf("test: error")

	//Output:

}
