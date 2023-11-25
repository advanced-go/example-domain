package slo

import "fmt"

func Example_addEntry() {
	addEntry([]Entry{{Controller: "percentile", Threshold: "99/1s", StatusCodes: ""}})
	printEntries(list)
	//fmt.Printf("test: addEntry() -> %v %v %v\n", list[0].Controller, list[0].Threshold, list[0].StatusCodes)

	addEntry([]Entry{{Controller: "status-codes", Threshold: "10%", StatusCodes: "500,504"}})
	printEntries(list)
	//fmt.Printf("test: addEntry() -> %v\n", list)

	addEntry([]Entry{{Controller: "percentile", Threshold: "95/500ms", StatusCodes: ""}})
	printEntries(list)
	//fmt.Printf("test: addEntry() -> %v\n", list)

	//Output:
	//test: addEntry() -> {percentile, 99/1s, []}
	//test: addEntry() -> {percentile, 99/1s, []} {status-codes, 10%, 500,504}
	//test: addEntry() -> {percentile, 99/1s, []} {status-codes, 10%, 500,504} {percentile, 95/500ms, []}

}

func printEntries(entries []Entry) {
	s := ""
	for i, e := range entries {
		if i == 0 {
			fmt.Printf("test: addEntry() -> ")
		}
		code := e.StatusCodes
		if e.StatusCodes == "" {
			code = "[]"
		}
		//fmt.Printf("{%v, %v, %v} ", e.Controller, e.Threshold, code)
		s += fmt.Sprintf("{%v, %v, %v} ", e.Controller, e.Threshold, code)
	}
	fmt.Printf("%v\n", s)
}

func ExampleGetEntryByController() {
	addEntry([]Entry{{Controller: "percentile", Threshold: "99/1s", StatusCodes: ""}})
	addEntry([]Entry{{Controller: "status-codes", Threshold: "10%", StatusCodes: "500,504"}})

	ctrl := ""
	s := getEntriesByController(ctrl)
	fmt.Printf("test: getEntriesByController(%s) -> %v\n", ctrl, s)

	ctrl = "percentile"
	s = getEntriesByController(ctrl)
	if len(s) > 0 {
		s[0].Id = ""
	}
	fmt.Printf("test: getEntriesByController(%s) -> %v\n", ctrl, s)

	ctrl = "status-codes"
	s = getEntriesByController(ctrl)
	if len(s) > 0 {
		s[0].Id = ""
	}
	fmt.Printf("test: getEntriesByController(%s) -> %v\n", ctrl, s)

	addEntry([]Entry{{Controller: "percentile", Threshold: "95/500ms", StatusCodes: ""}})
	ctrl = "percentile"
	s = getEntriesByController(ctrl)
	if len(s) > 0 {
		s[0].Id = ""
	}
	fmt.Printf("test: getEntriesByController(%s) -> %v\n", ctrl, s)

	//Output:
	//test: getEntriesByController() -> []
	//test: getEntriesByController(percentile) -> [{0001-01-01 00:00:00 +0000 UTC  percentile 99/1s }]
	//test: getEntriesByController(status-codes) -> [{0001-01-01 00:00:00 +0000 UTC  status-codes 10% 500,504}]
	//test: getEntriesByController(percentile) -> [{0001-01-01 00:00:00 +0000 UTC  percentile 95/500ms }]

}
