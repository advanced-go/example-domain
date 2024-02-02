package activity

import (
	"fmt"
	"github.com/advanced-go/core/messaging"
	"net/http"
	"time"
)

var (
	agent messaging.Agent
)

func init() {
	var err error
	agent, err = messaging.NewDefaultAgent(PkgPath, messageHandler, false)
	if err != nil {
		fmt.Printf("init(\"%v\") failure: [%v]\n", PkgPath, err)
	}
	agent.Run()
}

func messageHandler(msg messaging.Message) {
	start := time.Now()
	//fmt.Printf("messageHandler() -> [msg%v]\n", msg)
	switch msg.Event {
	case messaging.StartupEvent:
		messaging.SendReply(msg, messaging.Status{Code: http.StatusOK, Duration: time.Since(start)})
	case messaging.ShutdownEvent:
	case messaging.PingEvent:
		messaging.SendReply(msg, messaging.Status{Code: http.StatusOK, Duration: time.Since(start)})
	}
}
