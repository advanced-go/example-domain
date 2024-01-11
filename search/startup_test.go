package search

import (
	"fmt"
	"github.com/advanced-go/core/messaging"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/core/uri"
	"net/http"
	"net/http/httptest"
)

func Example_Ping() {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("", "github.com/advanced-go/example-domain/service:ping", nil)
	nid, rsc, ok := uri.UprootUrn(r.URL.Path)
	messaging.ProcessPing[runtime.Output](w, nid)
	buf, status := runtime.NewBytes(w.Result())
	if !status.OK() {
		fmt.Printf("test: NewBytes() -> [status:%v]\n", status)
	}
	fmt.Printf("test: Ping() -> [nid:%v] [nss:%v] [ok:%v] [status:%v] [content:%v]\n", nid, rsc, ok, w.Result().StatusCode, string(buf))

	//Output:
	//test: Ping() -> [nid:github.com/advanced-go/example-domain/search] [nss:ping] [ok:true] [status:200] [content:Ping status: OK, resource: github.com/advanced-go/example-domain/search]

}
