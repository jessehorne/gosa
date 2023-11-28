package main

import (
	"fmt"
	"github.com/jessehorne/go-simplex/pkg/v1/gosa"
	"github.com/jessehorne/go-simplex/pkg/v1/structs"
	"log"
)

func main() {
	c, err := gosa.NewClient("127.0.0.1", "5224")
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Join Connection
	uri := "simplex:/invitation#/?v=1-4&smp=smp%3A%2F%2FKr9PAzYW5qCDt8G9hzlPQMNPdcXTfVrRuW34ZQEIx9A%3D%40localhost%3A5223%2F2tAWG-RicwBySX3imdeJ9FS_JAfe5Jp-%23%2F%3Fv%3D1-2%26dh%3DMCowBQYDK2VuAyEAx1XCaHrP0SX0uXB-XhPykUP0sO3i7A2DcKKW4dnYjGE%253D&e2e=v%3D1-2%26x3dh%3DMEIwBQYDK2VvAzkA462gO5mAIcuoLEdOTkeGJLqXVC88j6kHaNgIM57_2uD1yAbsuVoe0Q8KYKp-QOuF2vZMuzU7cEs%3D%2CMEIwBQYDK2VvAzkAf_6x9MFAblbgFG8NKMMceR9LO2fdXbJj60MOpUD0JAC0lMYnYMtO_Ez5elqJU0Me79uPve79904%3D"
	i := structs.NewDefaultXInfo()
	i.Params.Profile.DisplayName = "alice"
	corrID := "2"
	connID := "alice"

	// c *Client, connType int, corrID, connID, uri, i string
	conn := gosa.JoinConnection(c, gosa.ConnectionTypeOne, corrID, connID, uri, i)
	conn.On("connected", func(d map[string]string) {
		fmt.Println("JOINED: ", d["connID"], d["xinfo"])
	})

	c.Run()
}
