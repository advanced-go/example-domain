package timeseries

import (
	"fmt"
	"time"
)

func createEntry(ctrl string) Entry {
	return Entry{Traffic: "ingress",
		Start:       time.Now().UTC(),
		Duration:    time.Millisecond * 500,
		Controller:  ctrl,
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
	}
}

func Example_AddEntry() {

	AddEntry(createEntry("host"))
	fmt.Printf("test: AddEntry() -> %v\n", list)

	//Output:

}

func Example_GetEntriesByController() {
	deleteEntries()

	e := GetEntriesByController("host")
	fmt.Printf("test: GetEntriesByController() -> %v\n", e)

	AddEntry(createEntry("host"))
	AddEntry(createEntry("ingress"))
	AddEntry(createEntry("egress"))
	AddEntry(createEntry("host"))
	fmt.Printf("test: List() -> %v\n", list)

	e = GetEntriesByController("host")
	fmt.Printf("test: GetEntriesByController() -> %v\n", e)

	//Output:

}
