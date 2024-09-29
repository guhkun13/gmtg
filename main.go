package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/dhowden/numerus"
	"github.com/guhkun13/gmtg/config"
	"github.com/guhkun13/gmtg/file"
	"github.com/guhkun13/gmtg/libs"
	questionfactories "github.com/guhkun13/gmtg/question_factories"
	"github.com/guhkun13/gmtg/utils"
	log "github.com/rs/zerolog/log"
)

var questionFactories = make(map[string]questionfactories.QuestionFactory)

func main() {
	config.InitLogger()

	// commandFactories := map[string]commandfactories.CommandFactory{
	// 	libs.CommandCurrency: &commandfactories.CurrencyImpl{},
	// 	libs.CommandMineral:  &commandfactories.MineralImpl{},
	// }

	// questionFactories = map[string]questionfactories.QuestionFactory{
	// 	libs.QuestionHowMuchCurrency: questionfactories.NewHowManyQuestionImpl(libs.RegexHowManyCreditQuestion),
	// }

	fileReader := file.NewReader()
	libs.EmptyFileContent(libs.FileOutputName)

	fileInput := fileReader.ReadFile(libs.FileInputName)
	scanner := bufio.NewScanner(fileInput)

	idx := 0
	st := time.Now()
	for scanner.Scan() {
		txt := scanner.Text()
		evaluateText(txt)
		idx++
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal().Err(err).Msg("scannerfailed  ")
	}

	et := time.Now()
	fmt.Println("[Total execution time] : ", et.Sub(st))
}

func evaluateText(text string) {

	if utils.IsMatchNewCurrency(text) {
		AssignNewCurrency(text)
	} else if utils.IsMatchNewMineral(text) {
		AssignNewMineral(text)
	} else if utils.IsMatchHowMuchQuestion(text) {
		factoryType := libs.QuestionHowManyCredit
		questionFactory := questionFactories[factoryType]
		questionFactory.Answer(text)

	} else if utils.IsMatchHowManyCreditQuestion(text) {
		AnswerHowManyCreditQuestion(text)
	} else if utils.IsMatchCreditComparisonQuestion(text) {
		AnswerCreditComparisonQuestion(text)
	} else if utils.IsMatchCurrencyComparisonQuestion(text) {
		AnswerCurrencyComparisonQuestion(text)
	} else {
		fmt.Println(libs.ErrQuestionUnrecognized)
		utils.WriteToOutput(libs.ErrQuestionUnrecognized)
	}
}

// assigning value to Mineral
func AssignNewMineral(text string) {
	// log.Debug().Msg("AssignNewMineral")

	values := regexp.MustCompile(regexNewMineral).FindStringSubmatch(text)
	// log.Debug().Strs("values", values).Msg("FindStringSubmatch")

	currency := trimRight(values[1])
	mineral := trimRight(values[2])
	strTotalAmount := trimRight(values[3])

	// fmt.Println("currency", currency)
	// fmt.Println("mineral", mineral)
	// fmt.Println("strTotalAmount", strTotalAmount)

	romanStr, err := ConvertNewCurrencyToRoman(currency)
	if err != nil {
		log.Error().Err(err).Msg("ConvertNewCurrencyToRoman failed")
		return
	}

	// from romanValue we need to make sure is it really ROMAN numeral or not
	romanNum, err := numerus.Parse(romanStr)
	if err != nil {
		log.Error().Err(err).Msg("numerus.Parse failed")
	}

	// fmt.Println("intVal", intVal.Value())
	intTotalAmount, err := strconv.Atoi(strTotalAmount)
	if err != nil {
		log.Error().Err(err).Msg("Credit value must be number")
		return
	}

	// fmt.Println("intTotalAmount", intTotalAmount)
	// fmt.Println("romanNum.Value()", romanNum.Value())

	creditValue := float64(float64(intTotalAmount) / float64(romanNum.Value()))
	// fmt.Println("creditValue", creditValue)

	newMineralsMap[mineral] = creditValue

	// log.Debug().
	// 	Interface("value", newMineralsMap).
	// 	Msg("newMineralsMap")

}

