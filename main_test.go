package main

import (
	"testing"

	"github.com/guhkun13/gmtg/config"
	"github.com/guhkun13/gmtg/service"
)

func TestMain(t *testing.T) {

	regexps := config.InitRegexps()
	services := service.NewServices(regexps)

	testCasesAssignCurrencies := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "#1 case 1",
			input: "glob is I",
			want:  "",
		},
		{
			desc:  "#1 case 2",
			input: "prok is V",
			want:  "",
		},
		{
			desc:  "#1 case 3",
			input: "pish is X",
			want:  "",
		},
		{
			desc:  "#1 case 4",
			input: "tegj is L",
			want:  "",
		},
	}

	testCasesAssignMinerals := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "#2 case 1",
			input: "glob glob Silver is 34 Credits",
			want:  "",
		},
		{
			desc:  "#2 case 2",
			input: "glob prok Gold is 57800 Credits",
			want:  "",
		},
		{
			desc:  "#2 case 3",
			input: "pish pish Iron is 3910 Credits",
			want:  "",
		},
	}

	testCasesHowMuchQuestion := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "#3 case 1",
			input: "how much is pish tegj glob glob ?",
			want:  "pish tegj glob glob is 42",
		},
	}

	testCasesHowManyCreditQuestion := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "#4 case 1",
			input: "how many Credits is glob prok Silver ?",
			want:  "glob prok Silver is 68 Credits",
		},
		{
			desc:  "#4 case 2",
			input: "how many Credits is glob glob Gold ?",
			want:  "glob glob Gold is 28900 Credits",
		},
		{
			desc:  "#4 case 3",
			input: "how many Credits is glob glob glob glob glob glob Gold ?",
			want:  "Requested number is in invalid format",
		},
		{
			desc:  "#4 case 4",
			input: "how many Credits is pish tegj glob Iron ?",
			want:  "pish tegj glob Iron is 8015.5 Credits",
		},
		{
			desc:  "#4 case 5",
			input: "how many Credits is glob Silver ?",
			want:  "glob Silver is 17 Credits",
		},
		{
			desc:  "#4 case 6",
			input: "how many Credits is Silver ?",
			want:  "I have no idea what you are talking about",
		},
	}

	testCasesCreditComparisonQuestion := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "#5 case 1",
			input: "Is glob prok larger than pish pish?",
			want:  "glob prok is smaller than pish pish",
		},
		{
			desc:  "#5 case 2",
			input: "Is tegj glob glob smaller than glob prok?",
			want:  "tegj glob glob is larger than glob prok",
		},
	}

	testCasesCurrencyComparisonQuestion := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "#6 case 1",
			input: "Is glob prok larger than pish pish?",
			want:  "glob prok is smaller than pish pish",
		},
		{
			desc:  "#6 case 2",
			input: "Is tegj glob glob smaller than glob prok?",
			want:  "tegj glob glob is larger than glob prok",
		},
	}

	testCasesErrorUnknownText := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "#7 case 1",
			input: "how much wood could a woodchuck chuck if a woodchuck could chuck wood ?",
			want:  "I have no idea what you are talking about",
		},
	}

	testCases := append(testCasesAssignCurrencies, testCasesAssignMinerals...)
	testCases = append(testCases, testCasesHowMuchQuestion...)
	testCases = append(testCases, testCasesHowManyCreditQuestion...)
	testCases = append(testCases, testCasesCreditComparisonQuestion...)
	testCases = append(testCases, testCasesCurrencyComparisonQuestion...)
	testCases = append(testCases, testCasesErrorUnknownText...)

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			got := services.EvaluateText(tt.input)
			if got != tt.want {
				t.Errorf("EvaluateText [%v], got (%v), want (%v)", tt.input, got, tt.want)
			}
		})
	}
}
