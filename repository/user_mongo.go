package repository

import (
	"github.com/michelaquino/golang_api_skeleton/context"
	"github.com/michelaquino/golang_api_skeleton/models"
)

// UserMongoRepository is a user repository for MongoDB
type UserMongoRepository struct{}

// Insert is a method that insert a user on database
func (u UserMongoRepository) Insert(requestLogData models.RequestLogData, userToInsert models.UserModel) error {
	log := context.GetLogger()

	dbSession := context.GetMongoSession()
	defer dbSession.Close()

	connection := dbSession.DB("api").C("user")
	err := connection.Insert(&userToInsert)
	if err != nil {
		log.Error("UserMongoRepository", "Create", requestLogData.ID, requestLogData.OriginIP, "Create user", "Error", err.Error())
		return err
	}

	log.Error("UserMongoRepository", "Create", requestLogData.ID, requestLogData.OriginIP, "Create user", "Success", "")
	return nil
}
