package facade

import (
	"fasad/pkg/balance"
	"fasad/pkg/bank"
	"fasad/pkg/creditHistory"
	"fmt"
)

type CreditTaker interface {
	CreditTake() string
}
type credit struct {
	name string
	accountMoney int
	loanSize int
}
func (c *credit) CreditTake() string {
	balance, err := balance.NewBalance(c.accountMoney)
	if err != nil {
		fmt.Errorf("Error: the checking balance is not working.\n")
	}
	moneyInAccount := balance.BalanceCheck()
	bank, err := bank.NewBank()
	if err != nil {
		fmt.Errorf("Error: the Bank is not working.\n")
	}
	bankResult := bank.Verify(c.loanSize)
	if bankResult == "approved" {
		moneyInAccount += c.loanSize
	}
	creditHistory, err := creditHistory.NewHistoryChecker()
	if err != nil {
		fmt.Errorf("Error: the credit history is not working.\n")
	}
	return fmt.Sprintf("%s has been %s for the %d$ credit\n" +
		"With a %s credit standing. And he has a %d$ in balance.", c.name, bankResult, c.loanSize, creditHistory.HistoryCheck(c.name), moneyInAccount)
}

// NewCredit initializes the Credit.
func NewCreditTaker(personName string, money int, wantMoney int) (CreditTaker, error) {
	return &credit{
		name: personName,
		accountMoney: money,
		loanSize: wantMoney,
	}, nil
}
