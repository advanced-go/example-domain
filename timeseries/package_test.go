package timeseries

import (
	"fmt"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"net/http/httptest"
	"time"
)

func init() {
	AddEntry([]Entry{
		{
			Traffic:     "ingress",
			Start:       time.Now().UTC(),
			Duration:    time.Millisecond * 800,
			Controller:  "host",
			Region:      "usa",
			Zone:        "west",
			SubZone:     "",
			Service:     "access-log",
			InstanceId:  "123-456-789",
			RequestId:   "request-id-1",
			Url:         "https://access-log.com/example-domain/timeseries/entry",
			Route:       "primary",
			Protocol:    "http",
			Host:        "access-log.com",
			Path:        "/example-domain/timeseries/entry",
			Method:      "GET",
			StatusCode:  200,
			StatusFlags: "",
			Timeout:     500,
			RateLimit:   500,
			RateBurst:   100,
			RoutePct:    0,
		},
		{
			Traffic:     "egress",
			Start:       time.Now().UTC(),
			Duration:    time.Millisecond * 100,
			Controller:  "egress-update",
			Region:      "usa",
			Zone:        "east",
			SubZone:     "",
			Service:     "access-log",
			InstanceId:  "789-012-345",
			RequestId:   "request-id-2",
			Url:         "https://access-log.com/example-domain/timeseries/entry",
			Route:       "primary",
			Protocol:    "http",
			Host:        "access-log.com",
			Path:        "/example-domain/timeseries/entry",
			Method:      "PUT",
			StatusCode:  202,
			StatusFlags: "",
			Timeout:     400,
			RateLimit:   400,
			RateBurst:   50,
			RoutePct:    0,
		},
	},
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
