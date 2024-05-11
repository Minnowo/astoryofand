package util

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func IsEmptyOrWhitespace(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func BytesIsEmptyOrWhitespace(b []byte) bool {
	return len(bytes.TrimSpace(b)) == 0
}

func GetOrderID() string {
	return uuid.NewString()
}

func GetNewOrderName(uid string) string {

	formattedDateTime := time.Now().Format("2006-01-02_15-04-05")

	return fmt.Sprintf("%s%s.asc", formattedDateTime, uid)
}
