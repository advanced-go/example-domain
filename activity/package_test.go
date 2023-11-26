package activity

import (
	"fmt"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/http2"
	"net/http"
	"reflect"
)

func Example_PkgUri() {
	pkgUri2 := reflect.TypeOf(any(pkg{})).PkgPath()
	fmt.Printf("test: PkgPath  = \"%v\"\n", pkgUri2)

	//Output:
	//test: PkgPath  = "github.com/advanced-go/example-domain/activity"

}

func Example_HttpHandler() {
	access.EnableTestLogHandler()
	uri := "http://localhost:8080/github.com/advanced-go/example-domain/activity:entry"
	//uri := "/github.com/advanced-go/example-domain/activity:entry"

	r, _ := http.NewRequest("GET", uri, nil)
	w := http2.NewRecorder()
	HttpHandler(w, r)

	fmt.Printf("test: HttpHandler() -> %v", w.Result())

	//Output:

}
