package service

import (
	"testing"

	"github.com/guhkun13/gmtg/config"
)

const (
	howMuchQuestion            string = "how much question"
	howManyCreditQuestion      string = "how many credit question"
	creditComparisonQuestion   string = "credit comparison question"
	currencyComparisonQuestion string = "currency comparison question"
)

func TestQuestion(t *testing.T) {
	regexps := config.InitRegexps()
	srv := NewQuestionImpl(regexps)

	testCases := []struct {
		name         string
		question     string
		questionType string
		want         bool
		err          error
	}{
		{
			name:         "#1 case 1",
			question:     "how much is one two six?",
			questionType: howMuchQuestion,
			want:         true,
			err:          nil,
		},
		{
			name:         "#1 case 2. no question mark",
			question:     "how much is one two five", //no question mark, so it will false/unrecognized
			questionType: howMuchQuestion,
			want:         false,
			err:          nil,
		},
		{
			name:         "#1 case 3",
			question:     "how many Credits is satu Silver ?",
			questionType: howManyCreditQuestion,
			want:         true,
			err:          nil,
		},
		{
			name:         "#1 case 4. no mineral name",
			question:     "how many Credits is five six ?", //no mineral name, so it will false
			questionType: howManyCreditQuestion,
			want:         false,
			err:          nil,
		},
		{
			name:         "#1 case 5",
			question:     "Does two Iron has more Credits than one Gold ?",
			questionType: creditComparisonQuestion,
			want:         true,
			err:          nil,
		},
		{
			name:         "#1 case 6",
			question:     "Does Iron has more Credits than one Gold ", // no currency, so it will false
			questionType: creditComparisonQuestion,
			want:         false,
			err:          nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.questionType == howMuchQuestion {
				got := srv.IsMatchHowMuchQuestion(tt.question)
				if got != tt.want {
					t.Errorf("IsMatchHowMuchQuestion [%v], got %v, want %v", tt.question, got, tt.want)
				}
			} else if tt.questionType == howManyCreditQuestion {
				got := srv.IsMatchHowManyCreditQuestion(tt.question)
				if got != tt.want {
					t.Errorf("IsMatchHowManyCreditQuestion [%v], got %v, want %v", tt.question, got, tt.want)
				}
			} else if tt.questionType == creditComparisonQuestion {
				got := srv.IsMatchCreditComparisonQuestion(tt.question)
				if got != tt.want {
					t.Errorf("IsMatchHowManyCreditQuestion [%v], got %v, want %v", tt.question, got, tt.want)
				}
			} else {
				t.Errorf("unhandled questions type [%v] [%v]", tt.question, tt.questionType)
			}
		})
	}

}
