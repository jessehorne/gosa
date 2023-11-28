package gosa

import "github.com/jessehorne/go-simplex/pkg/v1/structs"

type User struct {
	ConfirmID    string
	SMPServerURI string
	XInfo        structs.XInfo
}

func NewUser(confirmID, uri, xinfo string) *User {
	return &User{
		ConfirmID:    confirmID,
		SMPServerURI: uri,
		XInfo:        structs.XInfoFromString(xinfo),
	}
}
