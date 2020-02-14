package facade

import (
	"fasad/pkg/creditHistory"
	"fasad/pkg/verifier"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	persons = []string{
		"John",
		"Brad",
		"Alex",
		"Fill",
		"Nik",
		"Jack",
	}
	balance = 			2100
	maxCredit = 		1000
	desiredCredit =		1500

)

func TestFacade(t *testing.T) {
	var expect bool
	for _, person := range persons {
		story := creditHistory.NewCreditHistory(person)
		isGoodCreditStory := story.HistoryCheck(person)
		controller := verifier.NewVerifier(maxCredit)
		isSumLessMax := controller.Verify(desiredCredit)
		facade := NewCredit(person, balance,desiredCredit)
		isBalanceGood := balance >= desiredCredit
		if (isSumLessMax && (isGoodCreditStory || isBalanceGood)) ||
			(isGoodCreditStory && (isSumLessMax || isBalanceGood)) {
			expect = true
		} else {
			expect = false
		}
		result := facade.TakeCredit()
		assert.Equal(t, expect, result, person)
	}
}
