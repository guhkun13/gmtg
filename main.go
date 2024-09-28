package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dhowden/numerus"
	"github.com/guhkun13/gmtg/config"
	log "github.com/rs/zerolog/log"
)

const (
	regexRomanChar string = "([IVXLCDM])"
	regexCurrency  string = `((?:[a-z]+\s?)+)`
	regexMineral   string = `([A-Z][a-z]+\s?)`
)

var (
	regexNewCurrency                string = fmt.Sprintf(`^%s is %s`, regexCurrency, regexRomanChar)
	regexNewMineral                 string = fmt.Sprintf(`^%s%s is (\d+) Credits$`, regexCurrency, regexMineral)
	regexHowMuchQuestion            string = fmt.Sprintf(`^[H|h]ow much is %s\?$`, regexCurrency)
	regexHowManyCreditQuestion      string = fmt.Sprintf(`^[H|h]ow many Credits is %s%s\?$`, regexCurrency, regexMineral)
	regexCreditComparisonQuestion   string = fmt.Sprintf(`^[D|d]oes %s%s has (less|more) Credits than %s%s\?`, regexCurrency, regexMineral, regexCurrency, regexMineral)
	regexCurrencyComparisonQuestion string = fmt.Sprintf(`^Is %s (larger|smaller) than %s\?$`, regexCurrency, regexCurrency)
)

var newCurrenciesMap = make(map[string]string)
var newMineralsMap = make(map[string]float64)

func main() {
	fmt.Println(regexNewCurrency)
	fmt.Println(regexNewMineral)
	fmt.Println(regexHowMuchQuestion)
	fmt.Println(regexHowManyCreditQuestion)
	fmt.Println(regexCreditComparisonQuestion)

	config.InitLogger()

	// read files

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	idx := 0
	st := time.Now()
	fmt.Println(">> Start at ", st)

	for scanner.Scan() {
		txt := scanner.Text()
		log.Info().Str("val", txt).Msgf("[%d]", idx)
		evaluateText(txt)
		idx++
	}

	// wg := &sync.WaitGroup{}
	// for scanner.Scan() {
	// 	wg.Add(1)
	// 	txt := scanner.Text()
	// 	log.Info().Str("val", txt).Msgf("[%d]", idx)
	// 	// evaluateText(wg, txt)
	// 	go asyncEvaluateText(wg, txt)

	// 	idx++
	// }
	// wg.Wait()

	err = scanner.Err()
	if err != nil {
		log.Fatal().Err(err).Msg("scannerfailed  ")
	}

	log.Debug().
		Interface("0-newCurrenciesMap", newCurrenciesMap).
		Interface("1-newMineralsMap", newMineralsMap).
		Msg("Value")

	et := time.Now()
	fmt.Println(">> Finish at ", et)

	fmt.Println(">> Execution time ", et.Sub(st))
}

func evaluateText(text string) {

	if IsMatchNewCurrency(text) {
		AssignNewCurrency(text)
	} else if IsMatchNewMineral(text) {
		AssignNewMineral(text)
	} else if IsMatchHowMuchQuestion(text) {
		AnswerHowMuchQuestion(text)
	} else if IsMatchHowManyCreditQuestion(text) {
		AnswerHowManyCreditQuestion(text)
	} else if IsMatchCreditComparisonQuestion(text) {
		AnswerCreditComparisonQuestion(text)
	} else if IsMatchCurrencyComparisonQuestion(text) {
		AnswerCurrencyComparisonQuestion(text)
	}
}

func IsMatchNewCurrency(text string) bool {
	return regexp.MustCompile(regexNewCurrency).Match([]byte(text))
}

func IsMatchNewMineral(text string) bool {
	return regexp.MustCompile(regexNewMineral).Match([]byte(text))
}

func IsMatchHowMuchQuestion(text string) bool {
	return regexp.MustCompile(regexHowMuchQuestion).Match([]byte(text))
}

func IsMatchHowManyCreditQuestion(text string) bool {
	return regexp.MustCompile(regexHowManyCreditQuestion).Match([]byte(text))
}

func IsMatchCreditComparisonQuestion(text string) bool {
	return regexp.MustCompile(regexCreditComparisonQuestion).Match([]byte(text))
}

func IsMatchCurrencyComparisonQuestion(text string) bool {
	return regexp.MustCompile(regexCurrencyComparisonQuestion).Match([]byte(text))
}

