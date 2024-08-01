package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Cart model `my_cart` table

type Cart struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	ProductId primitive.ObjectID `bson:"product_id" json:"product_id"`
	Checkout  bool               `bson:"checkout,omitempty" json:"checkout"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt  time.Time          `bson:"updated_at" json:"updated_at"`
}
