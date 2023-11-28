package main

import (
	"fmt"
	"github.com/jessehorne/go-simplex/pkg/v1/gosa"
	"log"
)

func main() {
	c, err := gosa.NewClient("127.0.0.1", "5224")
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Create Connections

	// Create connection
	// CorrID: 1
	// ConnID: bob
	conn := gosa.NewConnection(c, gosa.ConnectionTypeOne)
	conn.Create("1", "bob")
	conn.XInfo.Params.Profile.DisplayName = "bob"

	// called when the connection is ready for users to connect to it
	// example data...
	// {
	//    "uri": "SMP SERVER URI",
	// }
	conn.On("connection-ready", func(d map[string]string) {
		fmt.Println("CONNECTION READY: ", d)
	})

	// called when another user attempts to join your connection using the url YOU provided them
	// example data...
	// {
	//    "confirm": "abcdefghijkl",
	//    "serverURI": "smp://whatever.onion:1234",
	//    "xinfo": "!!! SEE structs.XInfo.. this is used for  !!!"
	// }
	conn.On("user-joined", func(d map[string]string) {
		fmt.Println("USER JOINED: ", d)
		conn.AllowConnection(d["confirm"], d["serverURI"], d["xinfo"])
	})

	// called when gosa receives any error from the agent
	// {
	//    "msg": "some error message here"
	// }
	conn.On("error", func(d map[string]string) {
		fmt.Println("ERROR: ", d)
	})

	c.Run()
}
