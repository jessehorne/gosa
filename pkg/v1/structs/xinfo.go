package structs

import (
	"encoding/json"
	"fmt"
)

type XInfoParams struct {
	Profile XInfoProfile `json:"profile"`
}

type XInfoProfile struct {
	ContactLink string           `json:"contactLink"`
	DisplayName string           `json:"displayName"`
	FullName    string           `json:"fullName"`
	Image       string           `json:"image"`
	Preferences XInfoPreferences `json:"preferences"`
}

type XInfoPreferences struct {
	Calls         XInfoCalls         `json:"calls"`
	FullDelete    XInfoFullDelete    `json:"fullDelete"`
	Reactions     XInfoReactions     `json:"reactions"`
	TimedMessages XInfoTimedMessages `json:"timedMessages"`
	Voice         XInfoVoice         `json:"voice"`
}

type XInfoCalls struct {
	Allow string `json:"allow"`
}

type XInfoFullDelete struct {
	Allow string `json:"allow"`
}

type XInfoReactions struct {
	Allow string `json:"allow"`
}

type XInfoTimedMessages struct {
	Allow string `json:"allow"`
}

type XInfoVoice struct {
	Allow string `json:"allow"`
}

type XInfo struct {
	V      string `json:"v"`
	Event  string `json:"event"`
	Params XInfoParams
}

func (i *XInfo) ToString() string {
	jsonData, err := json.Marshal(i)
	if err != nil {
		jsonData = []byte("{}")
	}
	return string(jsonData)
}

func XInfoFromString(x string) XInfo {
	i := NewDefaultXInfo()
	err := json.Unmarshal([]byte(x), &i)
	if err != nil {
		fmt.Println("ERROR UNMARSHING JSON: ", err)
	}
	return i
}

func NewDefaultXInfo() XInfo {
	return XInfo{
		V:     "1-4",
		Event: "x.info",
		Params: XInfoParams{
			Profile: XInfoProfile{
				ContactLink: "",
				DisplayName: "GosaClient",
				FullName:    "Gosa Client",
				Image:       "",
				Preferences: XInfoPreferences{
					Calls: XInfoCalls{
						Allow: "no",
					},
					FullDelete: XInfoFullDelete{
						Allow: "no",
					},
					Reactions: XInfoReactions{
						Allow: "yes",
					},
					TimedMessages: XInfoTimedMessages{
						Allow: "yes",
					},
					Voice: XInfoVoice{
						Allow: "no",
					},
				},
			},
		},
	}
}
