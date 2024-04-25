package service

import (
	"fmt"
	"github.com/advanced-go/stdlib/messaging"
	"net/http"
)

func ExamplePing() {
	r, _ := http.NewRequest("", PkgPath+":ping", nil)
	status := messaging.Ping(nil, r.URL)
	fmt.Printf("test: Ping() -> [status:%v]\n", status)

	//Output:
	//test: Ping() -> [status:OK]

}
