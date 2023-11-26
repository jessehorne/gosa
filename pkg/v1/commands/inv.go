package commands

import (
	"errors"
	"fmt"
	"strings"
)

type CommandInv struct {
	CorrID string
	ConnID string
	URI    string
}

func (c CommandInv) ToString() string {
	return fmt.Sprintf("%s\n%s\nINV %s\n", c.CorrID, c.CorrID, c.URI)
}

func (c CommandInv) GetType() int {
	return CommandTypeInv
}

func CommandInvParse(lines []string) (CommandInv, error) {
	var cmd CommandInv

	if len(lines) != 3 {
		return cmd, errors.New("invalid inv command from agent: invalid line count")
	}

	splitConf := strings.Split(lines[2], " ")

	if len(splitConf) != 2 {
		return cmd, errors.New("invalid INV command from agent")
	}

	if splitConf[0] != "INV" {
		return cmd, errors.New("invalid INV command from agent: doesn't include INV on third line")
	}

	cmd.CorrID = lines[0]
	cmd.ConnID = lines[1]
	cmd.URI = splitConf[1]

	return cmd, nil
}
