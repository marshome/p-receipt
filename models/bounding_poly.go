package models

type BoundingPoly struct {

	// The bounding polygon vertices.
	Vertices []*Vertex `json:"vertices,omitempty"`
}
