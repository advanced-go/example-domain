package slo

import (
	"fmt"
	"github.com/advanced-go/stdlib/messaging"
	"net/http"
)

func ExamplePing() {
	r, _ := http.NewRequest("", PkgPath+":ping", nil)
	//nid, rsc, ok := uri.UprootUrn(r.URL.Path)
	status := messaging.Ping(nil, r.URL)
	fmt.Printf("test: Ping() -> [status:%v]\n", status)

	//Output:
	//test: Ping() -> [status:OK]

}
