package verifier

type Verifier interface {
	Verify(loanSize int) bool
}

type verifier struct {
	maximumLoanAmount int
}

func (v *verifier) Verify(desiredCredit int) bool {
	if desiredCredit <= v.maximumLoanAmount {
		return true
	}
	return false
}

// NewBank initializes the Verifier.
func NewVerifier(maxCredit int) Verifier {
	return &verifier{
		maximumLoanAmount: maxCredit,
	}
}

