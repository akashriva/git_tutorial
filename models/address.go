package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// adderss model `addresses`table
type Address struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	City      string             `bson:"city" json:"city"`
	Country   string             `bson:"country" json:"country"`
	Street    string             `bson:"street" json:"street"`
	ZipCode   string             `bson:"zip_code" json:"zip_code"`
	UserId    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt  time.Time          `bson:"updated_at" json:"updated_at"` //refrence to user model
}
