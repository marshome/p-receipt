package models

type Paragraph struct {
	BoundingBox *BoundingPoly `json:"boundingBox,omitempty"`
	Property    *TextProperty `json:"property,omitempty"`
	Words       []*Word       `json:"words,omitempty"`
}
