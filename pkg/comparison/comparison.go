package comparison

// Verifier checks whether the desiredCredit is greater than the maxCredit.
type Verifier interface {
	Verify(loanSize int) bool
}

type verifier struct {
	maxCredit int
}

func (v *verifier) Verify(desiredCredit int) bool {
	return desiredCredit <= v.maxCredit
}

// NewVerifier initializes the Verifier.
func NewVerifier(maxCredit int) Verifier {
	return &verifier {
		maxCredit: maxCredit,
	}
}
