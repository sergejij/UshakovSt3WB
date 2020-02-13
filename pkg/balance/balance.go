package balance

type BalanceChecker interface {
	BalanceCheck() int
}
type balance struct {
	accountMoney int
}
func (b *balance) BalanceCheck() int {
	return b.accountMoney
}

// NewBalance initializes the BalanceChecker.
func NewBalance(accMoney int) (BalanceChecker, error) {
	return &balance{
		accountMoney: accMoney,
	}, nil
}
