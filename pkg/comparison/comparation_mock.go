package comparison

import (
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Verify(desiredCredit int) bool {
	args := m.Called(desiredCredit)
	if _, ok := args.Get(0).(int); ok {
		return true
	}
	return false
}
