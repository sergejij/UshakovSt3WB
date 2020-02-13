package facade

import (
	"fasad/pkg/balance"
	"fasad/pkg/verifier"
	"fasad/pkg/creditHistory"
	"fmt"
)

type CreditTaker interface {
	CreditTake() string
}

type credit struct {
	name			string
	accountMoney	int
	desiredCredit 	int
	maxCredit 		int
}

func (c *credit) CreditTake() string {
	var verdict string

	currentMoney := balance.NewAmount(c.accountMoney)
	moneyInAccount := currentMoney.AmountCheck()

	controller := verifier.NewVerifier(c.maxCredit)
	isVerified := controller.Verify(c.desiredCredit)

	history := creditHistory.NewHistoryChecker(c.name)
	isGoodCreditStory := history.HistoryCheck(c.name) // true/false - credit story

	return fmt.Sprintf("%s has been %s for the %d$ credit.\n" +
		"With a %s credit standing. And he has a %d$ in balance.", c.name,
		verdict, c.desiredCredit, history, moneyInAccount)
}

// NewCredit initializes the Credit.
func NewCreditTaker(personName string, money int, wantMoney int, maxSize int) CreditTaker {
	return &credit{
		name: 			personName,
		accountMoney: 	money,
		desiredCredit: 	wantMoney,
		maxCredit: 		maxSize,
	}
}
