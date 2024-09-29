package service

import (
	"github.com/guhkun13/gmtg/config"
	"github.com/guhkun13/gmtg/libs"
	"github.com/rs/zerolog/log"
)

type Services struct {
	Currency CurrencyIface
	Mineral  MineralIface
	Question QuestionIface
}

func NewServices(regexps *config.Regexps) Services {
	currencyService := NewCurrencyImpl(regexps.AssignCurrency)
	mineralService := NewMineralImpl(regexps.AssignMineral, currencyService)

	return Services{
		Currency: currencyService,
		Mineral:  mineralService,
		Question: NewQuestionImpl(regexps, currencyService, mineralService),
	}
}

func (s *Services) EvaluateText(text string) {
	var err error

	if s.Currency.IsMatchAssignValue(text) {
		err = s.Currency.AssignValue(text)
	} else if s.Mineral.IsMatchAssignValue(text) {
		err = s.Mineral.AssignValue(text)
	} else if s.Question.IsMatchHowMuchQuestion(text) {
		err = s.Question.AnswerHowMuchQuestion(text)
	} else if s.Question.IsMatchHowManyCreditQuestion(text) {
		err = s.Question.AnswerHowManyCreditQuestion(text)
	} else {
		err = libs.ErrUnrecognizedText
	}

	if err != nil {
		log.Error().Err(err).Msg("Error occured")
	}
}

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