// assigning new currency to roman value
func AssignNewCurrency(text string) {
	values := regexp.MustCompile(regexNewCurrency).FindStringSubmatch(text)
	currency := trimRight(values[1])
	roman := trimRight(values[2])
	newCurrenciesMap[currency] = roman
}

func ConvertNewCurrencyToRoman(text string) (result string, err error) {
	currencies := strings.Split(text, " ")
	if len(currencies) == 1 {
		return CurrencyToRoman(text)
	}

	finalValue := ""
	for _, c := range currencies {
		roman, err := CurrencyToRoman(c)
		if err != nil {
			return "", err
		}
		finalValue += roman
	}

	return finalValue, err
}

func CurrencyToRoman(currency string) (string, error) {
	// fmt.Println("CurrencyToRoman", currency)
	// fmt.Println("newCurrenciesMap", newCurrenciesMap)

	val, err := numerus.Parse(newCurrenciesMap[currency])
	if err != nil {
		return "", err
	}

	return val.String(), nil
}

// assigning value to Mineral
func AssignNewMineral(text string) {
	log.Debug().Msg("AssignNewMineral")

	values := regexp.MustCompile(regexNewMineral).FindStringSubmatch(text)
	log.Debug().Strs("values", values).Msg("FindStringSubmatch")

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
	log.Debug().Msg("AnswerHowMuchQuestion")

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

	fmt.Println("answer", answer)

}

func AnswerHowManyCreditQuestion(text string) {
	// log.Debug().Msg("AnswerHowManyCreditQuestion")

	values := regexp.MustCompile(regexHowManyCreditQuestion).FindStringSubmatch(text)
	// log.Debug().Strs("values", values).Msg("FindStringSubmatch")

	currency := trimRight(values[1])
	mineral := trimRight(values[2])

	creditValue, err := getMineralValue(currency, mineral)
	if err != nil {
		log.Error().Err(err).Msg("getMineralValue failed")
		return
	}

	currencyMineral := combineString(currency, mineral)
	answer := fmt.Sprintf("%s is %d Credits", currencyMineral, int(creditValue))
	fmt.Println("answer", answer)
}

func combineString(currency, mineral string) string {
	return fmt.Sprintf("%s %s", currency, mineral)
}

func AnswerCreditComparisonQuestion(text string) {
	log.Debug().Msg("AnswerCreditComparisonQuestion")

	pattern := regexp.MustCompile(regexCreditComparisonQuestion)
	values := pattern.FindStringSubmatch(text)
	log.Debug().Strs("values", values).Msg("FindStringSubmatch")

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
	fmt.Println("answer", answer)

}

func AnswerCurrencyComparisonQuestion(text string) {
	log.Debug().Msg("AnswerCurrencyComparisonQuestion")

	pattern := regexp.MustCompile(regexCurrencyComparisonQuestion)
	values := pattern.FindStringSubmatch(text)
	log.Debug().Strs("values", values).Msg("FindStringSubmatch")

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

	fmt.Println("leftCurrency =", leftCurrency)
	fmt.Println("leftCurrencyValue =", leftCurrencyValue)

	fmt.Println("rightCurrency =", rightCurrency)
	fmt.Println("rightCurrencyValue =", rightCurrencyValue)

	comparator := "smaller"
	if leftCurrencyValue > rightCurrencyValue {
		comparator = "larger"
	}

	answer := fmt.Sprintf("%s is %s than %s", leftCurrency, comparator, rightCurrency)
	fmt.Println("answer", answer)

}

func getCurrencyValue(text string) (int64, error) {

	romanStr, err := ConvertNewCurrencyToRoman(text)
	if err != nil {
		return 0, err
	}
	romanNum, err := numerus.Parse(romanStr)
	if err != nil {
		return 0, err
	}

	return int64(romanNum.Value()), nil
}

func getMineralValue(currency, mineral string) (value float64, err error) {
	strRoman, err := ConvertNewCurrencyToRoman(currency)
	if err != nil {
		log.Error().Err(err).Msg("ConvertNewCurrencyToRoman failed")
	}

	romanNum, err := numerus.Parse(strRoman)
	if err != nil {
		log.Error().Err(err).Msg("numerus.Parse failed")
		fmt.Println("Requested number is in invalid format")
		return
	}
	// fmt.Println("romanNum", romanNum.Value())

	creditValue := float64(newMineralsMap[mineral]) * float64(romanNum.Value())
	// fmt.Println("creditValue", creditValue)

	return creditValue, nil
}

func trimRight(text string) string {
	return strings.TrimRight(text, " ")
}
