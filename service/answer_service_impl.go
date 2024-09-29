package service

import (
	"fmt"
	"regexp"

	"github.com/guhkun13/gmtg/config"
	"github.com/guhkun13/gmtg/utils"
	"github.com/rs/zerolog/log"
)

type AnswerImpl struct {
	Regexps         *config.Regexps
	CurrencyService CurrencyIface
	MineralService  MineralIface
}

func NewAnswerImpl(regexps *config.Regexps,
	currencyService CurrencyIface,
	mineralService MineralIface,
) AnswerIface {
	return &AnswerImpl{
		Regexps:         regexps,
		CurrencyService: currencyService,
		MineralService:  mineralService,
	}
}

func (s *AnswerImpl) AnswerHowMuchQuestion(text string) (string, error) {
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

func (s *AnswerImpl) AnswerHowManyCreditQuestion(text string) (string, error) {
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

func (s *AnswerImpl) AnswerCreditComparisonQuestion(text string) (string, error) {
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

func (s *AnswerImpl) AnswerCurrencyComparisonQuestion(text string) (string, error) {
	values := regexp.MustCompile(s.Regexps.CurrencyComparisonQuestion).FindStringSubmatch(text)

	leftCurrency := utils.TrimRight(values[1])
	rightCurrency := utils.TrimRight(values[3])

	leftRomanNum, err := s.CurrencyService.GetValue(leftCurrency)
	if err != nil {
		log.Error().Err(err).Msg("CurrencyService.GetValue failed")
		return "", err
	}

	rightRomanNum, err := s.CurrencyService.GetValue(rightCurrency)
	if err != nil {
		log.Error().Err(err).Msg("CurrencyService.GetValue failed")
		return "", err
	}

	fmt.Println("leftCurrency =", leftCurrency)
	fmt.Println("leftRomanNum =", leftRomanNum)

	fmt.Println("rightCurrency =", rightCurrency)
	fmt.Println("rightRomanNum =", rightRomanNum)

	comparator := "smaller"
	if leftRomanNum.Value() > rightRomanNum.Value() {
		comparator = "larger"
	}

	answer := fmt.Sprintf("%s is %s than %s", leftCurrency, comparator, rightCurrency)

	return answer, nil
}
