package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/shaikdev/groshop-server/db"
	"github.com/shaikdev/groshop-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(body models.User) (primitive.ObjectID, error) {
	response, err := db.User.InsertOne(context.Background(), body)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	insertedID := response.InsertedID.(primitive.ObjectID)
	return insertedID, nil

}

func GetUser(id string, email string) (models.User, error) {
	filter := bson.M{}
	if id != "" {
		_id, _ := primitive.ObjectIDFromHex(id)
		filter["_id"] = _id
	}
	if email != "" {
		filter["email"] = email
	}
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

func UpdateUser(id string, body models.User) (bool, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	update := updateUserField(body)
	setBody := bson.M{"$set": update}
	response, err := db.User.UpdateOne(context.Background(), filter, setBody)
	if err != nil {
		return false, err
	}
	if response.ModifiedCount == 0 {
		return false, nil
	} else {
		return true, nil
	}

}

func updateUserField(body models.User) bson.M {
	update := bson.M{}
	if body.Email != "" {
		update["email"] = body.Email
	}
	if body.Name != "" {
		update["name"] = body.Name
	}
	return update
}

func DeleteUser(id string) (int, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	response, err := db.User.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return int(response.DeletedCount), nil

}

func DeleteAllUser() (int, error) {
	response, err := db.User.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		return 0, err
	}
	return int(response.DeletedCount), nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil

}

func ComparePasswords(hashedPassword string, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}

func GenerateJwtToken(user models.User) (string, error) {
	// Calculate the expiration time (30 days from now)
	expirationTime := time.Now().Add(30 * 24 * time.Hour)

	secretKey := os.Getenv("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"_id":  user.Id,
		"exp":  expirationTime.Unix(),
	})

	// Sign the token with your secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")

}
