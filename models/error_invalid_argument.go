package models

type ErrorInvalidArgument struct {
	Field   *string `json:"field"`
	Message *string `json:"message"`
}
