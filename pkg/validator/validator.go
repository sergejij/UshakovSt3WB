package validator

// Verifier checks whether the desiredCredit is greater than the maxCredit.
type NameValidator interface {
	Validate(loanSize int) bool
}

type nameValidator struct {
	maxCredit int
}

// Validate ...
func (v *nameValidator) Validate(desiredCredit int) bool {
	return desiredCredit <= v.maxCredit
}

// NewVerifier initializes the Verifier.
func NewNameValidator(maxCredit int) NameValidator {
	return &nameValidator{
		maxCredit: maxCredit,
	}
}
