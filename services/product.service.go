package services

import (
	"context"

	"github.com/shaikdev/groshop-server/db"
	"github.com/shaikdev/groshop-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(body models.Product) (primitive.ObjectID, error) {
	response, err := db.Product.InsertOne(context.Background(), body)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	insertId := response.InsertedID.(primitive.ObjectID)
	return insertId, nil
}

func GetProduct(id string) (models.Product, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id, "isdeleted": false}
	var product models.Product
	err := db.Product.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func GetProducts() ([]models.Product, error) {
	response, err := db.Product.Find(context.Background(), bson.M{"isdeleted": false})
	if err != nil {
		return []models.Product{}, err
	}
	var products []models.Product
	for response.Next(context.Background()) {
		var product models.Product
		decodeErr := response.Decode(&product)
		if decodeErr != nil {
			return []models.Product{}, decodeErr
		}
		products = append(products, product)
	}
	return products, nil
}

func UpdateProduct(id string, body models.Product) (bool, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id, "isdeleted": false}
	update := bson.M{}
	setBody := bson.M{"$set": update}
	_, err := db.Product.UpdateOne(context.Background(), filter, setBody)
	if err != nil {
		return false, err
	}
	return true, nil
}
func DeleteProduct(id string) (bool, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id, "isdeleted": false}
	update := bson.M{"isdeleted": true}
	setBody := bson.M{"$set": update}
	_, err := db.Product.UpdateOne(context.Background(), filter, setBody)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteProducts() (bool, error) {
	filter := bson.M{"isdeleted": false}
	update := bson.M{"isdeleted": true}
	setBody := bson.M{"$set": update}
	_, err := db.Product.UpdateOne(context.Background(), filter, setBody)
	if err != nil {
		return false, err
	}
	return true, nil
}
