package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "mongodb+srv://shaiksha19:shaiksha19@socialapp.5gly5.mongodb.net/groshop?retryWrites=true&w=majority"
const dbName string = "gropshop"
const userModel string = "user"

// TODO: important thing

var User *mongo.Collection

// TODO: connect with MongoDB

func init() {

	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	defer fmt.Println("MongoDB connected successfully")

	User = client.Database(dbName).Collection(userModel)

	fmt.Println("Collection instance is ready")

}
