package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/akashriva/gin_framework/helper"
	"github.com/akashriva/gin_framework/models"
)

// Implementing the Manager interface
func (mgr *manager) Insert(data interface{}, collectionName string) (interface{}, error) {
	log.Println(data)
	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)
	result, err := orgCollection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (mgr *manager) GetSingleRecordByEmail(email string, collectionName string) *models.Verification {
	resp := &models.Verification{}
	filter := bson.D{{"email", email}}
	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)
	_ = orgCollection.FindOne(context.TODO(), filter).Decode(&resp)
	fmt.Println(resp)
	return resp
}

func (mgr *manager) UpdateVerification(data models.Verification, collectionName string) error {
	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)
	filter := bson.D{{"email", data.Email}}
	update := bson.D{{"$set", data}}
	_, err := orgCollection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (mgr *manager) UpdateEmailVerifiedStatus(req models.Verification, collectionName string) error {
	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)
	filter := bson.D{{"email", req.Email}}
	update := bson.D{{"$set", req}}
	_, err := orgCollection.UpdateOne(context.TODO(), filter, update)
	return err
}


// Get single user from db
// func (mgr *manager) GetSingleRecordByEmailForUser(email, collectionName string) models.User {
// 	resp := models.User{}
// 	filter := bson.D{{"email", email}}
// 	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)

// 	_ = orgCollection.FindOne(context.TODO(), filter).Decode(&resp)
// 	fmt.Println(resp)
// 	return resp
// }

// func (mgr *manager) GetListProducts(page, limit, offset int, collectionName string) (products []models.Product, count int64, err error) {
// 	skip := ((page - 1) * limit)
// 	if offset > 0 {
// 		skip = offset
// 	}
// 	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)
// 	findOptions := options.Find()
// 	findOptions.SetSkip(int64(skip))
// 	findOptions.SetLimit(int64(limit))

// 	cur, err := orgCollection.Find(context.TODO(), bson.M{}, findOptions)
// 	err = cur.All(context.TODO(), &products)
// 	itemCount, err := orgCollection.CountDocuments(context.TODO(), bson.M{})
// 	return products, itemCount, err
// }

// func (mgr *manager) SearchProduct(page, limit, offset int, search string, collectionName string) (products []models.Product, count int64, err error) {
// 	skip := ((page - 1) * limit)
// 	if offset > 0 {
// 		skip = offset
// 	}

// 	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)
// 	findOptions := options.Find()
// 	findOptions.SetSkip(int64(skip))
// 	findOptions.SetLimit(int64(limit))

// 	searchFilter := bson.M{}
// 	if len(search) >= 3 {
// 		searchFilter["$or"] = []bson.M{
// 			{"name": primitive.Regex{Pattern: "." + search + ".", Options: "i"}},
// 			{"description": primitive.Regex{Pattern: "." + search + ".", Options: "i"}},
// 		}
// 	}
// 	cur, err := orgCollection.Find(context.TODO(), searchFilter, findOptions)
// 	cur.All(context.TODO(), &products)
// 	count, err = orgCollection.CountDocuments(context.TODO(), searchFilter)
// 	return products, count, err
// }

// func (mgr *manager) GetSingleProductById(id primitive.ObjectID, collectionName string) (product models.Product, err error) {
// 	filter := bson.D{{"_id", id}}
// 	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)

// 	err = orgCollection.FindOne(context.TODO(), filter).Decode(&product)
// 	return product, err
// }

// func (mgr *manager) UpdateProduct(p models.Product, collectionName string) error {
// 	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)
// 	filter := bson.D{{"_id", p.ID}}
// 	update := bson.D{{"$set", p}}

// 	_, err := orgCollection.UpdateOne(context.TODO(), filter, update)
// 	return err
// }

// func (mgr *manager) DeleteProduct(id primitive.ObjectID, collectionName string) error {
// 	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)
// 	filter := bson.D{{"_id", id}}

// 	_, err := orgCollection.DeleteOne(context.TODO(), filter)
// 	return err
// }

// func (mgr *manager) GetSingleAddress(id primitive.ObjectID, collectionName string) (address models.Address, err error) {
// 	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)
// 	filter := bson.D{{"user_id", id}}

// 	err = orgCollection.FindOne(context.TODO(), filter).Decode(&address)
// 	return address, err
// }

// func (mgr *manager) GetSingleUserByUserId(id primitive.ObjectID, collectionName string) (user models.User) {
// 	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)
// 	filter := bson.D{{"_id", id}}

// 	_ = orgCollection.FindOne(context.TODO(), filter).Decode(&user)
// 	return user
// }

// func (mgr *manager) GetCartObjectById(id primitive.ObjectID, collectionName string) (c models.Cart, err error) {
// 	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)
// 	filter := bson.D{{"_id", id}}

// 	err = orgCollection.FindOne(context.TODO(), filter).Decode(&c)
// 	return c, err
// }

// func (mgr *manager) UpdateUser(u models.User, collectionName string) error {
// 	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)
// 	filter := bson.D{{"_id", u.Id}}
// 	update := bson.D{{"$set", u}}

// 	_, err := orgCollection.UpdateOne(context.TODO(), filter, update)
// 	return err
// }
// func (mgr *manager) UpdateCartToCheckOut(c models.Cart, collectionName string) error {
// 	orgCollection := mgr.Connection.Database(helper.Database).Collection(collectionName)
// 	filter := bson.D{{"_id", c.ID}}
// 	update := bson.D{{"$set", c}}

// 	_, err := orgCollection.UpdateOne(context.TODO(), filter, update)
// 	return err
// }
