package service

import (
	"regexp"
	"strings"

	"github.com/dhowden/numerus"
	"github.com/guhkun13/gmtg/libs"
	"github.com/guhkun13/gmtg/utils"
)

var currenciesMap = make(map[string]string)

type RomanNumeral struct {
	String string
	Value  int64
}

type CurrencyImpl struct {
	Regexp string
}

func NewCurrencyImpl(regexp string) CurrencyIface {
	return &CurrencyImpl{
		Regexp: regexp,
	}
}

func (s *CurrencyImpl) IsMatchAssignValue(text string) bool {
	return regexp.MustCompile(s.Regexp).MatchString(text)
}

func (s *CurrencyImpl) AssignValue(text string) error {
	values := regexp.MustCompile(s.Regexp).FindStringSubmatch(text)
	if len(values) < 2 {
		return libs.ErrUnrecognizedText
	}
	currencyStr := utils.TrimRight(values[1])
	romanStr := utils.TrimRight(values[2])

	romanNumber, err := numerus.Parse(romanStr)
	if err != nil {
		return err
	}

	currenciesMap[currencyStr] = romanNumber.String()

	return nil
}

func (s *CurrencyImpl) GetValue(text string) (res RomanNumeral, err error) {

	if !strings.Contains(text, " ") {
		roman, errVal := s.validateAndGetRoman(text)
		if errVal != nil {
			err = errVal
			return
		}
		return roman, nil
	}
	currencies := strings.Split(text, " ")

	romanString := ""
	for _, currency := range currencies {
		roman, errVal := s.validateAndGetRoman(currency)
		if errVal != nil {
			err = errVal
			return
		}
		romanString += roman.String
	}

	res, err = toRoman(romanString)
	if err != nil {
		return
	}

	return res, nil
}

func (s *CurrencyImpl) validateAndGetRoman(text string) (res RomanNumeral, err error) {
	val, ok := s.isExist(text)
	if !ok {
		err = libs.ErrCurrencyDoesNotExist
		return
	}

	roman, errConv := toRoman(val)
	if errConv != nil {
		err = errConv
		return
	}

	return roman, nil
}

func toRoman(currency string) (RomanNumeral, error) {
	roman, err := numerus.Parse(currency)
	if err != nil {
		err = libs.ErrNumberInvalidFormat
		return RomanNumeral{}, err
	}

	return RomanNumeral{
		String: roman.String(),
		Value:  int64(roman.Value()),
	}, nil
}

func (s *CurrencyImpl) isExist(text string) (string, bool) {
	val, ok := currenciesMap[text]
	return val, ok
}
