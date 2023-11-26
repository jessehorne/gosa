package commands

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errorInvalidNewCommand = errors.New("invalid NEW command")
)

type CommandNew struct {
	CorrID string
	ConnID string
	Type   string // INV or CON
}

func (c CommandNew) ToString() string {
	return fmt.Sprintf("%s\n%s\nNEW T %s subscribe\n", c.CorrID, c.ConnID, c.Type)
}

func (c CommandNew) GetType() int {
	return CommandTypeNew
}

func NewCommandNew(corrID, connID, t string) *CommandNew {
	return &CommandNew{
		CorrID: corrID,
		ConnID: connID,
		Type:   t,
	}
}

func CommandNewParse(lines []string) (CommandNew, error) {
	var cmd CommandNew
	if len(lines) != 3 {
		return cmd, errors.New("invalid NEW command")
	}

	corrID := lines[0]
	connID := lines[1]
	cmdStr := lines[2]

	splitCmd := strings.Split(cmdStr, " ")
	if splitCmd[0] != "NEW" {
		return cmd, errorInvalidNewCommand
	}

	if splitCmd[1] != "T" {
		return cmd, errorInvalidNewCommand
	}

	if splitCmd[2] != "INV" && splitCmd[2] != "CON" {
		return cmd, errorInvalidNewCommand
	}

	if splitCmd[3] != "subscribe" {
		return cmd, errorInvalidNewCommand
	}

	cmd.CorrID = corrID
	cmd.ConnID = connID
	cmd.Type = splitCmd[2]

	return cmd, nil
}
