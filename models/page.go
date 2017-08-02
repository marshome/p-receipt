package models

type Page struct {
	Blocks   []*Block      `json:"blocks,omitempty"`
	Height   int64         `json:"height,omitempty"`
	Property *TextProperty `json:"property,omitempty"`
	Width    int64         `json:"width,omitempty"`
}
