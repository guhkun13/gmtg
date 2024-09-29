package service

import (
	"regexp"
	"strings"

	"github.com/dhowden/numerus"
	"github.com/guhkun13/gmtg/libs"
	"github.com/guhkun13/gmtg/utils"
)

var currenciesMap = make(map[string]string)

type CurrencyIface interface {
	AssignValue(text string) error
	GetValue(text string) (int64, error)
	CurrenciesToRoman(text string) (res numerus.Numeral, err error)
}

type CurrencyImpl struct {
	Regexp string
}

func NewCurrencyImpl(regexp string) CurrencyIface {
	return &CurrencyImpl{
		Regexp: regexp,
	}
}

func (s *CurrencyImpl) AssignValue(text string) error {
	values := regexp.MustCompile(s.Regexp).FindStringSubmatch(text)
	currencyStr := utils.TrimRight(values[1])
	romanStr := utils.TrimRight(values[2])

	romanNumber, err := numerus.Parse(romanStr)
	if err != nil {
		return err
	}

	currenciesMap[currencyStr] = romanNumber.String()
	return nil
}

func (s *CurrencyImpl) GetValue(text string) (int64, error) {
	val, ok := s.isExist(text)
	if !ok {
		err := libs.ErrCurrencyDoesNotExist
		return 0, err
	}

	romanNumber, err := toRoman(val)
	if err != nil {
		return 0, err
	}

	return int64(romanNumber.Value()), nil
}

func (s *CurrencyImpl) CurrenciesToRoman(text string) (res numerus.Numeral, err error) {
	currencies := strings.Split(text, " ")
	if len(currencies) == 1 {
		return toRoman(text)
	}

	romanString := ""
	for _, currency := range currencies {
		val, ok := s.isExist(currency)
		if !ok {
			err = libs.ErrCurrencyDoesNotExist
			return
		}
		roman, errConv := toRoman(val)
		if errConv != nil {
			err = errConv
			return
		}
		romanString += roman.String()
	}

	res, err = toRoman(romanString)
	if err != nil {
		return
	}

	return res, nil
}

func toRoman(currency string) (roman numerus.Numeral, err error) {
	roman, err = numerus.Parse(currency)
	if err != nil {
		err = libs.ErrNumberInvalidFormat
		return
	}

	return roman, nil
}

func (s *CurrencyImpl) isExist(text string) (string, bool) {
	val, ok := currenciesMap[text]
	return val, ok
}
