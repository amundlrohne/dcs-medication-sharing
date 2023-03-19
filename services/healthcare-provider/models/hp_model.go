package models

type Provider struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Country  string `json:"country,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
	Username string `json:"username,omitempty" validate:"required"`
}

type User struct {
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}
