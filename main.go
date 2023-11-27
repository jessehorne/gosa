package main

import (
	"fmt"
	"github.com/jessehorne/go-simplex/pkg/v1/gosa"
	"github.com/jessehorne/go-simplex/pkg/v1/messages"
	"log"
)

func main() {
	c, err := gosa.NewClient("127.0.0.1", "5224")
	if err != nil {
		log.Fatalln(err)
		return
	}

	// when you connect to the agent server
	c.On("agent-connect", func() {
		fmt.Println("CALLBACK: agent-connect")
	})

	// called when initial data is sent...a-cmd-new will be called for all new messages from the agent
	c.On("agent-ready", func() {
		fmt.Println("CALLBACK: agent-ready")
	})

	// when you get a CONF messages from the agent...
	c.On("a-msg-conf", func(m messages.MessageConf) {
		fmt.Println("CALLBACK: a-cmd-conf | MSG: ", m.SMPServerURI)
	})

	// when you get a INV messages from the agent...
	c.On("a-msg-inv", func(m messages.MessageInv) {
		fmt.Println("CALLBACK: a-cmd-inv | MSG: INV ", m.URI)
	})

	// when you get a ERR messages from the agent...
	c.On("a-msg-err", func(m messages.MessageError) {
		fmt.Println("CALLBACK: a-cmd-err | MSG: ", m.First, m.Second, m.Third)
	})

	// run before closing down...on crash or ctrl-c
	c.On("close", func() {
		fmt.Println("\nCALLBACK: close")
	})

	c.NewConnection("1", "bob", gosa.ConnModeInvite)

	c.Run()
}
