package activity

import (
	"fmt"
	"github.com/advanced-go/core/messaging"
	"github.com/advanced-go/core/uri"
	"net/http"
)

func ExamplePing() {
	r, _ := http.NewRequest("", "github/advanced-go/example-domain/activity:ping", nil)
	nid, rsc, ok := uri.UprootUrn(r.URL.Path)
	status := messaging.Ping(nil, nid)
	fmt.Printf("test: Ping() -> [nid:%v] [nss:%v] [ok:%v] [status-code:%v]\n", nid, rsc, ok, status.Code)

	//Output:
	//test: Ping() -> [nid:github/advanced-go/example-domain/activity] [nss:ping] [ok:true] [status-code:200]

}
