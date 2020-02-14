package facade

type creditReputation interface {
	CheckHistory(personName string) bool
}

type verifier interface {
	Verify(loanSize int) bool
}

type Credit interface {
	Take(controller verifier, history creditReputation) bool
}

type credit struct {
	name			string
	accountMoney	int
	desiredCredit 	int
}

func (c *credit) Take(controller verifier, repute creditReputation) bool {
	var result bool
	isSumLessMax := controller.Verify(c.desiredCredit)
	isGoodCreditHistory := repute.CheckHistory(c.name)
	isBalanceGood := c.accountMoney >= c.desiredCredit
	if (isSumLessMax && (isGoodCreditHistory || isBalanceGood)) ||
		(isGoodCreditHistory && (isSumLessMax || isBalanceGood)) {
		result = true
	}
	return result
}

// NewCredit initializes the Credit.
func NewCredit(name string, accountMoney int, desiredCredit int) Credit {
	return &credit{
		name: 			name,
		accountMoney: 	accountMoney,
		desiredCredit: 	desiredCredit,
	}
}
