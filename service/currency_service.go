package service

type CurrencyIface interface {
	IsMatchAssignValue(text string) bool
	AssignValue(text string) error
	GetValue(text string) (RomanNumeral, error)
}
