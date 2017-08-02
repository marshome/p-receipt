package models

type Word struct {
	BoundingBox *BoundingPoly `json:"boundingBox,omitempty"`
	Property    *TextProperty `json:"property,omitempty"`
	Symbols     []*Symbol     `json:"symbols,omitempty"`
}
