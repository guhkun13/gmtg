package service

import (
	"regexp"

	"github.com/dhowden/numerus"
	"github.com/guhkun13/gmtg/utils"
)

var mineralsMap = make(map[string]float64)

type Mineral struct {
	Name  string
	Value float64
}

type MineralIface interface {
	AssignValue(text string) error
	GetValue(text string) (int64, error)
}

type MineralImpl struct {
	Regexp          string
	CurrencyService service.Currency
}

func NewMineralImpl(regexp string, currencyService service.Currency) MineralIface {
	return &MineralImpl{
		Regexp:          regexp,
		CurrencyService: currencyService,
	}
}

func (s *MineralImpl) AssignValue(text string) error {
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

func (s *MineralImpl) GetValue(text string) (int64, error) {

	creditValue := float64(mineralsMap[mineral]) * float64(romanNum.Value())
	// fmt.Println("creditValue", creditValue)

	return creditValue, nil
}

func (s *MineralImpl) isExist(text string) (float64, bool) {
	val, exist := mineralsMap[text]

	return val, exist
}
