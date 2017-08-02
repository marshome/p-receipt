package models

type Symbol struct {
	BoundingBox *BoundingPoly `json:"boundingBox,omitempty"`
	Property    *TextProperty `json:"property,omitempty"`
	Text        string        `json:"text,omitempty"`
}
