package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jessehorne/go-simplex/pkg/v1/structs"
	"strings"
)

type CommandConf struct {
	ConnID       string
	Key          string
	SMPServerURI string
	Port         string
	XInfo        structs.XInfo
}

func (c CommandConf) ToString() string {
	var jsonData []byte
	var err error

	jsonData, err = json.Marshal(c.XInfo)
	if err != nil {
		jsonData = []byte{}
	}
	return fmt.Sprintf("%s\nCONF %s %s %s\n%s\n", c.ConnID, c.Key, c.SMPServerURI, c.Port, jsonData)
}

func (c CommandConf) GetType() int {
	return CommandTypeConf
}

func CommandConfParse(lines []string) (CommandConf, error) {
	var cmd CommandConf

	if len(lines) != 3 {
		return cmd, errors.New("invalid CONF command from agent: invalid line count")
	}

	splitConf := strings.Split(lines[1], " ")

	if len(splitConf) != 4 {
		return cmd, errors.New("invalid CONF command from agent")
	}

	cmd.ConnID = lines[0]
	cmd.Key = splitConf[1]
	cmd.SMPServerURI = splitConf[2]
	cmd.Port = splitConf[3]

	var x structs.XInfo
	err := json.Unmarshal([]byte(lines[2]), &x)
	if err != nil {
		return cmd, errors.New("couldn't unmarshal CONF agent XInfo data")
	}

	return cmd, nil
}
