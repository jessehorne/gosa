package gosa

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
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
	commandQueue        []string

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

func (c *Client) callback(name string, data interface{}) {
	cb, ok := c.callbacks[name]
	if ok {
		if name == "connection" {
			cb.(func())()
		} else if name == "close" {
			cb.(func())()
		} else if name == "cmd" {
			cb.(func(Command))(data.(Command))
		}
	}
}

func (c *Client) Run() error {
	// get messages in a goroutine
	go func() {
		s := bufio.NewScanner(c.reader)
		for s.Scan() {
			c.onMessage(s.Text())
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
	c.callback("connection", nil)

	// run commands as they enter the queue, in order.
	// (we don't run commands instantly...)
	go func(c *Client) {
		for {
			if len(c.commandQueue) > 0 {
				var cmd string
				cmd, c.commandQueue = c.commandQueue[0], c.commandQueue[1:]
				c.writer.Write([]byte(cmd))
			}
		}
	}(c)

	// wait and block for command to finish
	return c.command.Wait()
}

func (c *Client) onMessage(data string) {
	if data == "Welcome to SMP agent v5.4.0.5" {
		c.ready = true
		return
	}

	if !c.ready {
		return
	}

	if len(data) < 3 {
		return
	}

	if data[:3] == "INV" {
		uri := parseForINV(data)
		if uri != "" {
			//c.responseInviteLinks = append(c.responseInviteLinks, uri)

			c.callback("cmd", Command{
				Type: "INV",
				Data: uri,
			})
		}
	}

}

func (c *Client) send(data string) {
	c.commandQueue = append(c.commandQueue, data)
}

func (c *Client) Close() {
	c.callback("close", nil)

	c.writer.Close()
	c.reader.Close()
}

func (c *Client) NewConnection(connMode, name string) {
	cmd := fmt.Sprintf("567\n567\nNEW T %s %s\n", connMode, name)
	c.send(cmd)

	//// wait for some time to get a response
	//waitTime := time.Now()
	//for {
	//	if len(c.responseInviteLinks) > 0 {
	//		uri := c.responseInviteLinks[0]
	//		c.responseInviteLinks = c.responseInviteLinks[1:]
	//
	//		return uri, nil
	//	}
	//
	//	ready := time.Since(waitTime) > time.Duration(c.commandTimeout)*time.Second
	//	if ready {
	//		return "", errors.New("timeout")
	//	}
	//}
}
