package google

import (
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

func Example_Get() {
	req, _ := http.NewRequest("", "http://localhost:8080"+"/"+PkgPath+"?q=test", nil)
	resp, status := Get(req)
	if buf, ok := resp.([]byte); ok {
		if buf != nil {
		}
	}

	fmt.Printf("test: Get(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, status.ContentHeader().Get(http2.ContentType), status.ContentHeader().Get(http2.ContentLength))

	//Output:
	//test: Do(http://localhost:8080/advanced-go/example-domain/google/search?q=test) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-length:100835]

}

func Example_getHandler() {
	req, _ := http.NewRequest("", "http://localhost:8080"+PkgPath+"?q=test", nil)
	resp, status := getHandler[runtime.TestError](req)
	if buf, ok := resp.([]byte); ok {
		if buf != nil {
		}
	}

	fmt.Printf("test: getHandler(%v) -> [status:%v] [content-type:%v] [content-length:%v]\n", req.URL.String(), status, status.ContentHeader().Get(http2.ContentType), status.ContentHeader().Get(http2.ContentLength))

	//Output:
	//test: doHandler(http://localhost:8080/advanced-go/example-domain/google/search?q=test) -> [status:OK] [content-type:text/html; charset=ISO-8859-1] [content-length:100835]

}

func Example_Another_Test() {
	fmt.Printf("test: error")

	//Output:

}
