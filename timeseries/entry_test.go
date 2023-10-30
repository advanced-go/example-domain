package timeseries

import (
	"fmt"
	"time"
)

func createEntry(ctrl string) []EntryV1 {
	return []EntryV1{{Traffic: "ingress",
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
	}}
}

func Example_addEntry() {

	addEntry(createEntry("host"))
	fmt.Printf("test: addEntry() -> %v\n", list)

	//Output:
	//test: addEntry() -> [{ingress 2023-10-21 19:01:59.1121739 +0000 UTC 500ms host us west  test-service 123-456-789 request-id https://service.com/path primary http service.com /path GET 200  500 500 100 0}]

}

func Example_getEntriesByController() {
	deleteEntries()

	e := getEntriesByController("host")
	fmt.Printf("test: getEntriesByController() -> %v\n", e)

	addEntry(createEntry("host"))
	addEntry(createEntry("ingress"))
	addEntry(createEntry("egress"))
	addEntry(createEntry("host"))
	fmt.Printf("test: list() -> %v\n", list)

	e = getEntriesByController("host")
	fmt.Printf("test: getEntriesByController() -> %v\n", e)

	//Output:

}
