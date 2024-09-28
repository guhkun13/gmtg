package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/guhkun13/gmtg/config"
	log "github.com/rs/zerolog/log"
)

const (
	regexNewCurrency string = "^([a-z]+) is ([IVXLCDM])$"
)

var newCurrenciesMap = make(map[string]string)

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
	for scanner.Scan() {
		txt := scanner.Text()
		log.Info().Str("val", txt).Msgf("[%d]", idx)
		evaluateText(txt)
		idx++

	}

	err = scanner.Err()
	if err != nil {
		log.Fatal().Err(err).Msg("scannerfailed  ")
	}

}

func evaluateText(text string) {

	if IsMatchNewCurrency(text) {
		AssignNewCurrency(text)
	}
}

func IsMatchNewCurrency(text string) bool {
	pattern := regexp.MustCompile(regexNewCurrency)
	return pattern.Match([]byte(text))
}

// assiging new currency to roman value
func AssignNewCurrency(text string) {
	log.Debug().Msg("AssignNewCurrency")

	pattern := regexp.MustCompile(regexNewCurrency)
	values := pattern.FindStringSubmatch(text)
	log.Debug().Strs("values", values).Msg("FindStringSubmatch")

	currency, roman := values[1], values[2]

	newCurrenciesMap[currency] = roman

	log.Debug().
		Str("1-currency", currency).
		Str("2-roman", roman).
		Interface("3-newCurrenciesMap", newCurrenciesMap).
		Msg("Value")

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
