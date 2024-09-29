package service

import (
	"regexp"

	"github.com/guhkun13/gmtg/config"
)

type QuestionImpl struct {
	Regexps *config.Regexps
}

func NewQuestionImpl(regexps *config.Regexps) QuestionIface {
	return &QuestionImpl{
		Regexps: regexps,
	}
}

func (s *QuestionImpl) IsMatchHowMuchQuestion(text string) bool {
	return regexp.MustCompile(s.Regexps.HowMuchQuestion).Match([]byte(text))
}

func (s *QuestionImpl) IsMatchHowManyCreditQuestion(text string) bool {
	return regexp.MustCompile(s.Regexps.HowManyCreditQuestion).Match([]byte(text))
}

func (s *QuestionImpl) IsMatchCreditComparisonQuestion(text string) bool {
	return regexp.MustCompile(s.Regexps.CreditComparisonQuestion).Match([]byte(text))
}

func (s *QuestionImpl) IsMatchCurrencyComparisonQuestion(text string) bool {
	return regexp.MustCompile(s.Regexps.CurrencyComparisonQuestion).Match([]byte(text))
}
