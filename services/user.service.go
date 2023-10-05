package services

import (
	"context"

	"github.com/shaikdev/groshop-server/db"
	"github.com/shaikdev/groshop-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(body models.User) (primitive.ObjectID, error) {
	response, err := db.User.InsertOne(context.Background(), body)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	insertedID := response.InsertedID.(primitive.ObjectID)
	return insertedID, nil

}

func GetUserById(id string) (models.User, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var user models.User
	err := db.User.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUsers() ([]primitive.M, error) {
	response, err := db.User.Find(context.Background(), bson.D{{}})
	if err != nil {
		return []primitive.M{}, err
	}
	var users []primitive.M
	for response.Next(context.Background()) {
		var user bson.M
		usersErr := response.Decode(&user)
		if usersErr != nil {
			return []primitive.M{}, usersErr
		}
		users = append(users, user)

	}
	return users, nil

}

func UpdateUser(id string, body models.User) error {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	update := bson.M{}
	setBody := bson.M{"$set": update}
	_, err := db.User.UpdateOne(context.Background(), filter, setBody)
	return err

}

func DeleteUser(id string) int {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	response, _ := db.User.DeleteOne(context.Background(), filter)
	return int(response.DeletedCount)

}

func DeleteAllUser() int {
	response, _ := db.User.DeleteMany(context.Background(), bson.D{{}})
	return int(response.DeletedCount)
}
