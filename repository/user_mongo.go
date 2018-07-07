package repository

import (
	"github.com/michelaquino/golang_api_skeleton/context"
	"github.com/michelaquino/golang_api_skeleton/models"
	"github.com/michelaquino/golang_api_skeleton/mongo"
)

var (
	userMongoCollectionName = "user"
)

// UserMongoRepository is a user repository for MongoDB
type UserMongoRepository struct{}

// Insert is a method that insert a user on database
func (u UserMongoRepository) Insert(requestLogData models.RequestLogData, userToInsert models.UserModel) error {
	log := context.GetLogger()

	// Execute the insert
	err := mongo.Insert(userMongoCollectionName, &userToInsert)
	if err != nil {
		log.Error("UserMongoRepository", "Create", requestLogData.ID, requestLogData.OriginIP, "Create user", "Error", err.Error())
		return err
	}

	log.Info("UserMongoRepository", "Create", requestLogData.ID, requestLogData.OriginIP, "Create user", "Success", "")
	return nil
}
