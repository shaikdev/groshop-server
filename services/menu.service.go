package services

import (
	"context"
	"time"

	"github.com/shaikdev/groshop-server/db"
	"github.com/shaikdev/groshop-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMenu(body models.Menu) (primitive.ObjectID, error) {
	response, err := db.Menu.InsertOne(context.Background(), body)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	responseId := response.InsertedID.(primitive.ObjectID)
	return responseId, nil
}

func GetMenu(id string, name string) (models.Menu, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"isdeleted": false}

	if id != "" {
		filter["_id"] = _id
	}
	if name != "" {
		filter["name"] = name
	}

	var menu models.Menu
	err := db.Menu.FindOne(context.Background(), filter).Decode(&menu)
	if err != nil {
		return models.Menu{}, err
	}
	return menu, nil
}

func GetMenus() ([]models.Menu, error) {
	filter := bson.M{"isdeleted": false}
	response, err := db.Menu.Find(context.Background(), filter)
	if err != nil {
		return []models.Menu{}, err
	}
	var menus []models.Menu
	for response.Next(context.Background()) {
		var menu models.Menu
		decodeMenuErr := response.Decode(&menu)
		if decodeMenuErr != nil {
			return []models.Menu{}, decodeMenuErr
		}
		menus = append(menus, menu)
	}
	return menus, nil
}

func UpdateMenu(id string, body models.Menu) (bool, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id, "isdeleted": false}
	update := bson.M{"modifiedat": time.Now()}
	if body.Image != "" {
		update["image"] = body.Image
	}
	if body.Name != "" {
		update["name"] = body.Name
	}
	setBody := bson.M{"$set": update}
	_, err := db.Menu.UpdateOne(context.Background(), filter, setBody)
	if err != nil {
		return false, err
	}
	return true, nil

}

func DeleteMenu(id string) (bool, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id, "isdeleted": false}
	setBody := bson.M{"$set": bson.M{"isdeleted": true}}
	_, err := db.Menu.UpdateOne(context.Background(), filter, setBody)
	if err != nil {
		return false, err
	}
	return true, nil

}

func DeleteMenus() (bool, error) {
	filter := bson.M{"isdeleted": false}
	setBody := bson.M{"$set": bson.M{"isdeleted": true}}
	_, err := db.Menu.UpdateMany(context.Background(), filter, setBody)
	if err != nil {
		return false, err
	}
	return true, nil
}
