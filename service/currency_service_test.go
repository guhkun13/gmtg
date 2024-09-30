package service

import (
	"testing"

	"github.com/guhkun13/gmtg/config"
	"github.com/guhkun13/gmtg/libs"
)

func TestCurrencyIsMatchAssignValue(t *testing.T) {
	regexp := config.RegexAssignCurrency
	srv := NewCurrencyImpl(regexp)

	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "positive case 1",
			input: "glob is I",
			want:  true,
		},
		{
			name:  "positive case 2",
			input: "prok is V",
			want:  true,
		},
		{
			name:  "positive case 3",
			input: "pish is X",
			want:  true,
		},
		{
			name:  "positive case 4",
			input: "tegj is L",
			want:  true,
		},
		{
			name:  "negative case 1",
			input: "",
			want:  false,
		},
		{
			name:  "negative case 2",
			input: "dua is II",
			want:  false,
		},
		{
			name:  "negative case 3",
			input: "what is this",
			want:  false,
		},
		{
			name:  "negative case 4",
			input: "limalima is VV",
			want:  false,
		},
		{
			name:  "negative case 5",
			input: "satu Is I",
			want:  false,
		},
		{
			name:  "negative case 6",
			input: "Satu is I",
			want:  false,
		},
		{
			name:  "negative case 7",
			input: "SATU IS I",
			want:  false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := srv.IsMatchAssignValue(tt.input)
			if got != tt.want {
				t.Errorf("IsMatchAssignValue [%v], got %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCurrencyAssignValue(t *testing.T) {
	regexp := config.RegexAssignCurrency
	srv := NewCurrencyImpl(regexp)

	testCases := []struct {
		name  string
		input string
		want  error
	}{
		{
			name:  "positive case 1",
			input: "glob is I",
			want:  nil,
		},
		{
			name:  "positive case 2",
			input: "prok is V",
			want:  nil,
		},
		{
			name:  "positive case 3",
			input: "pish is X",
			want:  nil,
		},
		{
			name:  "positive case 4",
			input: "tegj is L",
			want:  nil,
		},
		{
			name:  "negative case 1",
			input: "",
			want:  libs.ErrUnrecognizedText,
		},
		{
			name:  "negative case 2",
			input: "Satu is I",
			want:  libs.ErrUnrecognizedText,
		},
		{
			name:  "negative case 3",
			input: "Satu is I",
			want:  libs.ErrUnrecognizedText,
		},
		{
			name:  "negative case 4",
			input: "tiga is III",
			want:  libs.ErrUnrecognizedText,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := srv.AssignValue(tt.input)
			if got != tt.want {
				t.Errorf("AssignValue [%v], got %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCurrencyGetValue(t *testing.T) {
	regexp := config.RegexAssignCurrency
	srv := NewCurrencyImpl(regexp)

	testCasesInput := []struct {
		name  string
		input string
		want  error
	}{
		{
			name:  "#1 case 1",
			input: "satu is I",
			want:  nil,
		},
		{
			name:  "#1 case 2",
			input: "lima is V",
			want:  nil,
		},
		{
			name:  "#1 case 3",
			input: "sepuluh is X",
			want:  nil,
		},
	}

	for _, tt := range testCasesInput {
		t.Run(tt.name, func(t *testing.T) {
			got := srv.AssignValue(tt.input)
			if got != tt.want {
				t.Errorf("AssignValue [%v], got %v, want %v", tt.input, got, tt.want)
			}
		})
	}

	testCasesGetValue := []struct {
		name    string
		input   string
		wantStr string
		wantInt int64
		err     error
	}{
		{
			name:    "#2 case 1",
			input:   "satu",
			wantStr: "I",
			wantInt: 1,
			err:     nil,
		},
		{
			name:    "#2 case 2",
			input:   "satu satu",
			wantStr: "II",
			wantInt: 2,
			err:     nil,
		},
		{
			name:    "#2 case 3",
			input:   "lima",
			wantStr: "V",
			wantInt: 5,
			err:     nil,
		},
		{
			name:    "#2 case 4",
			input:   "sebelas",
			wantStr: "",
			wantInt: 0,
			err:     libs.ErrCurrencyDoesNotExist,
		},
		{
			name:    "#2 case 5",
			input:   "sepuluh",
			wantStr: "X",
			wantInt: 10,
			err:     nil,
		},
	}
	for _, ttg := range testCasesGetValue {
		t.Run(ttg.name, func(t *testing.T) {
			got, err := srv.GetValue(ttg.input)
			if got.String != ttg.wantStr || got.Value != ttg.wantInt || ttg.err != err {
				t.Errorf("GetValue [%v], got %v, want %v, %v, %v", ttg.input, got, ttg.wantStr, ttg.wantInt, ttg.err)
			}
		})
	}
}
