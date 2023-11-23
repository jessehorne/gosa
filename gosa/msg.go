package gosa

import (
	"strings"
)

// returns either the invitation URI or "" if there was an error
func parseForINV(data string) string {
	parsed := strings.Split(data, " ")

	if len(parsed) < 2 {
		return ""
	}

	if parsed[0] == "INV" {
		return parsed[1]
	}

	return ""
}
