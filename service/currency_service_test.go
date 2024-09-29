package service

import (
	"testing"

	"github.com/guhkun13/gmtg/config"
	"github.com/guhkun13/gmtg/libs"
)

func TestIsMatchAssignValue(t *testing.T) {
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

func TestAssignValue(t *testing.T) {
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
