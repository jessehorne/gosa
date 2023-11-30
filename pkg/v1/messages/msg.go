package messages

import (
	"errors"
	"fmt"
	"strings"
)

type MessageMsg struct {
	ConnID  string
	Type    string // OK or ERR
	R       string
	B       string
	S       string
	Content string
}

func (m MessageMsg) ToString() string {
	return fmt.Sprintf("\n%s\nMSG %s R=%s B=%s S=%s F %s\n%s\n",
		m.ConnID, m.Type, m.R, m.B, m.S, len(m.Content), m.Content)
}

func (m MessageMsg) GetType() int {
	return MessageTypeMsg
}

func MessageMsgParse(lines []string) (MessageMsg, error) {
	var msg MessageMsg

	if len(lines) != 4 {
		return msg, errors.New("invalid msg message from agent: invalid line count")
	}

	// MSG OK R=6,2023-11-29T04:27:29.241Z B=IjmF5i94fYo6hk+ymsmVivf1DUJoNYRI,2023-11-29T04:26:30.000Z S=5 F 111
	splitConf := strings.Split(lines[2], " ")

	if splitConf[0] != "MSG" {
		return msg, errors.New("invalid msg message from agent: third line doesn't begin with MSG")
	}

	if len(splitConf) != 7 {
		return msg, errors.New("invalid msg message from agent: not good data")
	}

	msg.ConnID = lines[1]
	msg.Type = splitConf[1]
	msg.R = splitConf[2][2:]
	msg.B = splitConf[3][2:]
	msg.S = splitConf[4][2:]
	msg.Content = lines[3]

	return msg, nil
}
