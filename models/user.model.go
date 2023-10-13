package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty"`
	Email      string             `json:"email,omitempty"`
	ProfilePic string             `json:"profile_pic,omitempty"`
	Password   string             `json:"password"`
	Address    []*Address         `json:"address"`
}

type Address struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DoorNumber string             `json:"door_number,omitempty"`
	StreetName string             `json:"street_name,omitempty"`
	PinCode    int                `json:"pincode,omitempty"`
}

func (u *User) UserBodyCheck() (bool, string) {
	if u.Name == "" {
		return true, "Name is required"
	} else if u.Email == "" {
		return true, "Email is required"
	} else if u.Password == "" {
		return true, "Password is required"
	} else {
		return false, ""
	}

}

func (u *User) LoginBodyCheck() (bool, string) {
	if u.Email == "" {
		return true, "Email is required"
	} else if u.Password == "" {
		return true, "Password is required"
	} else {
		return false, ""
	}
}

func (u *User) EditUser() (bool, string) {
	if u.Email == "" {
		return true, "Email is required"
	} else if u.Name == "" {
		return true, "Name is required"
	} else {
		return false, ""
	}
}

func (a *Address) AddressBodyCheck() (bool, string) {
	if a.DoorNumber == "" {
		return true, "Door number is required"
	} else if a.StreetName == "" {
		return true, "Street name is required"
	} else if a.PinCode == 0 {
		return true, "Pin code is required"
	} else {
		return false, ""
	}
}

func InitializeAddress() *User {
	return &User{
		Address: []*Address{},
	}
}
