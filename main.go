package main

import (
	"fmt"
	"github.com/jessehorne/go-simplex/gosa"
	"log"
)

func main() {
	c, err := gosa.NewClient("127.0.0.1", "5224")
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Create conversation
	//uri, err := c.NewConnection(gosa.ConnModeInvite, "subscribe")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//c.Close()
	//
	//fmt.Println(uri)

	// when you connect to the agent server
	c.On("connection", func() {
		fmt.Println("connected")
	})

	// when you get a command from the agent
	c.On("cmd", func(cmd gosa.Command) {
		fmt.Println("COMMAND: ", cmd.Type, cmd.Data)
	})

	// run before closing down...on crash or ctrl-c
	c.On("close", func() {
		fmt.Println("\ncleaning up...")
	})

	c.NewConnection(gosa.ConnModeInvite, "subscribe")

	c.Run()
}
