package validator

import (
	"github.com/stretchr/testify/mock"
)

// Mock of validator
type Mock struct {
	mock.Mock
}

// Validate ...
func (m *Mock) Validate(desiredCredit int) bool {
	args := m.Called(desiredCredit)
	if _, ok := args.Get(0).(int); ok {
		return true
	}
	return false
}
