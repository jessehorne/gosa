package main

import (
	"fmt"
	"github.com/jessehorne/go-simplex/pkg/v1/commands"
	"github.com/jessehorne/go-simplex/pkg/v1/gosa"
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

	// when you get a NEW command from the agent...you shouldn't??
	c.On("a-cmd-new", func(s string) {
		fmt.Println("CALLBACK: a-cmd-new | COMMAND: ", s)
	})

	// when you get a CONF command from the agent...
	c.On("a-cmd-conf", func(c commands.CommandConf) {
		fmt.Println("CALLBACK: a-cmd-conf | COMMAND: ", c.SMPServerURI)
	})

	// when you get a INV command from the agent...
	c.On("a-cmd-inv", func(c commands.CommandInv) {
		fmt.Println("CALLBACK: a-cmd-inv | COMMAND: INV ", c.URI)
	})

	// when you get a ERR command from the agent...
	c.On("a-cmd-err", func(c commands.CommandError) {
		fmt.Println("CALLBACK: a-cmd-err | COMMAND: ", c.First, c.Second, c.Third)
	})

	// run before closing down...on crash or ctrl-c
	c.On("close", func() {
		fmt.Println("\nCALLBACK: close")
	})

	c.NewConnection("1", "bob", gosa.ConnModeInvite)

	c.Run()
}
