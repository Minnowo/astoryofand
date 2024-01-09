package model

import "strings"

func IsEmptyOrWhitespace(s string) bool {
	return strings.TrimSpace(s) == ""
}
