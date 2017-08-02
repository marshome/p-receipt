package models

type DetectedLanguage struct {

	// Confidence of detected language. Range [0, 1].
	Confidence float64 `json:"confidence,omitempty"`

	// The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	// information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode string `json:"languageCode,omitempty"`
}
