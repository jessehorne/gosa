package commands

import (
	"fmt"
	"github.com/jessehorne/go-simplex/pkg/v1/structs"
)

type CommandAccept struct {
	CorrID       string
	ConnID       string
	InvitationID string
	XInfo        structs.XInfo
}

func (c CommandAccept) ToString() string {
	return fmt.Sprintf("%s\n%s\nACPT %s %d\n%s\n",
		c.CorrID,
		c.ConnID,
		c.InvitationID,
		len(c.XInfo.ToString()),
		c.XInfo.ToString())
}

func (c CommandAccept) GetType() int {
	return CommandTypeAccept
}

func NewCommandAccept(corrID, connID, invID string, xinfo structs.XInfo) *CommandAccept {
	return &CommandAccept{
		CorrID:       corrID,
		ConnID:       connID,
		InvitationID: invID,
		XInfo:        xinfo,
	}
}
