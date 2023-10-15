package action

import (
	"fmt"
)

func Example_AddEntry() {

	AddEntry(Entry{Agent: "agent-controller",
		Assignment: "usa:west::test-service:0123456789",
		Controller: "host",
		Behavior:   "RateLimiter",
		Action:     "Set limit = 250",
	},
	)

	fmt.Printf("test: AddEntry() -> %v\n", list)

	//Output:
	//test: AddEntry() -> [{0001-01-01 00:00:00 +0000 UTC  agent-controller usa:west::test-service:0123456789 host RateLimiter Set limit = 250}]

}
