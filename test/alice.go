package main

import (
	"fmt"
	"github.com/jessehorne/go-simplex/pkg/v1/gosa"
	"github.com/jessehorne/go-simplex/pkg/v1/structs"
	"log"
	"time"
)

func main() {
	c, err := gosa.NewClient("127.0.0.1", "5224")
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Join Connection
	uri := "simplex:/invitation#/?v=1-4&smp=smp%3A%2F%2FKr9PAzYW5qCDt8G9hzlPQMNPdcXTfVrRuW34ZQEIx9A%3D%40localhost%3A5223%2FF6CaScZLkEUu_-tyQVOC0H2UeZ77Vx5d%23%2F%3Fv%3D1-2%26dh%3DMCowBQYDK2VuAyEAXpI0hiRHMf7sqK1Xqm-l7CCao1QuJsD-NY0zTfn-IzA%253D&e2e=v%3D1-2%26x3dh%3DMEIwBQYDK2VvAzkAio58PRb78z-jW3YnrtIU90nGDSxbQBTkizxWS4a0uIbperISi23hzqDmSYZQYjG9yq8UnTg0eWA%3D%2CMEIwBQYDK2VvAzkArC73uLRwPBFVczIUBOFOpEfn3rL6HOlP9hiW0Q0MXBnFDSXxA2H9AdNyUrwpFoYGLUhajVEw40s%3D"
	i := structs.NewDefaultXInfo()
	i.Params.Profile.DisplayName = "alice"
	corrID := "2"
	connID := "alice2"

	// c *Client, connType int, corrID, connID, uri, i string
	conn := gosa.JoinConnection(c, gosa.ConnectionTypeOne, corrID, connID, uri, i)
	conn.On("connected", func(d map[string]string) {
		fmt.Println("JOINED: ", d["connID"], d["xinfo"])
		go func() {
			for x := 0; x < 2; x++ {
				conn.SendMessage("hello from alice " + string(x))
				time.Sleep(3 * time.Second)
			}
		}()
	})

	conn.On("message", func(d map[string]string) {
		// ack message so that you can get next ones :-)
		fmt.Println("RECEIVED MSG: ", d["msgId"], d["content"])
		conn.AckMessage(d["msgId"])
	})

	c.Run()
}
