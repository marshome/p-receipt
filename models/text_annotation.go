package models

type TextAnnotation struct {
	Pages []*Page `json:"pages,omitempty"`
	Text  string  `json:"text,omitempty"`
}
