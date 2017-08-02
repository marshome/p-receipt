package models

type DetectedBreak struct {

	// True if break prepends the element.
	IsPrefix bool `json:"isPrefix,omitempty"`

	// Detected break type.
	Type string `json:"type,omitempty"`
}
