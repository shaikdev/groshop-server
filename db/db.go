package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName string = "gropshop"
const userModel string = "user"
const menuModel string = "menu"

// TODO: important thing

var User *mongo.Collection
var Menu *mongo.Collection

// TODO: connect with MongoDB

func init() {
	// Load environment variables from .env file
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("Error loading .env file", envErr)
	}

	connectionString := os.Getenv("DB")

	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	defer fmt.Println("MongoDB connected successfully")

	User = client.Database(dbName).Collection(userModel)
	Menu = client.Database(dbName).Collection(menuModel)

	fmt.Println("Collection instance is ready")

}
