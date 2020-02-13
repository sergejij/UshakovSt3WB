package bank

type Verifier interface {
	Verify(loanSize int) string
}
type bank struct {

}
func (b *bank) Verify(loanSize int) string {
	if loanSize < 500 {
		return "approved"
	}
	return "denied"
}

// NewBank initializes the Verifier.
func NewBank() (Verifier, error) {
	return &bank{}, nil
}
