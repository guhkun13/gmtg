package utils

import (
	"regexp"

	"github.com/guhkun13/gmtg/config"
)

func IsMatchHowMuchQuestion(text string) bool {
	return regexp.MustCompile(config.RegexHowMuchQuestion).Match([]byte(text))
}

func IsMatchHowManyCreditQuestion(text string) bool {
	return regexp.MustCompile(config.RegexHowManyCreditQuestion).Match([]byte(text))
}

func IsMatchCreditComparisonQuestion(text string) bool {
	return regexp.MustCompile(config.RegexCreditComparisonQuestion).Match([]byte(text))
}

func IsMatchCurrencyComparisonQuestion(text string) bool {
	return regexp.MustCompile(config.RegexCurrencyComparisonQuestion).Match([]byte(text))
}
