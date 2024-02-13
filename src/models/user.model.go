package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserName  string             `bson:"userName"`
	UserEmail string             `bson:"userEmail"`
	Password  string             `bson:"userPassword"`
	IsAdmin   bool               `bson:"isAdmin,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
}