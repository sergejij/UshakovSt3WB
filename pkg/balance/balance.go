package balance

type AmountChecker interface {
	AmountCheck() int
}

type amount struct {
	accountMoney int
}

func (b *amount) AmountCheck() int {
	return b.accountMoney
}

// NewAmount initializes the AmountChecker.
func NewAmount(accMoney int) AmountChecker {
	return &amount{
		accountMoney: accMoney,
	}
}
