package config

import "fmt"

const (
	regexRomanChar string = "([IVXLCDM])"
	regexCurrency  string = `((?:[a-z]+\s?)+)`
	regexMineral   string = `([A-Z][a-z]+\s?)`
)

var (
	RegexAssignCurrency             string = fmt.Sprintf(`^%s is %s`, regexCurrency, regexRomanChar)
	RegexAssignMineral              string = fmt.Sprintf(`^%s%s is (\d+) Credits$`, regexCurrency, regexMineral)
	RegexHowMuchQuestion            string = fmt.Sprintf(`^[H|h]ow much is %s\?$`, regexCurrency)
	RegexHowManyCreditQuestion      string = fmt.Sprintf(`^[H|h]ow many Credits is %s%s\?$`, regexCurrency, regexMineral)
	RegexCreditComparisonQuestion   string = fmt.Sprintf(`^[D|d]oes %s%s has (less|more) Credits than %s%s\?`, regexCurrency, regexMineral, regexCurrency, regexMineral)
	RegexCurrencyComparisonQuestion string = fmt.Sprintf(`^Is %s (larger|smaller) than %s\?$`, regexCurrency, regexCurrency)
)

type Regexps struct {
	AssignCurrency             string
	AssignMineral              string
	HowMuchQuestion            string
	HowManyCreditQuestion      string
	CreditComparisonQuestion   string
	CurrencyComparisonQuestion string
}

func InitRegexps() Regexps {
	return Regexps{
		AssignCurrency:             RegexAssignCurrency,
		AssignMineral:              RegexAssignMineral,
		HowMuchQuestion:            RegexHowMuchQuestion,
		HowManyCreditQuestion:      RegexHowManyCreditQuestion,
		CreditComparisonQuestion:   RegexCreditComparisonQuestion,
		CurrencyComparisonQuestion: RegexCurrencyComparisonQuestion,
	}
}
