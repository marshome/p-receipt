package models

type Page struct {

	// List of blocks of text, images etc on this page.
	Blocks []*Block `json:"blocks,omitempty"`

	// Page height in pixels.
	Height int64 `json:"height,omitempty"`

	// Additional information detected on the page.
	Property *TextProperty `json:"property,omitempty"`

	// Page width in pixels.
	Width int64 `json:"width,omitempty"`
}
