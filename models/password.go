package models

type ValidatePasswordPayload struct {
	Password string `json:"password" validate:"required,string"`
}
