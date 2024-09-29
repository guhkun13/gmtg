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
		log.Error().Err(err).Msg("CurrencyService.GetValue failed")
		return err
	}

	answer := fmt.Sprintf("%s is %d", currency, romanNum.Value())

	log.Info().
		Str("0-question", text).
		Str("1-answer", answer).
		Msg("result")
	// WriteToOutput(answer)
	return nil
}

func (s *QuestionImpl) AnswerHowManyCreditQuestion(text string) error {
	values := regexp.MustCompile(s.Regexps.HowManyCreditQuestion).FindStringSubmatch(text)

	currencies := utils.TrimRight(values[1])
	mineralName := utils.TrimRight(values[2])

	mineral, err := s.MineralService.GetValue(currencies, mineralName)
	if err != nil {
		log.Error().Err(err).Msg("MineralService.GetValue failed")
		return err
	}

	currencyMineral := utils.CombineString(currencies, mineralName)
	answer := fmt.Sprintf("%s is %d Credits", currencyMineral, int(mineral.Credit))

	log.Info().
		Str("0-question", text).
		Str("1-answer", answer).
		Msg("result")

	return nil
}
