package activity

import (
	"fmt"
)

func Example_AddEntry() {

	AddEntry(Entry{ActivityID: "uuid",
		Agent:      "agent-controller",
		Assignment: "usa:west::test-service:0123456789",
		FrameUri:   "host-frame",
		Action:     "Analyzing observation",
	},
	)

	fmt.Printf("test: AddEntry() -> %v\n", list)

	//Output:
	//test: AddEntry() -> [{0001-01-01 00:00:00 +0000 UTC uuid agent-controller usa:west::test-service:0123456789 host-frame Analyzing observation}]

}
