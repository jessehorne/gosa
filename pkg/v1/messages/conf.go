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
	Key          string
	SMPServerURI string
	Port         string
	XInfo        structs.XInfo
}

func (m MessageConf) ToString() string {
	var jsonData []byte
	var err error

	jsonData, err = json.Marshal(m.XInfo)
	if err != nil {
		jsonData = []byte{}
	}
	return fmt.Sprintf("%s\nCONF %s %s %s\n%s\n", m.ConnID, m.Key, m.SMPServerURI, m.Port, jsonData)
}

func (m MessageConf) GetType() int {
	return MessageTypeConf
}

func MessageConfParse(lines []string) (MessageConf, error) {
	var msg MessageConf

	if len(lines) != 3 {
		return msg, errors.New("invalid CONF message from agent: invalid line count")
	}

	splitConf := strings.Split(lines[1], " ")

	if len(splitConf) != 4 {
		return msg, errors.New("invalid CONF message from agent")
	}

	msg.ConnID = lines[0]
	msg.Key = splitConf[1]
	msg.SMPServerURI = splitConf[2]
	msg.Port = splitConf[3]

	var x structs.XInfo
	err := json.Unmarshal([]byte(lines[2]), &x)
	if err != nil {
		return msg, errors.New("couldn't unmarshal CONF agent XInfo data")
	}

	return msg, nil
}
