package facade

type historyReputation interface {
	CheckHistory(name string) bool
}

type nameValidator interface {
	Validate(loanSize int) bool
}

// Credit is a service that takes credit
type Credit interface {
	Take(name string, accountMoney int, desiredCredit int) bool
}

type credit struct {
	nameValidator     nameValidator
	historyReputation historyReputation
}

// Take ...
func (c *credit) Take(name string, accountMoney int, desiredCredit int) bool {
	var (
		isSumLessMax, isGoodHistory, isBalanceGood bool
	)
	isSumLessMax = c.nameValidator.Validate(desiredCredit)
	isGoodHistory = c.historyReputation.CheckHistory(name)
	isBalanceGood = accountMoney >= desiredCredit
	return (isSumLessMax && isGoodHistory) ||
		(isBalanceGood && isGoodHistory) ||
		(isSumLessMax && isBalanceGood)
}

// NewCredit initializes the Credit.
func NewCredit(nameValidator nameValidator, historyReputation historyReputation) Credit {
	return &credit{
		nameValidator:     nameValidator,
		historyReputation: historyReputation,
	}
}
