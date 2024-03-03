package models

// UserModel is a struct that represents the user
type UserModel struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
}
