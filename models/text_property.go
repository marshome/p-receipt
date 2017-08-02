package models

type TextProperty struct {

	// Detected start or end of a text segment.
	DetectedBreak *DetectedBreak `json:"detectedBreak,omitempty"`

	// A list of detected languages together with confidence.
	DetectedLanguages []*DetectedLanguage `json:"detectedLanguages,omitempty"`
}
