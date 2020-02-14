package verifier

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testVerifier = "Test verifier"
	maximumCredit = "The maximum loan amount"
	desiredCredit = "Desired credit"
	maxCreditSize = 1000
)

var (
	desiredCredits = []int{
		100,
		200,
		300,
		1500,
		500,
		5000,
	}

	want = []bool {
		true,
		true,
		true,
		false,
		true,
		false,
	}
)

func TestVerifier_Verify(t *testing.T) {
	t.Run(testVerifier, func(t *testing.T) {
		controller := NewVerifier(maxCreditSize)
		for i, desired := range desiredCredits {
			result:= controller.Verify(desired)
			assert.Equal(t, want[i], result, fmt.Sprintf("%s - %d\n%s - %d",
				desiredCredit, desired, maximumCredit, maxCreditSize))
		}
	})
}
