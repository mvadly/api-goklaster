package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

type AuthLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (auth AuthLogin) Validate() error {
	return validation.ValidateStruct(&auth,
		validation.Field(&auth.Username, validation.Required, validation.Length(2, 50)),
		validation.Field(&auth.Password, validation.Required, validation.Length(2, 50)),
	)
}
