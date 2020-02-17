package history

import "github.com/stretchr/testify/mock"

type Mock struct {
	mock.Mock
}

func (m *Mock) CheckHistory(name string) bool{
	args := m.Called(name)
	if _, ok := args.Get(0).(bool); ok {
		return true
	}
	return false
}
