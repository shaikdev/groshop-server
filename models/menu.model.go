package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Image      string             `json:"image,omitempty"`
	Name       string             `json:"name,omitempty"`
	CreatedAt  time.Time          `json:"created_at,omitempty"`
	ModifiedAt time.Time          `json:"modified_at,omitempty"`
	IsDeleted  bool               `json:"is_deleted"`
}

func (m *Menu) MenuBodyCheck() (bool, string) {
	if m.Image == "" {
		return true, "Image field is required"
	} else if m.Name == "" {
		return true, "Name field is required"
	} else {
		return false, ""
	}
}
