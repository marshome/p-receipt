package models

type ErrorInvalidArgument struct {

	// Error code
	// Required: true
	Field *string `json:"field"`

	// Error msg
	// Required: true
	Message *string `json:"message"`
}
