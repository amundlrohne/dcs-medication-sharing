package models

type Consent struct {
	ToPublicKey   string `json:"topublickey,omitempty" validate:"required"`
	FromPublicKey string `json:"frompublickey,omitempty" validate:"required"` // use for GET method
	ExpDate       string `json:"expdate,omitempty" validate:"required"`
	DateCreated   string `json:"datecreated,omitempty" validate:"required"`
}
