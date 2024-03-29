package entryv1

import (
	"fmt"
)

func createEntry(ctrl string) []Entry {
	return []Entry{{
		//CreatedTS  0
		Traffic: "ingress",
		//Start:      0),
		Duration:    500,
		RequestId:   "request-id",
		Url:         "https://service.com/path",
		Protocol:    "http",
		Host:        "search.com",
		Path:        "/path",
		Method:      "GET",
		StatusCode:  200,
		StatusFlags: "",
		Timeout:     500,
		RateLimit:   500,
		RateBurst:   100,
	}}
}

func Example_addEntry() {

	addEntries(nil, createEntry("host"))
	fmt.Printf("test: addEntry() -> %v\n", list)

	//Output:
	//test: addEntry() -> [{0001-01-01 00:00:00 +0000 UTC ingress 0001-01-01 00:00:00 +0000 UTC 500 request-id https://service.com/path http search.com /path GET 200  500 500 100}]

}

/*
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
	//test: getEntriesByController() -> []
	//test: list() -> [{0001-01-01 00:00:00 +0000 UTC ingress 0001-01-01 00:00:00 +0000 UTC 500 host us west  test-search 123-456-789 request-id https://service.com/path primary http search.com /path GET 200  500 500 100 0} {0001-01-01 00:00:00 +0000 UTC ingress 0001-01-01 00:00:00 +0000 UTC 500 ingress us west  test-search 123-456-789 request-id https://service.com/path primary http search.com /path GET 200  500 500 100 0} {0001-01-01 00:00:00 +0000 UTC ingress 0001-01-01 00:00:00 +0000 UTC 500 egress us west  test-search 123-456-789 request-id https://service.com/path primary http search.com /path GET 200  500 500 100 0} {0001-01-01 00:00:00 +0000 UTC ingress 0001-01-01 00:00:00 +0000 UTC 500 host us west  test-search 123-456-789 request-id https://service.com/path primary http search.com /path GET 200  500 500 100 0}]
	//test: getEntriesByController() -> [{0001-01-01 00:00:00 +0000 UTC ingress 0001-01-01 00:00:00 +0000 UTC 500 host us west  test-search 123-456-789 request-id https://service.com/path primary http search.com /path GET 200  500 500 100 0} {0001-01-01 00:00:00 +0000 UTC ingress 0001-01-01 00:00:00 +0000 UTC 500 host us west  test-search 123-456-789 request-id https://service.com/path primary http search.com /path GET 200  500 500 100 0}]

}


*/
