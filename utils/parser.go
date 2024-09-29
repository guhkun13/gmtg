package utils

import (
	"regexp"

	"github.com/guhkun13/gmtg/libs"
)

func IsMatchNewCurrency(text string) bool {
	return regexp.MustCompile(libs.RegexNewCurrency).Match([]byte(text))
}

func IsMatchNewMineral(text string) bool {
	return regexp.MustCompile(libs.RegexNewMineral).Match([]byte(text))
}

func IsMatchHowMuchQuestion(text string) bool {
	return regexp.MustCompile(libs.RegexHowMuchQuestion).Match([]byte(text))
}

func IsMatchHowManyCreditQuestion(text string) bool {
	return regexp.MustCompile(libs.RegexHowManyCreditQuestion).Match([]byte(text))
}

func IsMatchCreditComparisonQuestion(text string) bool {
	return regexp.MustCompile(libs.RegexCreditComparisonQuestion).Match([]byte(text))
}

func IsMatchCurrencyComparisonQuestion(text string) bool {
	return regexp.MustCompile(libs.RegexCurrencyComparisonQuestion).Match([]byte(text))
}
