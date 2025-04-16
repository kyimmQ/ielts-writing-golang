package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" bson:"_id"`
	Username string    `json:"username" bson:"username"`
	Name     string    `json:"name" bson:"name"`
	Email    string    `json:"email" bson:"email"`
	Password string    `json:"password" bson:"password"`
}
