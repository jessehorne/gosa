package commands

import "fmt"

type CommandAck struct {
	CorrID    string
	ConnID    string
	MessageID string
}

func (c CommandAck) ToString() string {
	return fmt.Sprintf("%s\n%s\nACK %s\n",
		c.CorrID, c.ConnID, c.MessageID)
}

func (c CommandAck) GetType() int {
	return CommandTypeAck
}

func NewCommandAck(corrID, connID, msgID string) *CommandAck {
	return &CommandAck{
		CorrID:    corrID,
		ConnID:    connID,
		MessageID: msgID,
	}
}
