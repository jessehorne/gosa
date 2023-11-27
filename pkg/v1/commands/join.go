package commands

import (
	"fmt"
	"github.com/jessehorne/go-simplex/pkg/v1/structs"
)

type CommandJoin struct {
	CorrID string
	ConnID string
	URI    string
	XInfo  structs.XInfo
}

func (c CommandJoin) ToString() string {
	info := c.XInfo.ToString()
	l := len(info)

	return fmt.Sprintf("%s\n%s\nJOIN T %s subscribe %d\n%s\n",
		c.CorrID,
		c.ConnID,
		c.URI,
		l,
		info)
}

func (c CommandJoin) GetType() int {
	return CommandTypeJoin
}

func NewCommandJoin(corrID, connID, uri string, i structs.XInfo) *CommandJoin {
	return &CommandJoin{
		CorrID: corrID,
		ConnID: connID,
		URI:    uri,
		XInfo:  i,
	}
}
