package commands

const (
	CommandTypeNew = iota
	CommandTypeConf
	CommandTypeInv
)

type Command interface {
	ToString() string
	GetType() int
}

func ToCommand(s []string) Command {
	cmdNew, err := CommandNewParse(s)
	if err == nil {
		return cmdNew
	}

	cmdConf, err := CommandConfParse(s)
	if err == nil {
		return cmdConf
	}

	cmdInv, err := CommandInvParse(s)
	if err == nil {
		return cmdInv
	}

	return nil
}
