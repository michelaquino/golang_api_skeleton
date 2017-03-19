package repository

import (
	"github.com/michelaquino/golang_api_skeleton/models"
	testifyMock "github.com/stretchr/testify/mock"
)

// UserRepositoryMock is a mock that implements the UserRepository interface
type UserRepositoryMock struct {
	testifyMock.Mock
}

// Insert is a method that insert a user on database
func (mock *UserRepositoryMock) Insert(requestLogData models.RequestLogData, userToInsert models.UserModel) error {
	args := mock.Called(requestLogData, userToInsert)
	return args.Error(0)
}
