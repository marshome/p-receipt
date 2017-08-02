package models

type Block struct {
	BlockType   string        `json:"blockType,omitempty"`
	BoundingBox *BoundingPoly `json:"boundingBox,omitempty"`
	Paragraphs  []*Paragraph  `json:"paragraphs,omitempty"`
	Property    *TextProperty `json:"property,omitempty"`
}
