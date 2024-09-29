package service

type QuestionIface interface {
	IsMatchHowMuchQuestion(text string) bool
	IsMatchHowManyCreditQuestion(text string) bool
	IsMatchCreditComparisonQuestion(text string) bool
	IsMatchCurrencyComparisonQuestion(text string) bool
}
