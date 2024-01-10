package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func IsEmptyOrWhitespace(s string) bool {
	return strings.TrimSpace(s) == ""
}

func GetOrderID() string {
	return uuid.NewString()
}

func GetNewOrderName(uid string) string {

	formattedDateTime := time.Now().Format("2006-01-02_15-04-05")

	return fmt.Sprintf("%s%s.asc", formattedDateTime, uid)
}
