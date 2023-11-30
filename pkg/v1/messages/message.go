package messages

const (
	MessageTypeConf = iota
	MessageTypeInv
	MessageTypeInfo
	MessageTypeError
	MessageTypeMsg
)

type Message interface {
	ToString() string
	GetType() int
}

func ToMessage(s []string) Message {
	msgMsg, err := MessageMsgParse(s)
	if err == nil {
		return msgMsg
	}

	msgConf, err := MessageConfParse(s)
	if err == nil {
		return msgConf
	}

	msgInv, err := MessageInvParse(s)
	if err == nil {
		return msgInv
	}

	msgInfo, err := MessageInfoParse(s)
	if err == nil {
		return msgInfo
	}

	msgErr, err := MessageErrorParse(s)
	if err == nil {
		return msgErr
	}

	return nil
}
