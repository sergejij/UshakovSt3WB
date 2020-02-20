package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	maxCreditSizes = []int{
		1000,
		1500,
		2000,
		3000,
		1000,
		2000,
		0,
		10,
	}
	desiredCredits = []int{
		0,
		100,
		200,
		300,
		1600,
		500,
		5000,
		15000,
	}
	want = []bool{
		true,
		true,
		true,
		true,
		false,
		true,
		false,
		false,
	}
)

// TestVerifier_Verify ...
func TestVerifier_Verify(t *testing.T) {
	for i := 0; i < len(desiredCredits); i++ {
		comparison := NewNameValidator(maxCreditSizes[i])
		res := comparison.Validate(desiredCredits[i])
		assert.Equal(t, want[i], res)
	}
}
