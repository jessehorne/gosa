package commands

import (
	"fmt"
	"github.com/jessehorne/go-simplex/pkg/v1/structs"
)

type CommandLet struct {
	CorrID    string
	ConnID    string
	ConfirmID string
	XInfo     structs.XInfo
}

func (c CommandLet) ToString() string {
	info := c.XInfo.ToString()
	l := len(info)

	return fmt.Sprintf("%s\n%s\nLET %s %d\n%s\n",
		c.CorrID,
		c.ConnID,
		c.ConfirmID,
		l,
		info,
	)
}

func (c CommandLet) GetType() int {
	return CommandTypeLet
}

func NewCommandLet(corrID, connID, confirmID string, i structs.XInfo) *CommandLet {
	return &CommandLet{
		CorrID:    corrID,
		ConnID:    connID,
		ConfirmID: confirmID,
		XInfo:     i,
	}
}
