package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Medication struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Dose  string             `json:"dose,omitempty" validate:"required"`
	Title string             `json:"title,omitempty" validate:"required"`
}
