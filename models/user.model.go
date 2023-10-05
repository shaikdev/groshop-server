package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty"`
	Email      string             `json:"email,omitempty"`
	ProfilePic string             `json:"profile_pic"`
	Address    []*Address         `json:"address"`
}

type Address struct {
	DoorNumber  string `json:"door_number,omitempty"`
	StreetName  string `json:"street_name,omitempty"`
	VillageName string `json:"village_name,omitempty"`
	PinCode     int    `json:"pin_code,omitempty"`
}

func (u *User) UserBodyCheck() (bool, string) {

	if u.Name == "" {
		return true, "Name is required"
	} else if u.Email == "" {
		return true, "Email is required"
	}
	return false, ""
}
