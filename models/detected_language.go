package models

type DetectedLanguage struct {
	Confidence   float64 `json:"confidence,omitempty"`
	LanguageCode string  `json:"languageCode,omitempty"`
}
