package context2

import (
	"fmt"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

func Example_StatusContext() {
	status := runtime.NewStatus(http.StatusGatewayTimeout)
	ctx := NewStatusContext(nil, status)

	status = StatusFromContext(ctx)

	fmt.Printf("test: NewStatusContext() -> %v\n", status)

	//Output:
	//test: NewStatusContext() -> Timeout

}