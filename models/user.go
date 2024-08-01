package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user model table `user`
type User struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Role          string             `bson:"role" json:"role"`
	FirstName     string             `bson:"first_name" json:"first_name"`
	LastName      string             `bson:"last_name" json:"last_name"`
	Email         string             `bson:"email" json:"email"`
	Phone         string             `bson:"phone" json:"phone"`
	CountryCode   string             `bson:"country_code" json:"country_code"`
	PassWord      string             `bson:"password" json:"password"`
	IsEmailVerify bool               `bson:"is_email_verify" json:"is_email_verify"`
	IsPhoneVerfy  bool               `bson:"is_phone_verify" json:"is_phone_verify"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt      time.Time          `bson:"updated_at" json:"updated_at"`
	UserUniqeId   string             `bson:"user_unique_id" json:"user_unique_id"`
}
