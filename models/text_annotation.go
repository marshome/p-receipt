package models

type TextAnnotation struct {

	// List of pages detected by OCR.
	Pages []*Page `json:"pages,omitempty"`

	// UTF-8 text detected on the pages.
	Text string `json:"text,omitempty"`
}
