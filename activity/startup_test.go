package activity

import (
	"fmt"
	"github.com/advanced-go/core/http2/http2test"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/core/uri"
	"github.com/advanced-go/messaging/mux"
	"net/http"
)

func Example_Ping() {
	w := http2test.NewRecorder()
	r, _ := http.NewRequest("", "github.com/advanced-go/example-domain/activity:ping", nil)
	nid, rsc, ok := uri.UprootUrn(r.URL.Path)
	mux.ProcessPing[runtime.Output](w, nid)
	buf, status := runtime.NewBytes(w.Result())
	if !status.OK() {
		fmt.Printf("test: NewBytes() -> [status:%v]\n", status)
	}
	fmt.Printf("test: Ping() -> [nid:%v] [nss:%v] [ok:%v] [status:%v] [content:%v]\n", nid, rsc, ok, w.Result().StatusCode, string(buf))

	//Output:
	//test: Ping() -> [nid:github.com/advanced-go/example-domain/activity] [nss:ping] [ok:true] [status:200] [content:Ping status: OK, resource: github.com/advanced-go/example-domain/activity]

}
