package repository

import "github.com/michelaquino/golang_api_skeleton/models"

// UserRepository is a interface that define the methods to user repository
type UserRepository interface {

	// Insert is a method that insert a user on database
	Insert(requestLogData models.RequestLogData, userToInsert models.UserModel) error
}
