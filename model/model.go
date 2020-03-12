package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Film recode one film infomation.
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
}

// Pic is the smallest unit.
type Pic struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FID      primitive.ObjectID `json:"FID,omitempty" bson:"FID,omitempty"`
	Camera   string             `json:"Camera,omitempty" bson:"Camera,omitempty"`
	Lens     string             `json:"Lens,omitempty" bson:"Lens,omitempty"`
	Aperture string             `json:"Aperture,omitempty" bson:"Aperture,omitempty"`
	Shutter  string             `json:"Shutter,omitempty" bson:"Shutter,omitempty"`
	Notes    string             `json:"Notes,omitempty" bson:"Notes,omitempty"`
}
