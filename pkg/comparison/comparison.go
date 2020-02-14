package comparison

type verifier struct {
	maxCredit int
}

type Verifier interface {
	Verify(loanSize int) bool
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
