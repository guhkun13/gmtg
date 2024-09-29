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

func (s *QuestionImpl) AnswerHowMuchQuestion(text string) (string, error) {
	values := regexp.MustCompile(s.Regexps.HowMuchQuestion).FindStringSubmatch(text)
	currency := utils.TrimRight(values[1])
	romanNum, err := s.CurrencyService.GetValue(currency)
	if err != nil {
		log.Error().Err(err).Msg("CurrencyService.GetValue failed")
		return "", err
	}

	answer := fmt.Sprintf("%s is %d", currency, romanNum.Value())

	return answer, nil
}

func (s *QuestionImpl) AnswerHowManyCreditQuestion(text string) (string, error) {
	values := regexp.MustCompile(s.Regexps.HowManyCreditQuestion).FindStringSubmatch(text)

	currencies := utils.TrimRight(values[1])
	mineralName := utils.TrimRight(values[2])

	mineral, err := s.MineralService.GetValue(currencies, mineralName)
	if err != nil {
		log.Error().Err(err).Msg("MineralService.GetValue failed")
		return "", err
	}

	currencyMineral := utils.CombineString(currencies, mineralName)
	answer := fmt.Sprintf("%s is %d Credits", currencyMineral, int(mineral.Credit))

	return answer, nil
}

func (s *QuestionImpl) AnswerCreditComparisonQuestion(text string) (string, error) {
	values := regexp.MustCompile(s.Regexps.CreditComparisonQuestion).FindStringSubmatch(text)

	leftCurrency := utils.TrimRight(values[1])
	leftMineralName := utils.TrimRight(values[2])
	rightCurrency := utils.TrimRight(values[4])
	rightMineralName := utils.TrimRight(values[5])

	// fmt.Println("leftCurrency =", leftCurrency)
	// fmt.Println("leftMineralName =", leftMineralName)
	// fmt.Println("rightCurrency =", rightCurrency)
	// fmt.Println("rightMineralName =", rightMineralName)

	leftMineral, err := s.MineralService.GetValue(leftCurrency, leftMineralName)
	if err != nil {
		log.Error().Err(err).Msg("MineralService.GetValue left side failed")
		return "", err
	}

	rightMineral, err := s.MineralService.GetValue(rightCurrency, rightMineralName)
	if err != nil {
		log.Error().Err(err).Msg("MineralService.GetValue right side failed")
		return "", err
	}

	leftCurrencyMineral := utils.CombineString(leftCurrency, leftMineralName)
	rightCurrencyMineral := utils.CombineString(rightCurrency, rightMineralName)

	comparator := "less"
	if leftMineral.Credit > rightMineral.Credit {
		comparator = "more"
	}

	answer := fmt.Sprintf("%s has %s Credits than %s", leftCurrencyMineral, comparator, rightCurrencyMineral)

	return answer, nil
}
