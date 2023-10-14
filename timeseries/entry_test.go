package timeseries

import (
	"fmt"
	"time"
)

func Example_AddEntry() {

	AddEntry(Entry{Traffic: "ingress",
		Start:       time.Now().UTC(),
		Duration:    time.Millisecond * 500,
		Region:      "us",
		Zone:        "west",
		SubZone:     "",
		Service:     "test-service",
		InstanceId:  "123-456-789",
		RequestId:   "request-id",
		Url:         "http://service.com/path",
		Route:       "primary",
		Protocol:    "http",
		Host:        "service.com",
		Path:        "/path",
		Method:      "GET",
		StatusCode:  200,
		StatusFlags: "",

		// Needed to verify client controller configuration matches configuration in cloud
		// Can this be replaced with a periodic audit?
		Timeout:   500,
		RateLimit: 500,
		RateBurst: 100,
		RoutePct:  0,
	},
	)

	fmt.Printf("test: AddEntry() -> %v\n", list)

	//Output:

}
