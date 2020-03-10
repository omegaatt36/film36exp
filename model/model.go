package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Film struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Vendor          string             `json:"Vendor,omitempty" bson:"Vendor,omitempty"`
	Production      string             `json:"Production,omitempty" bson:"Production,omitempty"`
	ISO             string             `json:"ISO,omitempty" bson:"ISO,omitempty"`
	ForceProcessing string             `json:"ForceProcessing,omitempty" bson:"ForceProcessing,omitempty"`
	FilmSize        string             `json:"FilmSize,omitempty" bson:"FilmSize,omitempty"`
	FilmType        string             `json:"FilmType,omitempty" bson:"FilmType,omitempty"`
	CreateAt        string             `json:"CreateAt,omitempty" bson:"CreateAt,omitempty"`
	UptateAt        string             `json:"UptateAt,omitempty" bson:"UptateAt,omitempty"`
	Description     string             `json:"Description,omitempty" bson:"Description,omitempty"`
	Pics            []*Pic             `json:"Pics,omitempty" bson:"Pics,omitempty"`
}

// Pic is the smallest unit in this system.
type Pic struct {
	ID       string `josn:"ID"`
	Camera   string `josn:"Camera"`
	Lens     string `josn:"Lens"`
	Aperture string `josn:"Aperture"`
	Shutter  string `josn:"Shutter"`
	Notes    string `josn:"Notes"`
}
