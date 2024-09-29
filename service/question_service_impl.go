package service

import (
	"fmt"
	"regexp"

	"github.com/guhkun13/gmtg/config"
	"github.com/guhkun13/gmtg/utils"
	"github.com/rs/zerolog/log"
)

type Question struct {
	Value  string
	Answer string
}

type QuestionImpl struct {
	Regexps         *config.Regexps
	CurrencyService CurrencyIface
	MineralService  MineralIface
}

func NewQuestionImpl(regexps *config.Regexps,
	currencyService CurrencyIface,
	mineralService MineralIface,
) QuestionIface {
	return &QuestionImpl{
		Regexps:         regexps,
		CurrencyService: currencyService,
		MineralService:  mineralService,
	}
}

func (s *QuestionImpl) IsMatchHowMuchQuestion(text string) bool {
	return regexp.MustCompile(s.Regexps.HowMuchQuestion).Match([]byte(text))
}

func (s *QuestionImpl) IsMatchHowManyCreditQuestion(text string) bool {
	return regexp.MustCompile(s.Regexps.HowManyCreditQuestion).Match([]byte(text))
}

func (s *QuestionImpl) AnswerHowMuchQuestion(text string) error {
	values := regexp.MustCompile(s.Regexps.HowMuchQuestion).FindStringSubmatch(text)

	currency := utils.TrimRight(values[1])

	romanNum, err := s.CurrencyService.GetValue(currency)
	if err != nil {
		log.Error().Err(err).Msg("ConvertNewCurrencyToRoman failed")
		return err
	}

	answer := fmt.Sprintf("%s is %d", currency, romanNum.Value())

	fmt.Println(answer)
	// WriteToOutput(answer)
	return nil
}
