package timeseries

import (
	"fmt"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/messaging/core"
	"github.com/advanced-go/messaging/exchange"
	"time"
)

var (
	agent exchange.Agent
)

func init() {
	var status runtime.Status
	agent, status = exchange.NewDefaultAgent(PkgPath, messageHandler, false)
	if !status.OK() {
		fmt.Printf("init(\"%v\") failure: [%v]\n", PkgPath, status)
	}
	agent.Run()
}

func messageHandler(msg core.Message) {
	start := time.Now()
	//fmt.Printf("messageHandler() -> [msg%v]\n", msg)
	switch msg.Event {
	case core.StartupEvent:
		core.SendReply(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
	case core.ShutdownEvent:
	case core.PingEvent:
		core.SendReply(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
	}
}
