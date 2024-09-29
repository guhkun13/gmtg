package utils

import (
	"regexp"

	"github.com/guhkun13/gmtg/config"
)

func IsMatchCreditComparisonQuestion(text string) bool {
	return regexp.MustCompile(config.RegexCreditComparisonQuestion).Match([]byte(text))
}

func IsMatchCurrencyComparisonQuestion(text string) bool {
	return regexp.MustCompile(config.RegexCurrencyComparisonQuestion).Match([]byte(text))
}
