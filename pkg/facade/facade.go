package facade

// reputation checks whether a person has a good credit reputation.
type reputation interface {
	CheckHistory(name string) bool
}

// verifier checks whether the desiredCredit is greater than the maxCredit.
type verifier interface {
	Verify(loanSize int) bool
}

// Credit is a service that takes credit
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
