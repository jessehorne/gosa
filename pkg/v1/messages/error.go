package messages

import (
	"errors"
	"fmt"
	"strings"
)

type MessageError struct {
	First  string
	Second string
	Third  string
}

func (m MessageError) ToString() string {
	return fmt.Sprintf("%s\n%s\n%s\n", m.First, m.Second, m.Third)
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

	msg.First = lines[0]
	msg.Second = lines[1]
	msg.Third = lines[2]

	return msg, nil
}
