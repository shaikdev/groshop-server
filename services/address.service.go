package services

import (
	"context"

	"github.com/shaikdev/groshop-server/db"
	"github.com/shaikdev/groshop-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAddress(id string, body models.User) (int64, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	setQuery := bson.M{"$push": bson.M{"address": bson.M{"$each": body.Address, "$position": 0}}}
	response, err := db.User.UpdateOne(context.Background(), filter, setQuery)
	if err != nil {
		return 0, err
	}
	return response.ModifiedCount, nil
}

func UpdateAddress(id string, body models.Address) (int64, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	body.Id = _id
	filter := bson.M{"address._id": _id}
	setQuery := bson.M{"$set": bson.M{"address.$": body}}
	response, err := db.User.UpdateOne(context.Background(), filter, setQuery)
	if err != nil {
		return 0, err
	}
	return response.ModifiedCount, nil
}

func DeleteAddressById(id string, addressId string) (int64, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	address, _ := primitive.ObjectIDFromHex(addressId)
	setQuery := bson.M{"$pull": bson.M{"address": bson.M{"_id": address}}}
	response, err := db.User.UpdateOne(context.Background(), filter, setQuery)
	if err != nil {
		return 0, err
	}
	return response.ModifiedCount, nil
}
