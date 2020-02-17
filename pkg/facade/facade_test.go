package facade

import (
	"fasad/pkg/comparison"
	"fasad/pkg/history"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testName = 				"Test facade"
	methodVerify = 			"Verify"
	methodCheckHistory = 	"CheckHistory"
	name = 					"Alex"
	accountMoney = 			400
	desiredCredit = 		200
)

// TestCredit_Take ...
func TestCredit_Take(t *testing.T) {
	t.Run(testName, func(t *testing.T) {
		myCredit := NewCredit(name, accountMoney, desiredCredit)
		comparisonMock := new(comparison.Mock)
		comparisonMock.On(methodVerify, desiredCredit).Return(true).Once()

		historyMock := new(history.Mock)
		historyMock.On(methodCheckHistory, name).Return(true).Once()

		res := myCredit.Take(comparisonMock, historyMock)
		assert.Equal(t, true, res)
	})

}

