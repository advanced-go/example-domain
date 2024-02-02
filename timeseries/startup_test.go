package timeseries

import (
	"fmt"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/core/uri"
	"net/http"
	"net/http/httptest"
)

func _Example_Ping() {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("", "github/advanced-go/example-domain/timeseries:ping", nil)
	nid, rsc, ok := uri.UprootUrn(r.URL.Path)
	//messaging.ProcessPing[runtime.Output](w, nid)
	buf, status := runtime.ReadAll(w.Result().Body, nil)
	if !status.OK() {
		fmt.Printf("test: NewBytes() -> [status:%v]\n", status)
	}
	fmt.Printf("test: Ping() -> [nid:%v] [nss:%v] [ok:%v] [status:%v] [content:%v]\n", nid, rsc, ok, w.Result().StatusCode, string(buf))

	//Output:
	//test: Ping() -> [nid:github/advanced-go/example-domain/timeseries] [nss:ping] [ok:true] [status:200] [content:Ping status: OK, resource: github/advanced-go/example-domain/timeseries]

}
