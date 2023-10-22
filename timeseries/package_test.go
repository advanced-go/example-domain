package timeseries

import (
	"fmt"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"net/http/httptest"
	"time"
)

func init() {
	AddEntry([]Entry{{Traffic: "ingress",
		Start:       time.Now().UTC(),
		Duration:    time.Millisecond * 500,
		Controller:  "host",
		Region:      "us",
		Zone:        "west",
		SubZone:     "",
		Service:     "test-service",
		InstanceId:  "123-456-789",
		RequestId:   "request-id",
		Url:         "https://service.com/path",
		Route:       "primary",
		Protocol:    "http",
		Host:        "service.com",
		Path:        "/path",
		Method:      "GET",
		StatusCode:  200,
		StatusFlags: "",
		Timeout:     500,
		RateLimit:   500,
		RateBurst:   100,
		RoutePct:    0,
	}},
	)
}

func Example_entryHandler() {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "www.someuri.com", nil)
	entryHandler[runtime.DebugError](rec, req)

	//resp := exchange.R
	fmt.Printf("test: Response() %v\n", rec)

	//Output:
}
