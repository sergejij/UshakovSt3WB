package comparison

import (
	"github.com/stretchr/testify/mock"
)

// Mock of comparison
type Mock struct {
	mock.Mock
}

// Verify ...
func (m *Mock) Verify(desiredCredit int) bool {
	args := m.Called(desiredCredit)
	if _, ok := args.Get(0).(int); ok {
		return true
	}
	return false
}
