package gosa

import (
	"bufio"
	"fmt"
	"github.com/jessehorne/go-simplex/pkg/v1/commands"
	"github.com/jessehorne/go-simplex/pkg/v1/messages"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

type Client struct {
	Address             string
	Port                string
	writer              io.WriteCloser
	reader              io.ReadCloser
	command             *exec.Cmd
	responseInviteLinks []string
	commandTimeout      int // seconds to wait for each command to process
	ready               bool
	running             bool

	callbacks map[string]interface{}
}

func NewClient(addr string, port string) (*Client, error) {
	address := fmt.Sprintf("%s:%s", addr, port)

	cmd := exec.Command("openssl", "s_client", "-connect", address)
	writer, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	reader, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	client := &Client{
		Address: addr,
		Port:    port,

		writer: writer,
		reader: reader,

		command:        cmd,
		commandTimeout: 5,

		running:   true,
		callbacks: map[string]interface{}{},
	}

	return client, nil
}

func (c *Client) On(cb string, f interface{}) {
	c.callbacks[cb] = f
}

func (c *Client) callback(name string, data messages.Message) {
	cb, ok := c.callbacks[name]
	if ok {
		if name == "agent-connection" {
			cb.(func())()
		} else if name == "close" {
			cb.(func())()
		} else if name == "a-msg-conf" {
			cb.(func(messages.MessageConf))(data.(messages.MessageConf))
		} else if name == "a-msg-inv" {
			cb.(func(messages.MessageInv))(data.(messages.MessageInv))
		} else if name == "a-msg-err" {
			cb.(func(messages.MessageError))(data.(messages.MessageError))
		}
	}
}

func (c *Client) OnMessage(s []string) {
	msg := messages.ToMessage(s)

	if msg == nil {
		fmt.Println("NULL: ", fmt.Sprintf("%s\n%s\n%s\n", s[0], s[1], s[2]))
		return
	}

	// call low level agent callback
	if msg.GetType() == messages.MessageTypeConf {
		c.callback("a-msg-conf", msg)
	} else if msg.GetType() == messages.MessageTypeInv {
		c.callback("a-msg-inv", msg)
	} else if msg.GetType() == messages.MessageTypeError {
		c.callback("a-msg-err", msg)
	}
}

func (c *Client) Run() error {
	// get messages in a goroutine
	go func() {
		s := bufio.NewScanner(c.reader)

		count := 0 // count lines
		var messageBuffer []string
		for s.Scan() {
			if !c.ready {
				c.waitForReady(s.Text())
			} else {
				messageBuffer = append(messageBuffer, strings.TrimSuffix(s.Text(), "\r"))

				count += 1

				if count > 2 {
					c.OnMessage(messageBuffer)
					count = 0
					messageBuffer = []string{}
				}
			}
		}
	}()

	// capture CTRL-C and cleanup
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		c.Close()
		os.Exit(1)
	}()

	// start the command but don't block
	if err := c.command.Start(); err != nil {
		return err
	}

	// connection was made if Start doesn't return an error
	c.callback("agent-connection", nil)

	// wait and block for command to finish
	return c.command.Wait()
}

func (c *Client) waitForReady(data string) {
	if data == "Welcome to SMP agent v5.4.0.5" {
		c.ready = true
		return
	}
}

func (c *Client) Close() {
	c.callback("close", nil)

	c.writer.Close()
	c.reader.Close()
}

func (c *Client) NewConnection(corrID, connID, t string) {
	cmd := commands.NewCommandNew(corrID, connID, t)
	c.send(cmd.ToString())
}

func (c *Client) send(data string) {
	c.writer.Write([]byte(data))
}
