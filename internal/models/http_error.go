package models

type HttpError struct {
	Error   int    `json:"error" example:"404"`
	Message string `json:"message" example:"NOT_FOUND"`
	Details string `json:"details,omitempty" example:"User with ID 123 does not exist"`
}
