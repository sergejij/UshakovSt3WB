package facade

import (
	"fasad/pkg/creditHistory"
	"fasad/pkg/models"
	"fasad/pkg/verifier"
)

type Credit interface {
	TakeCredit()	bool
}

type credit struct {
	name			string
	accountMoney	int
	desiredCredit 	int
}

func (c *credit) TakeCredit() bool {
	var result bool
	controller := verifier.NewVerifier(models.MaxCredit)
	isSumLessMax := controller.Verify(c.desiredCredit)
	story := creditHistory.NewCreditHistory(c.name)
	isGoodCreditStory := story.HistoryCheck(c.name)
	isBalanceGood := c.accountMoney >= c.desiredCredit
	if (isSumLessMax && (isGoodCreditStory || isBalanceGood)) ||
		(isGoodCreditStory && (isSumLessMax || isBalanceGood)) {
		result = true
	} else {
		result = false
	}
	return result
}

// NewCredit initializes the Credit.
func NewCredit(personName string, money int, wantMoney int) Credit {
	return &credit{
		name: 			personName,
		accountMoney: 	money,
		desiredCredit: 	wantMoney,
	}
}
