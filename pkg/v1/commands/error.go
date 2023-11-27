package commands

import (
	"errors"
	"fmt"
	"strings"
)

type CommandError struct {
	First  string
	Second string
	Third  string
}

func (c CommandError) ToString() string {
	return fmt.Sprintf("%s\n%s\n%s\n", c.First, c.Second, c.Third)
}

func (c CommandError) GetType() int {
	return CommandTypeError
}

func CommandErrorParse(lines []string) (CommandError, error) {
	var cmd CommandError

	if len(lines) != 3 {
		return cmd, errors.New("invalid error command from agent: invalid line count")
	}

	splitLast := strings.Split(lines[2], " ")

	if splitLast[0] != "ERR" {
		return cmd, errors.New("invalid error command: last line doesn't contain ERR")
	}

	cmd.First = lines[0]
	cmd.Second = lines[1]
	cmd.Third = lines[2]

	return cmd, nil
}
