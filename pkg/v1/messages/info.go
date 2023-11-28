package messages

import (
	"errors"
	"fmt"
	"github.com/jessehorne/go-simplex/pkg/v1/structs"
	"strings"
)

type MessageInfo struct {
	ConnID string
	XInfo  structs.XInfo // the profile of the user you connected to
}

func (m MessageInfo) ToString() string {
	return fmt.Sprintf("\n%s\nINFO %d\n%s\n",
		m.ConnID, len(m.XInfo.ToString()), m.XInfo.ToString())
}

func (m MessageInfo) GetType() int {
	return MessageTypeInfo
}

func MessageInfoParse(lines []string) (MessageInfo, error) {
	var msg MessageInfo

	if len(lines) != 4 {
		return msg, errors.New("invalid INFO message from agent: invalid line count")
	}

	splitConf := strings.Split(lines[2], " ")

	if len(splitConf) != 2 {
		return msg, errors.New("invalid INFO message from agent")
	}

	if splitConf[0] != "INFO" {
		return msg, errors.New("invalid INFO message from agent: not INFO command")
	}

	msg.ConnID = lines[1]
	msg.XInfo = structs.XInfoFromString(lines[3])

	return msg, nil
}
