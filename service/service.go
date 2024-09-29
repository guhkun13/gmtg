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
	answer := ""
	isQuestion := true

	if s.Currency.IsMatchAssignValue(text) {
		isQuestion = false
		err = s.Currency.AssignValue(text)
	} else if s.Mineral.IsMatchAssignValue(text) {
		isQuestion = false
		err = s.Mineral.AssignValue(text)
	} else if s.Question.IsMatchHowMuchQuestion(text) {
		answer, err = s.Question.AnswerHowMuchQuestion(text)
	} else if s.Question.IsMatchHowManyCreditQuestion(text) {
		answer, err = s.Question.AnswerHowManyCreditQuestion(text)
	} else {
		err = libs.ErrUnrecognizedText
	}

	if isQuestion {
		if err != nil {
			answer = err.Error()
		}
		log.Info().
			Str("0-question", text).
			Str("1-answer", answer).
			Msg("result")
	}

}

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
