package mocks

import "github.com/stretchr/testify/mock"

type verifierMock struct {
	mock.Mock
}

type VerifierMock interface {
	Verify(desiredCredit int) (bool, error)
}

func (v *verifierMock) Verify(desiredCredit int) (bool, error) {
	args := v.Called(desiredCredit)
	return args.Bool(0), args.Error(1)
}

func NewVerifierMock(maxCredit int) VerifierMock {
	return &verifierMock {}
}
