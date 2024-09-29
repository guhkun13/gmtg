package utils

import (
	"fmt"
	"strings"
)

func TrimRight(text string) string {
	return strings.TrimRight(text, " ")
}

func CombineString(currency, mineral string) string {
	return fmt.Sprintf("%s %s", currency, mineral)
}
