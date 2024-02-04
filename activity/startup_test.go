package activity

import (
	"fmt"
	"github.com/advanced-go/core/messaging"
	"net/http"
)

func ExamplePing() {
	r, _ := http.NewRequest("", PkgPath+"ping", nil)
	status := messaging.Ping(nil, r.URL)
	fmt.Printf("test: Ping() -> [status-code:%v]\n", status.Code)

	//Output:
	//test: Ping() -> [status-code:200]

}
