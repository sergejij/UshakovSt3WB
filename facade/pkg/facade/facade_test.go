package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"task1/pkg/history"
	"task1/pkg/validator"
)

const (
	testName           = "Test facade"
	methodValidate     = "Validate"
	methodCheckHistory = "CheckHistory"
	name               = "Alex"
	accountMoney       = 400
	desiredCredit      = 200
)

// TestCredit_Take ...
func TestCredit_Take(t *testing.T) {
	t.Run(testName, func(t *testing.T) {
		comparisonMock := new(validator.Mock)
		comparisonMock.On(methodValidate, desiredCredit).Return(true).Once()

		historyMock := new(history.Mock)
		historyMock.On(methodCheckHistory, name).Return(true).Once()

		myCredit := NewCredit(comparisonMock, historyMock)
		res := myCredit.Take(name, accountMoney, desiredCredit)
		assert.Equal(t, true, res)
	})

}
