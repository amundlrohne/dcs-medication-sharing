package models

type Consent struct {
	ToPublicKey   string `json:"topublickey,omitempty" validate:"required"`
	FromPublicKey string `json:"frompublickey,omitempty" validate:"required"`
	ExpDate       string `json:"expdate,omitempty" validate:"required"`
}
