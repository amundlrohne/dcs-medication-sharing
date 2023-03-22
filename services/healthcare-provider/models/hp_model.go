package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Provider struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Country  string             `json:"country,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
	Username string             `json:"username,omitempty" validate:"required"`
}

type User struct {
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type JWT struct {
	Username string `json:"username,omitempty" validate:"required"`
	Token    string `json:"token,omitempty" validate:"required"`
}
