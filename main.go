package main

import (
	"fmt"
	"github.com/jessehorne/go-simplex/pkg/v1/gosa"
	"github.com/jessehorne/go-simplex/pkg/v1/messages"
	"github.com/jessehorne/go-simplex/pkg/v1/structs"
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
		fmt.Println("CALLBACK: a-cmd-inv | MSG: INV ", m.CorrID, m.ConnID, m.URI)
	})

	// when you get a ERR messages from the agent...
	c.On("a-msg-err", func(m messages.MessageError) {
		fmt.Println("CALLBACK: a-cmd-err | MSG: ", m.First, m.Second, m.Third)
	})

	// run before closing down...on crash or ctrl-c
	c.On("close", func() {
		fmt.Println("\nCALLBACK: close")
	})

	//c.NewConnection("1", "bob", gosa.ConnModeInvite)
	//c.NewConnection("1", "bob", gosa.ConnModeConnection)
	testURI := "simplex:/invitation#/?v=1-4&smp=smp%3A%2F%2FKr9PAzYW5qCDt8G9hzlPQMNPdcXTfVrRuW34ZQEIx9A%3D%40localhost%3A5223%2FCHbBxAPoLkJUPz5BarcqRbHA-_YogRl2%23%2F%3Fv%3D1-2%26dh%3DMCowBQYDK2VuAyEApJxXvnNu8isvb43vjj1svYrroFHfcocq7vhkVTbwY3k%253D&e2e=v%3D1-2%26x3dh%3DMEIwBQYDK2VvAzkAW9dVkc-X6qFArFniQMFB-M6Gvpxh-2rTTZW3yGvXL9leImy2aSljAVjV6j3Y6KpOpGtUvBbC1dc%3D%2CMEIwBQYDK2VvAzkANTSe-hsM2wy8SpluVnXcRZdY8FDmcYslEyhkHIKesDJwF96toC8HYN95gj1qLbmA2pitpIogANw%3D"
	testInfo := structs.NewDefaultXInfo()
	corrID := "2"
	connID := "alice"
	c.JoinConnection(corrID, connID, testURI, testInfo)

	c.Run()
}
