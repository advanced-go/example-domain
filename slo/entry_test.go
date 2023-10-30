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
	//test: addEntry() -> [{percentile 99/1s }]
	//test: addEntry() -> [{percentile 99/1s } {status-codes 10% 500,504}]
	//test: addEntry() -> [{percentile 99/1s } {status-codes 10% 500,504} {percentile 95/500ms }]

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

	//test: getEntriesByController() -> []
	//test: getEntriesByController(percentile) -> [{percentile 99/1s }]
	//test: getEntriesByController(status-codes) -> [{status-codes 10% 500,504}]
	//test: getEntriesByController(percentile) -> [{percentile 95/500ms }]

}
