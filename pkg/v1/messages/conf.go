package messages

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jessehorne/go-simplex/pkg/v1/structs"
	"strings"
)

type MessageConf struct {
	ConnID       string
	ConfirmID    string
	SMPServerURI string
	XInfoLength  string
	XInfo        structs.XInfo
}

func (m MessageConf) ToString() string {
	var jsonData []byte
	var err error

	jsonData, err = json.Marshal(m.XInfo)
	if err != nil {
		jsonData = []byte{}
	}
	return fmt.Sprintf("%s\nCONF %s %s %s\n%s\n",
		m.ConnID, m.ConfirmID, m.SMPServerURI, m.XInfoLength, string(jsonData))
}

func (m MessageConf) GetType() int {
	return MessageTypeConf
}

func MessageConfParse(lines []string) (MessageConf, error) {
	var msg MessageConf

	if len(lines) != 4 {
		return msg, errors.New("invalid CONF message from agent: invalid line count")
	}

	splitConf := strings.Split(lines[2], " ")

	if len(splitConf) != 4 {
		return msg, errors.New("invalid CONF message from agent")
	}

	if splitConf[0] != "CONF" {
		return msg, errors.New("invalid CONF message from agent: not CONF command")
	}

	msg.ConnID = lines[1]
	msg.ConfirmID = splitConf[1]
	msg.SMPServerURI = splitConf[2]
	msg.XInfoLength = splitConf[3]

	var x structs.XInfo
	err := json.Unmarshal([]byte(lines[3]), &x)
	if err != nil {
		return msg, errors.New("couldn't unmarshal CONF agent XInfo data")
	}
	msg.XInfo = x

	return msg, nil
}
