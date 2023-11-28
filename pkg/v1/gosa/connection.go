package gosa

import (
	"github.com/jessehorne/go-simplex/pkg/v1/commands"
	"github.com/jessehorne/go-simplex/pkg/v1/structs"
)

const (
	ConnectionStateNew = iota
	ConnectionStateConnected
	ConnectionStateApproved
	ConnectionStateDeleted
)

const (
	ConnectionTypeOne  = "INV"
	ConnectionTypeMany = "CON"
)

type Connection struct {
	Client    *Client
	Type      string
	CorrID    string
	ConnID    string
	ConfirmID string
	URI       string
	State     int
	Callbacks map[string]func(map[string]string)
	Users     map[string]*User
	XInfo     structs.XInfo
}

func NewConnection(c *Client, connType string) *Connection {
	return &Connection{
		Client:    c,
		Type:      connType,
		State:     ConnectionStateNew,
		Callbacks: map[string]func(map[string]string){},
		Users:     map[string]*User{},
		XInfo:     structs.NewDefaultXInfo(),
	}
}

func JoinConnection(c *Client, connType, corrID, connID, uri string, i structs.XInfo) *Connection {
	conn := &Connection{
		Client:    c,
		Type:      connType,
		State:     ConnectionStateNew,
		Callbacks: map[string]func(map[string]string){},
		Users:     map[string]*User{},
		XInfo:     i,
		CorrID:    corrID,
		ConnID:    connID,
		URI:       uri,
	}

	cmd := commands.NewCommandJoin(corrID, connID, uri, i)
	c.send(cmd.ToString())

	c.registerConnection(conn)

	return conn
}

func (c *Connection) Create(corrID, connID string) {
	c.CorrID = corrID
	c.ConnID = connID

	// register connection with gosa
	c.Client.registerConnection(c)

	// send command to agent to create new connection
	cmd := commands.NewCommandNew(corrID, connID, c.Type)
	c.Client.send(cmd.ToString())
}

func (c *Connection) On(name string, f func(map[string]string)) {
	c.Callbacks[name] = f
}

func (c *Connection) Callback(name string, data map[string]string) {
	cb, ok := c.Callbacks[name]
	if ok {
		cb(data)
	}
}

func (c *Connection) AllowConnection(confirmID, uri, info string) {
	// add user to connections list of users
	c.Users[confirmID] = NewUser(confirmID, uri, info)

	if c.Type == ConnectionTypeOne {
		// if we found the user from confirmID, let's send a LET
		// LET will accept the connection and give them our user profile info
		cmd := commands.NewCommandLet(c.CorrID, c.ConnID, confirmID, c.XInfo)
		c.Client.send(cmd.ToString())
	}
}
