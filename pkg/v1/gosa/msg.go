package gosa

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/jessehorne/go-simplex/pkg/v1/commands"
)

type Msg struct {
	CorrID  string
	ConnID  string
	Content string
	Acked   bool
	Cmd     *commands.CommandSend
	Message Message
}

type Message struct {
	Event  string        `json:"event"`
	MsgID  string        `json:"msgId"`
	Params MessageParams `json:"params"`
}

type MessageParams struct {
	Content MessageContent `json:"content"`
}

type MessageContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (m *Message) ToString() string {
	s, err := json.Marshal(m)
	if err != nil {
		return "{}"
	}

	return string(s)
}

func NewMsg(corrID, connID, content string) *Msg {
	// generate random 12 bytes and base64 url encode it
	token := make([]byte, 12)
	rand.Read(token)
	messageID := base64.URLEncoding.EncodeToString(token)

	jsonMessage := Message{
		Event: "x.msg.new",
		MsgID: messageID,
		Params: MessageParams{
			Content: MessageContent{
				Type: "text",
				Text: content,
			},
		},
	}

	cmd := commands.NewCommandSend(corrID, connID, jsonMessage.ToString())

	return &Msg{
		CorrID:  corrID,
		ConnID:  connID,
		Content: content,
		Cmd:     cmd,
		Message: jsonMessage,
	}
}
