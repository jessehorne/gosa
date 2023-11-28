package messages

import (
	"errors"
	"fmt"
	"strings"
)

type MessageError struct {
	CorrID string
	ConnID string
	Msg    string
}

func (m MessageError) ToString() string {
	return fmt.Sprintf("%s\n%s\n%s\n", m.CorrID, m.ConnID, m.Msg)
}

func (m MessageError) GetType() int {
	return MessageTypeError
}

func MessageErrorParse(lines []string) (MessageError, error) {
	var msg MessageError

	if len(lines) != 3 {
		return msg, errors.New("invalid error message from agent: invalid line count")
	}

	splitLast := strings.Split(lines[2], " ")

	if splitLast[0] != "ERR" {
		return msg, errors.New("invalid error message: last line doesn't contain ERR")
	}

	msg.CorrID = lines[0]
	msg.ConnID = lines[1]
	msg.Msg = lines[2]

	return msg, nil
}
