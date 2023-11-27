package messages

const (
	MessageTypeConf = iota
	MessageTypeInv
	MessageTypeError
)

type Message interface {
	ToString() string
	GetType() int
}

func ToMessage(s []string) Message {
	msgConf, err := MessageConfParse(s)
	if err == nil {
		return msgConf
	}

	msgInv, err := MessageInvParse(s)
	if err == nil {
		return msgInv
	}

	msgErr, err := MessageErrorParse(s)
	if err == nil {
		return msgErr
	}

	return nil
}
