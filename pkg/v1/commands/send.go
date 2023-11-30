package commands

import (
	"fmt"
)

type CommandSend struct {
	CorrID  string
	ConnID  string
	Content string
}

func (c CommandSend) ToString() string {
	return fmt.Sprintf("%s\n%s\nSEND F %d\n%s\n",
		c.CorrID, c.ConnID, len(c.Content), c.Content)
}

func (c CommandSend) GetType() int {
	return CommandTypeSend
}

func NewCommandSend(corrID, connID, content string) *CommandSend {
	return &CommandSend{
		CorrID:  corrID,
		ConnID:  connID,
		Content: content,
	}
}
