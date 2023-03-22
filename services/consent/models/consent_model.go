package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Consent struct {
	ID            primitive.ObjectID `bson:"_id"`
	ToPublicKey   string             `json:"topublickey,omitempty" validate:"required"`
	FromPublicKey string             `json:"frompublickey,omitempty" validate:"required"` // use for GET method
	ExpDate       string             `json:"expdate,omitempty" validate:"required"`
	DateCreated   string             `json:"datecreated,omitempty" validate:"required"`
}
