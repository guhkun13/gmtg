package service

import (
	"github.com/dhowden/numerus"
)

type CurrencyIface interface {
	IsMatchAssignValue(text string) bool
	AssignValue(text string) error
	GetValue(text string) (numerus.Numeral, error)
}
