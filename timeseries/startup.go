package timeseries

import (
	"github.com/advanced-go/core/runtime/startup"
	"sync/atomic"
)

var (
	c       = make(chan startup.Message, 1)
	started int64
	//duration = time.Second * 4
	location = PkgPath + "/startup"
)

// isPkgStarted - returns status of startup
func isPkgStarted() bool {
	return atomic.LoadInt64(&started) != 0
}

func init() {
	startup.Register(PkgUri, c)
	go receive()
}

var messageHandler startup.MessageHandler = func(msg startup.Message) {
	//start := time.Now()
	switch msg.Event {
	case startup.StartupEvent:
	case startup.ShutdownEvent:
	}
}

func receive() {
	for {
		select {
		case msg, open := <-c:
			if !open {
				return
			}
			go messageHandler(msg)
		default:
		}
	}
}
