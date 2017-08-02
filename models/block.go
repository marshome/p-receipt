package models

type Block struct {

	// Detected block type (text, image etc) for this block.
	BlockType string `json:"blockType,omitempty"`

	// The bounding box for the block.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `json:"boundingBox,omitempty"`

	// List of paragraphs in this block (if this blocks is of type text).
	Paragraphs []*Paragraph `json:"paragraphs,omitempty"`

	// Additional information detected for the block.
	Property *TextProperty `json:"property,omitempty"`
}
