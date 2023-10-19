package activity

import (
	"fmt"
)

func Example_AddEntry() {

	AddEntry(Entry{ActivityID: "activity-uuid",
		ActivityType: "trace",
		Agent:        "agent-controller",
		AgentUri:     "https://host/agent-path",
		Assignment:   "usa:west::test-service:0123456789",
		FrameUri:     "https://host/frame-path",
		Controller:   "host-controller",
		Behavior:     "RateLimiting",
		Description:  "Analyzing observation",
	},
	)

	fmt.Printf("test: AddEntry() -> %v\n", list)

	//Output:
	//test: AddEntry() -> [{0001-01-01 00:00:00 +0000 UTC activity-uuid trace agent-controller https://host/agent-path usa:west::test-service:0123456789 https://host/frame-path host-controller RateLimiting Analyzing observation}]

}

func Example_GetEntriesByType() {

	AddEntry(Entry{ActivityID: "urn:uuid:1",
		ActivityType: "trace",
		Agent:        "agent-controller",
		AgentUri:     "https://host/agent-path",
		Assignment:   "usa:west::test-service:0123456789",
		FrameUri:     "https://host/frame-path",
		Controller:   "host-controller",
		Behavior:     "RateLimiting",
		Description:  "Analyzing observation",
	},
	)

	AddEntry(Entry{ActivityID: "urn:uuid:2",
		ActivityType: "action",
		Agent:        "agent-controller",
		AgentUri:     "https://host/agent-path",
		Assignment:   "usa:west::test-service:0123456789",
		FrameUri:     "https://host/frame-path",
		Controller:   "host-controller",
		Behavior:     "RateLimiting",
		Description:  "Reduced rate limit",
	},
	)

	AddEntry(Entry{ActivityID: "urn:uuid:3",
		ActivityType: "action",
		Agent:        "agent-controller",
		Assignment:   "usa:west::test-service:0123456789",
		FrameUri:     "https://host/frame-path",
		Controller:   "host-controller",
		Behavior:     "RateLimiting",
		Description:  "Reduced rate burst",
	},
	)
	e, err := GetEntriesByType[[]Entry]("activity")
	fmt.Printf("test: GetEntriesByType[[]Entry](activity) [err:%v] [entry:%v]\n", err, e)

	buf, err2 := GetEntriesByType[[]byte]("activity")
	fmt.Printf("test: GetEntriesByType[[]byte](activity) [err:%v] [entry:%v]\n", err2, string(buf))

	e, err = GetEntriesByType[[]Entry]("trace")
	fmt.Printf("test: GetEntriesByType[[]Entry](trace) [err:%v] [entry:%v]\n", err, e)

	buf, err2 = GetEntriesByType[[]byte]("trace")
	fmt.Printf("test: GetEntriesByType[[]byte](trace) [err:%v] [entry:%v]\n", err2, string(buf))

	e, err = GetEntriesByType[[]Entry]("action")
	fmt.Printf("test: GetEntriesByType[[]Entry](action) [err:%v] [entry:%v]\n", err, e)

	buf, err2 = GetEntriesByType[[]byte]("action")
	fmt.Printf("test: GetEntriesByType[[]byte](action) [err:%v] [entry:%v]\n", err2, string(buf))

	//Output:
	//test: GetEntriesByType[[]Entry](activity) [err:<nil>] [entry:[]]
	//test: GetEntriesByType[[]byte](activity) [err:<nil>] [entry:null]
	//test: GetEntriesByType[[]Entry](trace) [err:<nil>] [entry:[{0001-01-01 00:00:00 +0000 UTC activity-uuid trace agent-controller https://host/agent-path usa:west::test-service:0123456789 https://host/frame-path host-controller RateLimiting Analyzing observation} {0001-01-01 00:00:00 +0000 UTC urn:uuid:1 trace agent-controller https://host/agent-path usa:west::test-service:0123456789 https://host/frame-path host-controller RateLimiting Analyzing observation}]]
	//test: GetEntriesByType[[]byte](trace) [err:<nil>] [entry:[{"CreatedTS":"0001-01-01T00:00:00Z","ActivityID":"activity-uuid","ActivityType":"trace","Agent":"agent-controller","AgentUri":"https://host/agent-path","Assignment":"usa:west::test-service:0123456789","FrameUri":"https://host/frame-path","Controller":"host-controller","Behavior":"RateLimiting","Description":"Analyzing observation"},{"CreatedTS":"0001-01-01T00:00:00Z","ActivityID":"urn:uuid:1","ActivityType":"trace","Agent":"agent-controller","AgentUri":"https://host/agent-path","Assignment":"usa:west::test-service:0123456789","FrameUri":"https://host/frame-path","Controller":"host-controller","Behavior":"RateLimiting","Description":"Analyzing observation"}]]
	//test: GetEntriesByType[[]Entry](action) [err:<nil>] [entry:[{0001-01-01 00:00:00 +0000 UTC urn:uuid:2 action agent-controller https://host/agent-path usa:west::test-service:0123456789 https://host/frame-path host-controller RateLimiting Reduced rate limit} {0001-01-01 00:00:00 +0000 UTC urn:uuid:3 action agent-controller  usa:west::test-service:0123456789 https://host/frame-path host-controller RateLimiting Reduced rate burst}]]
	//test: GetEntriesByType[[]byte](action) [err:<nil>] [entry:[{"CreatedTS":"0001-01-01T00:00:00Z","ActivityID":"urn:uuid:2","ActivityType":"action","Agent":"agent-controller","AgentUri":"https://host/agent-path","Assignment":"usa:west::test-service:0123456789","FrameUri":"https://host/frame-path","Controller":"host-controller","Behavior":"RateLimiting","Description":"Reduced rate limit"},{"CreatedTS":"0001-01-01T00:00:00Z","ActivityID":"urn:uuid:3","ActivityType":"action","Agent":"agent-controller","AgentUri":"","Assignment":"usa:west::test-service:0123456789","FrameUri":"https://host/frame-path","Controller":"host-controller","Behavior":"RateLimiting","Description":"Reduced rate burst"}]]

}
