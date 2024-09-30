package service

import (
	"testing"

	"github.com/guhkun13/gmtg/config"
)

func TestMineras(t *testing.T) {
	currencyService := NewCurrencyImpl(config.RegexAssignCurrency)
	srv := NewMineralImpl(config.RegexAssignMineral, currencyService)

	// first, set the currency value
	testCasesAssignCurrency := []struct {
		name  string
		input string
		err   error
	}{
		{
			name:  "#1 case 1",
			input: "one is I",
			err:   nil,
		},
		{
			name:  "#1 case 2",
			input: "five is V",
			err:   nil,
		},
	}

	for _, tt := range testCasesAssignCurrency {
		t.Run(tt.name, func(t *testing.T) {
			got := currencyService.AssignValue(tt.input)
			if got != tt.err {
				t.Errorf("currencyService.AssignValue [%v], got %v, want %v", tt.input, got, tt.err)
			}
		})
	}

	// second, set the Mineral value
	testCasesDefineMineral := []struct {
		name  string
		input string
		err   error
	}{
		{
			name:  "#2 case 1",
			input: "one Silver is 100 Credits",
			err:   nil,
		},
		{
			name:  "#1 case 2",
			input: "one Copper is 10 Credits",
			err:   nil,
		},
	}

	for _, tt := range testCasesDefineMineral {
		t.Run(tt.name, func(t *testing.T) {
			got := srv.AssignValue(tt.input)
			if got != tt.err {
				t.Errorf("srv.AssignValue [%v], got %v, want %v", tt.input, got, tt.err)
			}
		})
	}

	// third, get the Mineral value

	testCasesGetMineral := []struct {
		name        string
		currency    string
		mineralName string
		wantStr     string
		wantInt     int64
		err         error
	}{
		{
			name:        "#3 case 1",
			currency:    "one",
			mineralName: "Silver",
			wantStr:     "Silver",
			wantInt:     100,
			err:         nil,
		},
		{
			name:        "#1 case 2",
			currency:    "one",
			mineralName: "Copper",
			wantStr:     "Copper",
			wantInt:     10,
			err:         nil,
		},
		{
			name:        "#1 case 3",
			currency:    "five",
			mineralName: "Copper",
			wantStr:     "Copper",
			wantInt:     50,
			err:         nil,
		},
		{
			name:        "#1 case 3",
			currency:    "one five",
			mineralName: "Copper",
			wantStr:     "Copper",
			wantInt:     40,
			err:         nil,
		},
	}

	for _, tt := range testCasesGetMineral {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.GetValue(tt.currency, tt.mineralName)
			if got.Name != tt.wantStr || got.Credit != float64(tt.wantInt) || err != tt.err {
				t.Errorf("srv.GetValue [%v %v], got %v, want (%v %v)", tt.currency, tt.mineralName, got, tt.wantStr, tt.wantInt)
			}
		})
	}

}
