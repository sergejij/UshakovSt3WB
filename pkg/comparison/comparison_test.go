package comparison

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testVerifier = "Test comparison"
	maximumCredit = "The maximum loan amount"
	desiredCredit = "Desired credit"
	maxCreditSize = 1000
)

var (
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
	want = []bool {
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

/*func TestVerifier_Verify(t *testing.T) {
	t.Run(testVerifier, func(t *testing.T) {
		controller := NewVerifier(maxCreditSize)
		for i, desired := range desiredCredits {
			result:= controller.Verify(desired)
			assert.Equal(t, want[i], result, fmt.Sprintf("%s - %d\n%s - %d",
				desiredCredit, desired, maximumCredit, maxCreditSize))
		}
	})
}*/
func TestVerifier_Verify(t *testing.T) {
	comparison := NewVerifier(maxCreditSize)
	for i := 0; i < len(desiredCredits); i++ {
		res := comparison.Verify(desiredCredits[i])
		assert.Equal(t, want[i], res)
	}
}
