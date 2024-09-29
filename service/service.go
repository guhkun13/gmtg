package service

import (
	"github.com/guhkun13/gmtg/config"
	"github.com/guhkun13/gmtg/libs"
	"github.com/rs/zerolog/log"
)

type Services struct {
	Currency CurrencyIface
	Mineral  MineralIface
}

func NewServices(regexps config.Regexps) Services {
	currencyService := NewCurrencyImpl(regexps.AssignCurrency)

	return Services{
		Currency: currencyService,
		Mineral:  NewMineralImpl(regexps.AssignMineral, currencyService),
	}
}

func (s *Services) EvaluateText(text string) {
	var err error

	if s.Currency.IsMatchAssignValue(text) {
		err = s.Currency.AssignValue(text)
		if err != nil {
			log.Error().Err(err).Msg("Currency.AssignValue failed")
		}
	} else if s.Mineral.IsMatchAssignValue(text) {
		err = s.Mineral.AssignValue(text)
		if err != nil {
			log.Error().Err(err).Msg("Mineral.AssignValue failed")
		}
	} else {
		err = libs.ErrUnrecognizedText
		log.Error().Err(err).Msg("Err happened")
	}
}

// func evaluateText(text string) {

// 	if text {
// 		AssignNewCurrency(text)
// 	} else if utils.IsMatchNewMineral(text) {
// 		AssignNewMineral(text)
// 	} else if utils.IsMatchHowMuchQuestion(text) {
// 		AnswerHowMuchQuestion(text)
// 	}

// 	// else if utils.IsMatchHowManyCreditQuestion(text) {
// 	// 	AnswerHowManyCreditQuestion(text)
// 	// } else if utils.IsMatchCreditComparisonQuestion(text) {
// 	// 	AnswerCreditComparisonQuestion(text)
// 	// } else if utils.IsMatchCurrencyComparisonQuestion(text) {
// 	// 	AnswerCurrencyComparisonQuestion(text)
// 	// } else {
// 	// 	fmt.Println(libs.ErrQuestionUnrecognized)
// 	// 	utils.WriteToOutput(libs.ErrQuestionUnrecognized.Error())
// 	// }
// }

// func AnswerHowMuchQuestion(text string) {
// 	// log.Debug().Msg("AnswerHowMuchQuestion")

// 	values := regexp.MustCompile(regexHowMuchQuestion).FindStringSubmatch(text)
// 	// log.Debug().Strs("values", values).Msg("FindStringSubmatch")

// 	currency := trimRight(values[1])

// 	romanStr, err := ConvertNewCurrencyToRoman(currency)
// 	if err != nil {
// 		log.Error().Err(err).Msg("ConvertNewCurrencyToRoman failed")
// 	}

// 	// fmt.Println("romanStr", romanStr)

// 	romanNum, err := numerus.Parse(romanStr)
// 	if err != nil {
// 		log.Error().Err(err).Msg("numerus.Parse failed")
// 	}

// 	// fmt.Println("intValue", romanNum.Value())

// 	answer := fmt.Sprintf("%s is %d", currency, romanNum.Value())

// 	fmt.Println(answer)
// 	WriteToOutput(answer)
// }

// func AnswerCreditComparisonQuestion(text string) {
// 	// log.Debug().Msg("AnswerCreditComparisonQuestion")

// 	pattern := regexp.MustCompile(regexCreditComparisonQuestion)
// 	values := pattern.FindStringSubmatch(text)
// 	// log.Debug().Strs("values", values).Msg("FindStringSubmatch")

// 	leftCurrency := trimRight(values[1])
// 	leftMineral := trimRight(values[2])
// 	rightCurrency := trimRight(values[4])
// 	rightMineral := trimRight(values[5])

// 	// fmt.Println("leftCurrency =", leftCurrency)
// 	// fmt.Println("leftMineral =", leftMineral)
// 	// fmt.Println("rightCurrency =", rightCurrency)
// 	// fmt.Println("rightMineral =", rightMineral)

// 	// calculate left value
// 	leftCreditValue, err := getMineralValue(leftCurrency, leftMineral)
// 	if err != nil {
// 		log.Error().Err(err).Msg("getMineralValue for left side failed")
// 		return
// 	}

// 	// calculate right value
// 	rightCreditValue, err := getMineralValue(rightCurrency, rightMineral)
// 	if err != nil {
// 		log.Error().Err(err).Msg("getMineralValue for right side failed")
// 		return
// 	}
// 	// fmt.Println("leftCreditValue", leftCreditValue)
// 	// fmt.Println("rightCreditValue", rightCreditValue)

// 	leftCurrencyMineral := combineString(leftCurrency, leftMineral)
// 	rightCurrencyMineral := combineString(rightCurrency, rightMineral)
// 	comparator := "less"
// 	if leftCreditValue > rightCreditValue {
// 		comparator = "more"
// 	}

// 	answer := fmt.Sprintf("%s has %s Credits than %s", leftCurrencyMineral, comparator, rightCurrencyMineral)
// 	fmt.Println(answer)
// 	WriteToOutput(answer)
// }

// func AnswerCurrencyComparisonQuestion(text string) {
// 	// log.Debug().Msg("AnswerCurrencyComparisonQuestion")

// 	values := regexp.MustCompile(regexCurrencyComparisonQuestion).FindStringSubmatch(text)
// 	// log.Debug().Strs("values", values).Msg("FindStringSubmatch")

// 	leftCurrency := trimRight(values[1])
// 	rightCurrency := trimRight(values[3])

// 	leftCurrencyValue, err := getCurrencyValue(leftCurrency)
// 	if err != nil {
// 		log.Error().Err(err).Msg("getCurrencyValue failed")
// 		return
// 	}

// 	rightCurrencyValue, err := getCurrencyValue(rightCurrency)
// 	if err != nil {
// 		log.Error().Err(err).Msg("getCurrencyValue failed")
// 		return
// 	}

// 	// fmt.Println("leftCurrency =", leftCurrency)
// 	// fmt.Println("leftCurrencyValue =", leftCurrencyValue)

// 	// fmt.Println("rightCurrency =", rightCurrency)
// 	// fmt.Println("rightCurrencyValue =", rightCurrencyValue)

// 	comparator := "smaller"
// 	if leftCurrencyValue > rightCurrencyValue {
// 		comparator = "larger"
// 	}

// 	answer := fmt.Sprintf("%s is %s than %s", leftCurrency, comparator, rightCurrency)
// 	fmt.Println(answer)
// 	WriteToOutput(answer)
// }

// func (f *HowManyQuestionImpl) Answer(text string) {
// 	values := regexp.MustCompile(f.Regex).FindStringSubmatch(text)
// 	// log.Debug().Strs("values", values).Msg("FindStringSubmatch")

// 	currency := utils.TrimRight(values[1])
// 	mineral := utils.TrimRight(values[2])

// 	creditValue, err := utils.GetMineralValue(currency, mineral)
// 	if err != nil {
// 		// log.Error().Err(err).Msg("getMineralValue failed")
// 		return
// 	}

// 	currencyMineral := combineString(currency, mineral)
// 	answer := fmt.Sprintf("%s is %d Credits", currencyMineral, int(creditValue))

// 	fmt.Println(answer)
// 	WriteToOutput(answer)

// }
