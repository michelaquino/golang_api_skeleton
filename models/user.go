package models

// UserModel is a struct that represents the user
type UserModel struct {
	Name  string `bson:"name,omitempty"`
	Email string `bson:"email,omitempty"`
}
