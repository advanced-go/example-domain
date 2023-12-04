package activity

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
	status := exchange.Register(exchange.NewMailbox(PkgPath, false))
	if status.OK() {
		agent, status = exchange.NewAgent(PkgPath, messageHandler, nil)
	}
	if !status.OK() {
		fmt.Printf("init() failure: [%v]\n", PkgPath)
	}
	agent.Run()
}

func messageHandler(msg core.Message) {
	start := time.Now()
	//fmt.Printf("messageHandler() -> [msg%v]\n", msg)
	switch msg.Event {
	case core.StartupEvent:
	case core.ShutdownEvent:
	case core.PingEvent:
		core.SendReply(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
	}
}
