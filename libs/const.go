package libs

const (
	regexRomanChar string = "([IVXLCDM])"
	regexCurrency  string = `((?:[a-z]+\s?)+)`
	regexMineral   string = `([A-Z][a-z]+\s?)`
)

const (
	FileInputName  string = "input.txt"
	FileOutputName string = "output.txt"
)

const (
	CommandCurrency string = "command_currency"
	CommandMineral  string = "command_mineral"
)

const (
	QuestionHowMuchCurrency    string = "question_how_much_currency"
	QuestionHowManyCredit      string = "question_how_many_credit"
	QuestionCreditComparison   string = "question_credit_comparison"
	QuestionCurrencyComparison string = "question_currency_comparison"
)
