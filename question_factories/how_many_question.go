package questionfactories

import (
	"fmt"
	"regexp"

	"github.com/guhkun13/gmtg/utils"
)

type HowManyQuestionImpl struct {
	Regex string
}

func NewHowManyQuestionImpl(regex string) QuestionFactory {
	return &HowManyQuestionImpl{
		Regex: regex,
	}
}

func (f *HowManyQuestionImpl) Answer(text string) {
	values := regexp.MustCompile(f.Regex).FindStringSubmatch(text)
	// log.Debug().Strs("values", values).Msg("FindStringSubmatch")

	currency := utils.TrimRight(values[1])
	mineral := utils.TrimRight(values[2])

	creditValue, err := utils.GetMineralValue(currency, mineral)
	if err != nil {
		// log.Error().Err(err).Msg("getMineralValue failed")
		return
	}

	currencyMineral := combineString(currency, mineral)
	answer := fmt.Sprintf("%s is %d Credits", currencyMineral, int(creditValue))

	fmt.Println(answer)
	WriteToOutput(answer)

}
