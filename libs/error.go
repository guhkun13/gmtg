package libs

import "errors"

var (
	ErrQuestionUnrecognized = errors.New("I have no idea what you are talking about ")
	ErrNumberInvalidFormat  = errors.New("Requested number is in invalid format")
	ErrCurrencyDoesNotExist = errors.New("Currency does not exist")
)