func AnswerHowMuchQuestion(text string) {
	// log.Debug().Msg("AnswerHowMuchQuestion")

	values := regexp.MustCompile(regexHowMuchQuestion).FindStringSubmatch(text)
	// log.Debug().Strs("values", values).Msg("FindStringSubmatch")

	currency := trimRight(values[1])

	romanStr, err := ConvertNewCurrencyToRoman(currency)
	if err != nil {
		log.Error().Err(err).Msg("ConvertNewCurrencyToRoman failed")
	}

	// fmt.Println("romanStr", romanStr)

	romanNum, err := numerus.Parse(romanStr)
	if err != nil {
		log.Error().Err(err).Msg("numerus.Parse failed")
	}

	// fmt.Println("intValue", romanNum.Value())

	answer := fmt.Sprintf("%s is %d", currency, romanNum.Value())

	fmt.Println(answer)
	WriteToOutput(answer)
}

func AnswerCreditComparisonQuestion(text string) {
	// log.Debug().Msg("AnswerCreditComparisonQuestion")

	pattern := regexp.MustCompile(regexCreditComparisonQuestion)
	values := pattern.FindStringSubmatch(text)
	// log.Debug().Strs("values", values).Msg("FindStringSubmatch")

	leftCurrency := trimRight(values[1])
	leftMineral := trimRight(values[2])
	rightCurrency := trimRight(values[4])
	rightMineral := trimRight(values[5])

	// fmt.Println("leftCurrency =", leftCurrency)
	// fmt.Println("leftMineral =", leftMineral)
	// fmt.Println("rightCurrency =", rightCurrency)
	// fmt.Println("rightMineral =", rightMineral)

	// calculate left value
	leftCreditValue, err := getMineralValue(leftCurrency, leftMineral)
	if err != nil {
		log.Error().Err(err).Msg("getMineralValue for left side failed")
		return
	}

	// calculate right value
	rightCreditValue, err := getMineralValue(rightCurrency, rightMineral)
	if err != nil {
		log.Error().Err(err).Msg("getMineralValue for right side failed")
		return
	}
	// fmt.Println("leftCreditValue", leftCreditValue)
	// fmt.Println("rightCreditValue", rightCreditValue)

	leftCurrencyMineral := combineString(leftCurrency, leftMineral)
	rightCurrencyMineral := combineString(rightCurrency, rightMineral)
	comparator := "less"
	if leftCreditValue > rightCreditValue {
		comparator = "more"
	}

	answer := fmt.Sprintf("%s has %s Credits than %s", leftCurrencyMineral, comparator, rightCurrencyMineral)
	fmt.Println(answer)
	WriteToOutput(answer)
}

func AnswerCurrencyComparisonQuestion(text string) {
	// log.Debug().Msg("AnswerCurrencyComparisonQuestion")

	values := regexp.MustCompile(regexCurrencyComparisonQuestion).FindStringSubmatch(text)
	// log.Debug().Strs("values", values).Msg("FindStringSubmatch")

	leftCurrency := trimRight(values[1])
	rightCurrency := trimRight(values[3])

	leftCurrencyValue, err := getCurrencyValue(leftCurrency)
	if err != nil {
		log.Error().Err(err).Msg("getCurrencyValue failed")
		return
	}

	rightCurrencyValue, err := getCurrencyValue(rightCurrency)
	if err != nil {
		log.Error().Err(err).Msg("getCurrencyValue failed")
		return
	}

	// fmt.Println("leftCurrency =", leftCurrency)
	// fmt.Println("leftCurrencyValue =", leftCurrencyValue)

	// fmt.Println("rightCurrency =", rightCurrency)
	// fmt.Println("rightCurrencyValue =", rightCurrencyValue)

	comparator := "smaller"
	if leftCurrencyValue > rightCurrencyValue {
		comparator = "larger"
	}

	answer := fmt.Sprintf("%s is %s than %s", leftCurrency, comparator, rightCurrency)
	fmt.Println(answer)
	WriteToOutput(answer)
}
