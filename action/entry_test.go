package action

import (
	"fmt"
)

func Example_AddEntry() {

	AddEntry(Entry{Controller: "host",
		Behavior: "RateLimiter",
		Action:   "Set limit = 250",
	},
	)

	fmt.Printf("test: AddEntry() -> %v\n", list)

	//Output:
	//test: AddEntry() -> [{host RateLimiter Set limit = 250}]

}
