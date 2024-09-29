package libs

import "fmt"

var (
	RegexNewCurrency                string = fmt.Sprintf(`^%s is %s`, regexCurrency, regexRomanChar)
	RegexNewMineral                 string = fmt.Sprintf(`^%s%s is (\d+) Credits$`, regexCurrency, regexMineral)
	RegexHowMuchQuestion            string = fmt.Sprintf(`^[H|h]ow much is %s\?$`, regexCurrency)
	RegexHowManyCreditQuestion      string = fmt.Sprintf(`^[H|h]ow many Credits is %s%s\?$`, regexCurrency, regexMineral)
	RegexCreditComparisonQuestion   string = fmt.Sprintf(`^[D|d]oes %s%s has (less|more) Credits than %s%s\?`, regexCurrency, regexMineral, regexCurrency, regexMineral)
	RegexCurrencyComparisonQuestion string = fmt.Sprintf(`^Is %s (larger|smaller) than %s\?$`, regexCurrency, regexCurrency)
)
