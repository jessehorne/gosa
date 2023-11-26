package structs

type XInfo struct {
	V      string `json:"v"`
	Event  string `json:"event"`
	Params struct {
		Profile struct {
			ContactLink string `json:"contactLink"`
			DisplayName string `json:"displayName"`
			FullName    string `json:"fullName"`
			Image       string `json:"image"`
			Preferences struct {
				Calls struct {
					Allow string `json:"allow"`
				} `json:"calls"`
				FullDelete struct {
					Allow string `json:"allow"`
				} `json:"fullDelete"`
				Reactions struct {
					Allow string `json:"allow"`
				} `json:"reactions"`
				TimedMessages struct {
					Allow string `json:"allow"`
				} `json:"timedMessages"`
				Voice struct {
					Allow string `json:"allow"`
				} `json:"voice"`
			} `json:"preferences"`
		} `json:"profile"`
	} `json:"params"`
}
