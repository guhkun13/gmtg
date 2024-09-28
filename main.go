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
	regexNewCurrency           string = `^([a-z]+) is ([IVXLCDM])$`
	regexNewMineral            string = `^(([a-z]+ )+)([A-Z][a-z]+) is (\d+) Credits$`
	regexHowMuchQuestion       string = `^(how much is) (([a-z]+\s{0,1})+)\?$`
	regexHowManyCreditQuestion string = `^how many Credits is (([a-z]+ )+)([A-Z][a-z]+)\s{0,1}\?$`
)

var newCurrenciesMap = make(map[string]string)
var newMineralsMap = make(map[string]float64)

func main() {
	fmt.Println("Halo dunia")

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

// assigning new currency to roman value
func AssignNewCurrency(text string) {
	// log.Debug().Msg("AssignNewCurrency")

	pattern := regexp.MustCompile(regexNewCurrency)
	values := pattern.FindStringSubmatch(text)
	// log.Debug().Strs("values", values).Msg("FindStringSubmatch")

	currency, roman := values[1], values[2]

	newCurrenciesMap[currency] = roman

	// log.Debug().
	// 	Str("1-currency", currency).
	// 	Str("2-roman", roman).
	// 	Interface("3-newCurrenciesMap", newCurrenciesMap).
	// 	Msg("Value")

}

func ConvertNewCurrencyToRoman(text string) (roman string, err error) {
	currencies := strings.Split(text, " ")

	if len(currencies) == 1 {
		return CurrencyToRoman(text)
	}

	// fmt.Println("more than 1 currency, iterate!")

	finalRoman := ""
	for _, c := range currencies {
		roman, err = CurrencyToRoman(c)

		finalRoman += roman
	}

	// fmt.Println("finalRoman", finalRoman)
	return finalRoman, err
}

func CurrencyToRoman(currency string) (string, error) {
	val, err := numerus.Parse(newCurrenciesMap[currency])
	if err != nil {
		return "", err
	}

	return val.String(), nil
}

// assigning value to Mineral
func AssignNewMineral(text string) {
	// log.Debug().Msg("AssignNewMineral")

	pattern := regexp.MustCompile(regexNewMineral)
	values := pattern.FindStringSubmatch(text)
	// log.Debug().Strs("values", values).Msg("FindStringSubmatch")

	currency, mineral, strTotalAmount := values[1], values[3], values[4]

	// fmt.Println("currency", currency)
	// fmt.Println("mineral", mineral)
	// fmt.Println("strTotalAmount", strTotalAmount)

	romanStr, err := ConvertNewCurrencyToRoman(currency)
	if err != nil {
		log.Error().Err(err).Msg("ConvertNewCurrencyToRoman failed")
		return
	}
	// fmt.Println("romanValue", romanValue)

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

	creditValue := float64(float64(intTotalAmount) / float64(romanNum.Value()))

	// fmt.Println("creditValue", creditValue)
	newMineralsMap[mineral] = creditValue

	log.Debug().
		Interface("value", newMineralsMap).
		Msg("newMineralsMap")

}

func AnswerHowMuchQuestion(text string) {
	// log.Debug().Msg("AnswerHowMuchQuestion")

	pattern := regexp.MustCompile(regexHowMuchQuestion)
	values := pattern.FindStringSubmatch(text)
	// log.Debug().Strs("values", values).Msg("FindStringSubmatch")

	currency := values[2]
	currency = strings.TrimRight(currency, " ")
	// fmt.Println("currency", currency)

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
	log.Debug().Msg("AnswerHowManyCreditQuestion")

	pattern := regexp.MustCompile(regexHowManyCreditQuestion)
	values := pattern.FindStringSubmatch(text)
	log.Debug().Strs("values", values).Msg("FindStringSubmatch")

	currency := values[1]
	mineral := values[3]

	strRoman, err := ConvertNewCurrencyToRoman(currency)
	if err != nil {
		log.Error().Err(err).Msg("ConvertNewCurrencyToRoman failed")
	}

	romanNum, err := numerus.Parse(strRoman)
	if err != nil {
		log.Error().Err(err).Msg("numerus.Parse failed")
	}
	fmt.Println("romanNum", romanNum.Value())

	creditValue := float64(newMineralsMap[mineral]) * float64(romanNum.Value())
	fmt.Println("creditValue", creditValue)

	currencyMineral := currency + mineral
	answer := fmt.Sprintf("%s is %d Credits", currencyMineral, int(creditValue))
	fmt.Println("answer", answer)

}

// strVal := "IV"
// num, err := numerus.Parse(strVal)
// if err != nil {
// 	log.Error().Err(err).Msg("failed to parse")
// }

// log.Info().
// 	Str("From", strVal).
// 	Str("String", num.String()).
// 	Int("Int", int(num.Value())).Msg("Parse result")

// func asyncEvaluateText(wg *sync.WaitGroup, text string) {
// 	if IsMatchNewCurrency(text) {
// 		AssignNewCurrency(text)
// 	}

// 	wg.Done()
// }
