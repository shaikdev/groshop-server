package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductImage string             `json:"product_image,omitempty"`
	CoverPhoto   string             `json:"cover_photo,omitempty"`
	OfferValue   int                `json:"offer_value,omitempty"`
	Description  string             `json:"description,omitempty"`
	Rating       int                `json:"rating,omitempty"`
	ShopName     string             `json:"shop_name,omitempty"`
	Quantity     int                `json:"quantity,omitempty"`
	SearchKeys   []string           `json:"search_keys"`
	Search       string             `json:"search"`
	ProductKg    string             `json:"product_kg"`
	DeliveryTime string             `json:"delivery_time,omitempty"`
	IsDeleted    bool               `json:"is_deleted,omitempty"`
	CreatedAt    time.Time          `json:"created_at,omitempty"`
	ModifiedAt   time.Time          `json:"modified_at,omitempty"`
}

func (p *Product) ProductBodyCheck() (bool, string) {
	if p.ProductImage == "" {
		return true, "Product image field is required"
	} else if p.CoverPhoto == "" {
		return true, "Cover photo field is required"
	} else if p.Description == "" {
		return true, "Description field is required"
	} else if p.ShopName == "" {
		return true, "Shop name field is required"
	} else if len(p.SearchKeys) == 0 {
		return true, "Search keys field is required"
	} else if p.ProductKg == "" {
		return true, "Product kg field is required"
	} else if p.DeliveryTime == "" {
		return true, "Delivery time field is required"
	} else {
		return false, ""
	}
}
