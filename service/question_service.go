package service

type QuestionIface interface {
	IsMatchHowMuchQuestion(text string) bool
	IsMatchHowManyCreditQuestion(text string) bool
	AnswerHowMuchQuestion(text string) (string, error)
	AnswerHowManyCreditQuestion(text string) (string, error)
}
