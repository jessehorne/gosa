package commands

const (
	CommandTypeNew = iota
	CommandTypeConf
	CommandTypeInv
	CommandTypeError
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

	return nil
}
