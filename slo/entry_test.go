package slo

import "fmt"

func Example_addEntry() {
	addEntry([]EntryV1{{Controller: "percentile", Threshold: "99/1s", StatusCodes: ""}})
	fmt.Printf("test: addEntry() -> %v\n", list)

	addEntry([]EntryV1{{Controller: "status-codes", Threshold: "10%", StatusCodes: "500,504"}})
	fmt.Printf("test: addEntry() -> %v\n", list)

	addEntry([]EntryV1{{Controller: "percentile", Threshold: "95/500ms", StatusCodes: ""}})
	fmt.Printf("test: addEntry() -> %v\n", list)

	//Output:
	//test: addEntry() -> [{0001-01-01 00:00:00 +0000 UTC percentile 99/1s }]
	//test: addEntry() -> [{0001-01-01 00:00:00 +0000 UTC percentile 99/1s } {0001-01-01 00:00:00 +0000 UTC status-codes 10% 500,504}]
	//test: addEntry() -> [{0001-01-01 00:00:00 +0000 UTC percentile 99/1s } {0001-01-01 00:00:00 +0000 UTC status-codes 10% 500,504} {0001-01-01 00:00:00 +0000 UTC percentile 95/500ms }]

}

func ExampleGetEntryByController() {
	addEntry([]EntryV1{{Controller: "percentile", Threshold: "99/1s", StatusCodes: ""}})
	addEntry([]EntryV1{{Controller: "status-codes", Threshold: "10%", StatusCodes: "500,504"}})

	ctrl := ""
	s := getEntriesByController(ctrl)
	fmt.Printf("test: getEntriesByController(%s) -> %v\n", ctrl, s)

	ctrl = "percentile"
	s = getEntriesByController(ctrl)
	fmt.Printf("test: getEntriesByController(%s) -> %v\n", ctrl, s)

	ctrl = "status-codes"
	s = getEntriesByController(ctrl)
	fmt.Printf("test: getEntriesByController(%s) -> %v\n", ctrl, s)

	addEntry([]EntryV1{{Controller: "percentile", Threshold: "95/500ms", StatusCodes: ""}})

	ctrl = "percentile"
	s = getEntriesByController(ctrl)
	fmt.Printf("test: getEntriesByController(%s) -> %v\n", ctrl, s)

	//Output:
	//test: getEntriesByController() -> []
	//test: getEntriesByController(percentile) -> [{0001-01-01 00:00:00 +0000 UTC percentile 99/1s }]
	//test: getEntriesByController(status-codes) -> [{0001-01-01 00:00:00 +0000 UTC status-codes 10% 500,504}]
	//test: getEntriesByController(percentile) -> [{0001-01-01 00:00:00 +0000 UTC percentile 95/500ms }]

}
