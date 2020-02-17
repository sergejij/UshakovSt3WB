package history

import "github.com/stretchr/testify/mock"

// Mock of history
type Mock struct {
	mock.Mock
}

// CheckHistory ...
func (m *Mock) CheckHistory(name string) bool{
	args := m.Called(name)
	if _, ok := args.Get(0).(bool); ok {
		return true
	}
	return false
}
