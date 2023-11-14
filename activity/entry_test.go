package activity

import (
	"fmt"
)

func Example_addEntry() {

	addEntry([]EntryV1{{ActivityID: "activity-uuid",
		ActivityType: "trace",
		Agent:        "agent-controller",
		AgentUri:     "https://host/agent-path",
		Assignment:   "usa:west::test-service:0123456789",
		Controller:   "host-controller",
		Behavior:     "RateLimiting",
		Description:  "Analyzing observation",
	}},
	)

	fmt.Printf("test: addEntry() -> %v\n", list)

	//Output:
	//{ "activity": "trace" "agent": "agent-controller"  "controller": "host-controller"  "message": "Analyzing observation"  }
	//test: addEntry() -> [{0001-01-01 00:00:00 +0000 UTC activity-uuid trace agent-controller https://host/agent-path usa:west::test-service:0123456789 host-controller RateLimiting Analyzing observation}]

}

func Example_getEntriesByType() {

	addEntry([]EntryV1{{ActivityID: "urn:uuid:1",
		ActivityType: "trace",
		Agent:        "agent-controller",
		AgentUri:     "https://host/agent-path",
		Assignment:   "usa:west::test-service:0123456789",
		Controller:   "host-controller",
		Behavior:     "RateLimiting",
		Description:  "Analyzing observation",
	}},
	)

	addEntry([]EntryV1{{ActivityID: "urn:uuid:2",
		ActivityType: "action",
		Agent:        "agent-controller",
		AgentUri:     "https://host/agent-path",
		Assignment:   "usa:west::test-service:0123456789",
		Controller:   "host-controller",
		Behavior:     "RateLimiting",
		Description:  "Reduced rate limit",
	}},
	)

	addEntry([]EntryV1{{ActivityID: "urn:uuid:3",
		ActivityType: "action",
		Agent:        "agent-controller",
		Assignment:   "usa:west::test-service:0123456789",
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
		e, err = getEntriesByActivityType[[]EntryV1]("trace")
		fmt.Printf("test: getEntriesByActivityType[[]EntryV1](trace) [err:%v] [entry:%v]\n", err, e)

		buf, err2 = getEntriesByActivityType[[]byte]("trace")
		fmt.Printf("test: getEntriesByActivityType[[]byte](trace) [err:%v] [entry:%v]\n", err2, string(buf))

		e, err = getEntriesByActivityType[[]EntryV1]("action")
		fmt.Printf("test: getEntriesByActivityType[[]EntryV1](action) [err:%v] [entry:%v]\n", err, e)

		buf, err2 = getEntriesByActivityType[[]byte]("action")
		fmt.Printf("test: getEntriesByActivityType[[]byte](action) [err:%v] [entry:%v]\n", err2, string(buf))


	*/

	//Output:
	//{ "activity": "trace" "agent": "agent-controller"  "controller": "host-controller"  "message": "Analyzing observation"  }
	//{ "activity": "action" "agent": "agent-controller"  "controller": "host-controller"  "message": "Reduced rate limit"  }
	//{ "activity": "action" "agent": "agent-controller"  "controller": "host-controller"  "message": "Reduced rate burst"  }
	//test: getEntriesByType() []
	//test: getEntriesByType(trace) [{0001-01-01 00:00:00 +0000 UTC activity-uuid trace agent-controller https://host/agent-path usa:west::test-service:0123456789 host-controller RateLimiting Analyzing observation} {0001-01-01 00:00:00 +0000 UTC urn:uuid:1 trace agent-controller https://host/agent-path usa:west::test-service:0123456789 host-controller RateLimiting Analyzing observation}]
	//test: getEntriesByType(action) [{0001-01-01 00:00:00 +0000 UTC urn:uuid:2 action agent-controller https://host/agent-path usa:west::test-service:0123456789 host-controller RateLimiting Reduced rate limit} {0001-01-01 00:00:00 +0000 UTC urn:uuid:3 action agent-controller  usa:west::test-service:0123456789 host-controller RateLimiting Reduced rate burst}]

}

func Example_Log() {
	e := EntryV1{
		//CreatedTS:    time.Time{},
		ActivityID:   "",
		ActivityType: "trace",
		Agent:        "agent-test",
		AgentUri:     "",
		Assignment:   "",
		Controller:   "controller-test",
		Behavior:     "",
		Description:  "test description",
	}
	fmt.Printf("test: logActivity() -> %v\n", e)
	logActivity(e)

	//Output:
	//test: logActivity() -> {0001-01-01 00:00:00 +0000 UTC  trace agent-test   controller-test  test description}
	//{ "activity": "trace" "agent": "agent-test"  "controller": "controller-test"  "message": "test description"  }

}
