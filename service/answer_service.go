package service

type AnswerIface interface {
	AnswerHowMuchQuestion(text string) (string, error)
	AnswerHowManyCreditQuestion(text string) (string, error)
	AnswerCreditComparisonQuestion(text string) (string, error)
	AnswerCurrencyComparisonQuestion(text string) (string, error)
}
