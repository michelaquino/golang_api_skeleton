package repository

import (
	"context"
	"log/slog"

	"github.com/michelaquino/golang_api_skeleton/src/models"
	"github.com/michelaquino/golang_api_skeleton/src/mongo"
)

var userMongoCollectionName = "user"

// UserMongoRepository is a user repository for MongoDB.
type UserMongoRepository struct{}

// Insert is a method that inserts an user into database.
func (u UserMongoRepository) Insert(ctx context.Context, userToInsert models.UserModel) error {
	// Execute Mongo's Insert
	err := mongo.Insert(ctx, userMongoCollectionName, &userToInsert)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return err
	}

	slog.ErrorContext(ctx, err.Error())
	return nil
}
