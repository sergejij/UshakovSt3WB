package facade

type historyReputation interface {
	CheckHistory(name string) bool
}

type verifier interface {
	Verify(loanSize int) bool
}

// Credit is a service that takes credit
type Credit interface {
	Take(controller verifier, historyReputation historyReputation) bool
}

type credit struct {
	name          string
	accountMoney  int
	desiredCredit int
}

// Take ...
func (c *credit) Take(money, name string) (result bool) {
	var (
		isSumLessMax, isGoodHistory, isBalanceGood bool
	)
	isSumLessMax = controller.Verify(c.desiredCredit)
	isGoodHistory = historyReputation.CheckHistory(c.name)
	isBalanceGood = c.accountMoney >= c.desiredCredit
	return (isSumLessMax && isGoodHistory) ||
		(isBalanceGood && isGoodHistory) ||
		(isSumLessMax && isBalanceGood)
}

// NewCredit initializes the Credit.
func NewCredit(name string, accountMoney int, desiredCredit int) Credit {
	return &credit{
		name:          name,
		accountMoney:  accountMoney,
		desiredCredit: desiredCredit,
	}
}
