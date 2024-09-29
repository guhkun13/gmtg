package service

import (
	"github.com/guhkun13/gmtg/config"
	"github.com/guhkun13/gmtg/libs"
)

type Services struct {
	Currency CurrencyIface
	Mineral  MineralIface
	Question QuestionIface
	Answer   AnswerIface
}

func NewServices(regexps *config.Regexps) Services {
	currencyService := NewCurrencyImpl(regexps.AssignCurrency)
	mineralService := NewMineralImpl(regexps.AssignMineral, currencyService)

	return Services{
		Currency: currencyService,
		Mineral:  mineralService,
		Question: NewQuestionImpl(regexps),
		Answer:   NewAnswerImpl(regexps, currencyService, mineralService),
	}
}

func (s *Services) EvaluateText(text string) string {
	var err error
	answer := ""
	isQuestion := true

	if s.Currency.IsMatchAssignValue(text) {
		isQuestion = false
		err = s.Currency.AssignValue(text)
	} else if s.Mineral.IsMatchAssignValue(text) {
		isQuestion = false
		err = s.Mineral.AssignValue(text)
	} else if s.Question.IsMatchHowMuchQuestion(text) {
		answer, err = s.Answer.AnswerHowMuchQuestion(text)
	} else if s.Question.IsMatchHowManyCreditQuestion(text) {
		answer, err = s.Answer.AnswerHowManyCreditQuestion(text)
	} else if s.Question.IsMatchCreditComparisonQuestion(text) {
		answer, err = s.Answer.AnswerCreditComparisonQuestion(text)
	} else if s.Question.IsMatchCurrencyComparisonQuestion(text) {
		answer, err = s.Answer.AnswerCurrencyComparisonQuestion(text)
	} else {
		err = libs.ErrUnrecognizedText
	}

	if isQuestion {
		if err != nil {
			answer = err.Error()
		}
	}

	return answer

}
