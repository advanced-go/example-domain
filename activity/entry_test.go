package activity

import (
	"fmt"
)

func Example_addEntry() {

	addEntry([]entry{{ActivityID: "activity-uuid",
		ActivityType: "trace",
		Agent:        "agent-controller",
		AgentUri:     "https://host/agent-path",
		Assignment:   "usa:west::test-service:0123456789",
		FrameUri:     "https://host/frame-path",
		Controller:   "host-controller",
		Behavior:     "RateLimiting",
		Description:  "Analyzing observation",
	}},
	)

	fmt.Printf("test: addEntry() -> %v\n", list)

	//Output:
	//test: addEntry() -> [{0001-01-01 00:00:00 +0000 UTC activity-uuid trace agent-controller https://host/agent-path usa:west::test-service:0123456789 https://host/frame-path host-controller RateLimiting Analyzing observation}]

}

func Example_getEntriesByType() {

	addEntry([]entry{{ActivityID: "urn:uuid:1",
		ActivityType: "trace",
		Agent:        "agent-controller",
		AgentUri:     "https://host/agent-path",
		Assignment:   "usa:west::test-service:0123456789",
		FrameUri:     "https://host/frame-path",
		Controller:   "host-controller",
		Behavior:     "RateLimiting",
		Description:  "Analyzing observation",
	}},
	)

	addEntry([]entry{{ActivityID: "urn:uuid:2",
		ActivityType: "action",
		Agent:        "agent-controller",
		AgentUri:     "https://host/agent-path",
		Assignment:   "usa:west::test-service:0123456789",
		FrameUri:     "https://host/frame-path",
		Controller:   "host-controller",
		Behavior:     "RateLimiting",
		Description:  "Reduced rate limit",
	}},
	)

	addEntry([]entry{{ActivityID: "urn:uuid:3",
		ActivityType: "action",
		Agent:        "agent-controller",
		Assignment:   "usa:west::test-service:0123456789",
		FrameUri:     "https://host/frame-path",
		Controller:   "host-controller",
		Behavior:     "RateLimiting",
		Description:  "Reduced rate burst",
	}},
	)
	e := getEntriesByType("invalid")
	fmt.Printf("test: getEntriesByType() %v\n", e)

	e = getEntriesByType("trace")
	fmt.Printf("test: getEntriesByType(trace) %v\n", e)

	e = getEntriesByType("action")
	fmt.Printf("test: getEntriesByType(action) %v\n", e)

	/*
		e, err = getEntriesByActivityType[[]entry]("trace")
		fmt.Printf("test: getEntriesByActivityType[[]entry](trace) [err:%v] [entry:%v]\n", err, e)

		buf, err2 = getEntriesByActivityType[[]byte]("trace")
		fmt.Printf("test: getEntriesByActivityType[[]byte](trace) [err:%v] [entry:%v]\n", err2, string(buf))

		e, err = getEntriesByActivityType[[]entry]("action")
		fmt.Printf("test: getEntriesByActivityType[[]entry](action) [err:%v] [entry:%v]\n", err, e)

		buf, err2 = getEntriesByActivityType[[]byte]("action")
		fmt.Printf("test: getEntriesByActivityType[[]byte](action) [err:%v] [entry:%v]\n", err2, string(buf))


	*/

	//Output:
	//test: getEntriesByType() []
	//test: getEntriesByType(trace) [{0001-01-01 00:00:00 +0000 UTC activity-uuid trace agent-controller https://host/agent-path usa:west::test-service:0123456789 https://host/frame-path host-controller RateLimiting Analyzing observation} {0001-01-01 00:00:00 +0000 UTC urn:uuid:1 trace agent-controller https://host/agent-path usa:west::test-service:0123456789 https://host/frame-path host-controller RateLimiting Analyzing observation}]
	//test: getEntriesByType(action) [{0001-01-01 00:00:00 +0000 UTC urn:uuid:2 action agent-controller https://host/agent-path usa:west::test-service:0123456789 https://host/frame-path host-controller RateLimiting Reduced rate limit} {0001-01-01 00:00:00 +0000 UTC urn:uuid:3 action agent-controller  usa:west::test-service:0123456789 https://host/frame-path host-controller RateLimiting Reduced rate burst}]

}
