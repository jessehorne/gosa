package messages

import (
	"errors"
	"fmt"
	"strings"
)

type MessageInv struct {
	CorrID string
	ConnID string
	URI    string
}

func (m MessageInv) ToString() string {
	return fmt.Sprintf("%s\n%s\nINV %s\n", m.CorrID, m.CorrID, m.URI)
}

func (m MessageInv) GetType() int {
	return MessageTypeInv
}

func MessageInvParse(lines []string) (MessageInv, error) {
	var msg MessageInv

	if len(lines) != 3 {
		return msg, errors.New("invalid inv message from agent: invalid line count")
	}

	splitConf := strings.Split(lines[2], " ")

	if len(splitConf) != 2 {
		return msg, errors.New("invalid INV message from agent")
	}

	if splitConf[0] != "INV" {
		return msg, errors.New("invalid INV message from agent: doesn't include INV on third line")
	}

	msg.CorrID = lines[0]
	msg.ConnID = lines[1]
	msg.URI = splitConf[1]

	return msg, nil
}
