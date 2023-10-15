package activity

import (
	"fmt"
)

func Example_AddEntry() {

	AddEntry(Entry{Agent: "controller",
		Assignment: "region:zone:sub-zone:service:instanceID",
		Action:     "Analyzing observation",
	},
	)

	fmt.Printf("test: AddEntry() -> %v\n", list)

	//Output:
	//test: AddEntry() -> [{controller region:zone:sub-zone:service:instanceID Analyzing observation}]

}
