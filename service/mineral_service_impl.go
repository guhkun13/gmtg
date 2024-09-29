package service

import (
	"regexp"
	"strconv"

	"github.com/guhkun13/gmtg/libs"
	"github.com/guhkun13/gmtg/utils"
	"github.com/rs/zerolog/log"
)

var mineralsMap = make(map[string]float64)

type Mineral struct {
	Name   string
	Credit float64
}

type MineralImpl struct {
	Regexp          string
	CurrencyService CurrencyIface
}

func NewMineralImpl(regexp string, currencyService CurrencyIface) MineralIface {
	return &MineralImpl{
		Regexp:          regexp,
		CurrencyService: currencyService,
	}
}

func (s *MineralImpl) IsMatchAssignValue(text string) bool {
	return regexp.MustCompile(s.Regexp).Match([]byte(text))
}

func (s *MineralImpl) AssignValue(text string) error {
	values := regexp.MustCompile(s.Regexp).FindStringSubmatch(text)
	// log.Debug().Interface("values", values).Msg("FindStringSubmatch")

	currency := utils.TrimRight(values[1])
	mineralName := utils.TrimRight(values[2])
	strTotalAmount := utils.TrimRight(values[3])

	romanNum, err := s.CurrencyService.GetValue(currency)
	if err != nil {
		log.Error().Err(err).Str("input", currency).Msg("CurrencyService.GetValue failed")
		return err
	}

	intTotalAmount, err := strconv.Atoi(strTotalAmount)
	if err != nil {
		log.Error().Err(err).Str("input", strTotalAmount).Msg("Credit value must be number")
		return err
	}

	creditValue := float64(float64(intTotalAmount) / float64(romanNum.Value()))

	mineralsMap[mineralName] = creditValue
	return nil
}

func (s *MineralImpl) GetValue(currencies, mineral string) (res Mineral, err error) {

	val, ok := s.isExist(mineral)
	if !ok {
		err = libs.ErrMineralDoesNotExist
		return
	}

	roman, err := s.CurrencyService.GetValue(currencies)
	if err != nil {
		return
	}

	creditValue := val * float64(roman.Value())

	return Mineral{
		Name:   mineral,
		Credit: creditValue,
	}, nil

}

func (s *MineralImpl) isExist(text string) (float64, bool) {
	val, exist := mineralsMap[text]
	return val, exist
}
