package models

type DetectedBreak struct {
	IsPrefix bool   `json:"isPrefix,omitempty"`
	Type     string `json:"type,omitempty"`
}
