package slo

import (
	"fmt"
	"github.com/advanced-go/core/messaging"
	"github.com/advanced-go/core/uri"
	"net/http"
)

func ExamplePing() {
	r, _ := http.NewRequest("", "github/advanced-go/example-domain/slo:ping", nil)
	nid, rsc, ok := uri.UprootUrn(r.URL.Path)
	status := messaging.Ping(nil, nid)
	fmt.Printf("test: Ping() -> [nid:%v] [nss:%v] [ok:%v] [status:%v]\n", nid, rsc, ok, status)

	//Output:
	//test: Ping() -> [nid:github/advanced-go/example-domain/slo] [nss:ping] [ok:true] [status:200]

}
