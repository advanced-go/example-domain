package timeseries

import (
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/core/uri"
	"github.com/advanced-go/messaging/mux"
	"net/http"
	"net/http/httptest"
)

func Example_Ping() {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("", "github.com/advanced-go/example-domain/timeseries:ping", nil)
	nid, rsc, ok := uri.UprootUrn(r.URL.Path)
	mux.ProcessPing[runtime.Output](w, nid)
	buf, status := http2.ReadAll(w.Result())
	if !status.OK() {
		fmt.Printf("test: ReadAll() -> [status:%v]\n", status)
	}
	fmt.Printf("test: Ping() -> [nid:%v] [nss:%v] [ok:%v] [status:%v] [content:%v]\n", nid, rsc, ok, w.Result().StatusCode, string(buf))

	//Output:
	//test: Ping() -> [nid:github.com/advanced-go/example-domain/timeseries] [nss:ping] [ok:true] [status:200] [content:Ping status: OK, resource: github.com/advanced-go/example-domain/timeseries]

}
