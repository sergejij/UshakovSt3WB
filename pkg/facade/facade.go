package facade

type reputation interface {
	CheckHistory(personName string) bool
}

type verifier interface {
	Verify(loanSize int) bool
}

type Credit interface {
	Take(controller verifier, history reputation) bool
}

type credit struct {
	name			string
	accountMoney	int
	desiredCredit 	int
}

func (c *credit) Take(controller verifier, credit reputation) bool {
	var result bool
	isSumLessMax := controller.Verify(c.desiredCredit)
	isGoodHistory := credit.CheckHistory(c.name)
	isBalanceGood := c.accountMoney >= c.desiredCredit
	if (isSumLessMax && (isGoodHistory || isBalanceGood)) ||
		(isGoodHistory && (isSumLessMax || isBalanceGood)) {
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
